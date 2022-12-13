package trustmanagerprovider

import (
	"context"
	"terraform-provider-pingdirectory/internal/operations"
	"terraform-provider-pingdirectory/internal/resource/config"
	internaltypes "terraform-provider-pingdirectory/internal/types"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	client "github.com/pingidentity/pingdata-config-api-go-client"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &blindTrustManagerProviderResource{}
	_ resource.ResourceWithConfigure   = &blindTrustManagerProviderResource{}
	_ resource.ResourceWithImportState = &blindTrustManagerProviderResource{}
)

// Create a Blind Trust Manager Provider resource
func NewBlindTrustManagerProviderResource() resource.Resource {
	return &blindTrustManagerProviderResource{}
}

// blindTrustManagerProviderResource is the resource implementation.
type blindTrustManagerProviderResource struct {
	providerConfig internaltypes.ProviderConfiguration
	apiClient      *client.APIClient
}

// blindTrustManagerProviderResourceModel maps the resource schema data.
type blindTrustManagerProviderResourceModel struct {
	Name                     types.String `tfsdk:"name"`
	Enabled                  types.Bool   `tfsdk:"enabled"`
	IncludeJVMDefaultIssuers types.Bool   `tfsdk:"include_jvm_default_issuers"`
	LastUpdated              types.String `tfsdk:"last_updated"`
	Notifications            types.Set    `tfsdk:"notifications"`
	RequiredActions          types.Set    `tfsdk:"required_actions"`
}

// Metadata returns the resource type name.
func (r *blindTrustManagerProviderResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_blind_trust_manager_provider"
}

// GetSchema defines the schema for the resource.
func (r *blindTrustManagerProviderResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	schema := tfsdk.Schema{
		Description: "Manages a Blind Trust Manager Provider.",
		Attributes: map[string]tfsdk.Attribute{
			"name": {
				Description: "Name of the Trust Manager Provider.",
				Type:        types.StringType,
				Required:    true,
				PlanModifiers: []tfsdk.AttributePlanModifier{
					resource.RequiresReplace(),
				},
			},
			"enabled": {
				Description: "Indicate whether the Trust Manager Provider is enabled for use.",
				Type:        types.BoolType,
				Required:    true,
			},
			// Optional boolean fields must be Computed because PD gives them a default value
			"include_jvm_default_issuers": {
				Description: "Indicates whether certificates issued by an authority included in the JVM's set of default issuers should be automatically trusted, even if they would not otherwise be trusted by this provider.",
				Type:        types.BoolType,
				Optional:    true,
				Computed:    true,
			},
		},
	}
	config.AddCommonSchema(&schema)
	return schema, nil
}

// Configure adds the provider configured client to the resource.
func (r *blindTrustManagerProviderResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	providerCfg := req.ProviderData.(internaltypes.ResourceConfiguration)
	r.providerConfig = providerCfg.ProviderConfig
	r.apiClient = providerCfg.ApiClient
}

// Add optional fields to create request
func addOptionalBlindTrustManagerProviderFields(addRequest *client.AddBlindTrustManagerProviderRequest, plan blindTrustManagerProviderResourceModel) {
	// Non string values just have to be defined
	if internaltypes.IsDefined(plan.IncludeJVMDefaultIssuers) {
		boolVal := plan.IncludeJVMDefaultIssuers.ValueBool()
		addRequest.IncludeJVMDefaultIssuers = &boolVal
	}
}

// Create a new resource
func (r *blindTrustManagerProviderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan blindTrustManagerProviderResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	addRequest := client.NewAddBlindTrustManagerProviderRequest(plan.Name.ValueString(),
		[]client.EnumblindTrustManagerProviderSchemaUrn{client.ENUMBLINDTRUSTMANAGERPROVIDERSCHEMAURN_URNPINGIDENTITYSCHEMASCONFIGURATION2_0TRUST_MANAGER_PROVIDERBLIND},
		plan.Enabled.ValueBool())
	addOptionalBlindTrustManagerProviderFields(addRequest, plan)
	// Log request JSON
	requestJson, err := addRequest.MarshalJSON()
	if err == nil {
		tflog.Debug(ctx, "Add request: "+string(requestJson))
	}
	apiAddRequest := r.apiClient.TrustManagerProviderApi.AddTrustManagerProvider(config.BasicAuthContext(ctx, r.providerConfig))
	apiAddRequest = apiAddRequest.AddTrustManagerProviderRequest(
		client.AddBlindTrustManagerProviderRequestAsAddTrustManagerProviderRequest(addRequest))

	trustManagerResponse, httpResp, err := r.apiClient.TrustManagerProviderApi.AddTrustManagerProviderExecute(apiAddRequest)
	if err != nil {
		config.ReportHttpError(ctx, &resp.Diagnostics, "An error occurred while creating the Trust Manager Provider", err, httpResp)
		return
	}

	// Log response JSON
	responseJson, err := trustManagerResponse.MarshalJSON()
	if err == nil {
		tflog.Debug(ctx, "Add response: "+string(responseJson))
	}

	// Read the response into the state
	readBlindTrustManagerProviderResponse(ctx, trustManagerResponse.BlindTrustManagerProviderResponse, &plan)

	// Populate Computed attribute values
	plan.LastUpdated = types.StringValue(string(time.Now().Format(time.RFC850)))

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read a BlindTrustManagerProviderResponse object into the model struct
func readBlindTrustManagerProviderResponse(ctx context.Context, r *client.BlindTrustManagerProviderResponse, state *blindTrustManagerProviderResourceModel) {
	state.Name = types.StringValue(r.Id)
	state.Enabled = types.BoolValue(r.Enabled)
	state.IncludeJVMDefaultIssuers = internaltypes.BoolTypeOrNil(r.IncludeJVMDefaultIssuers)
	// Report any notifications from the Config API
	if r.Urnpingidentityschemasconfigurationmessages20 != nil {
		state.Notifications = internaltypes.GetStringSet(r.Urnpingidentityschemasconfigurationmessages20.Notifications)
		state.RequiredActions, _ = config.GetRequiredActionsSet(*r.Urnpingidentityschemasconfigurationmessages20)
		config.LogMessages(ctx, r.Urnpingidentityschemasconfigurationmessages20)
	} else {
		state.Notifications, _ = types.SetValue(types.StringType, []attr.Value{})
		state.RequiredActions, _ = types.SetValue(config.GetRequiredActionsObjectType(), []attr.Value{})
	}
}

// Read resource information
func (r *blindTrustManagerProviderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state blindTrustManagerProviderResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	trustManagerResponse, httpResp, err := r.apiClient.TrustManagerProviderApi.GetTrustManagerProvider(
		config.BasicAuthContext(ctx, r.providerConfig), state.Name.ValueString()).Execute()
	if err != nil {
		config.ReportHttpError(ctx, &resp.Diagnostics, "An error occurred while getting the Trust Manager Provider", err, httpResp)
		return
	}

	// Log response JSON
	responseJson, err := trustManagerResponse.MarshalJSON()
	if err == nil {
		tflog.Debug(ctx, "Read response: "+string(responseJson))
	}

	// Read the response into the state
	readBlindTrustManagerProviderResponse(ctx, trustManagerResponse.BlindTrustManagerProviderResponse, &state)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Create any update operations necessary to make the state match the plan
func createBlindTrustManagerProviderOperations(plan blindTrustManagerProviderResourceModel, state blindTrustManagerProviderResourceModel) []client.Operation {
	var ops []client.Operation

	operations.AddBoolOperationIfNecessary(&ops, plan.Enabled, state.Enabled, "enabled")
	operations.AddBoolOperationIfNecessary(&ops, plan.IncludeJVMDefaultIssuers, state.IncludeJVMDefaultIssuers, "include-jvm-default-issuers")
	return ops
}

// Update a resource
func (r *blindTrustManagerProviderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan blindTrustManagerProviderResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get the current state to see how any attributes are changing
	var state blindTrustManagerProviderResourceModel
	req.State.Get(ctx, &state)
	updateRequest := r.apiClient.TrustManagerProviderApi.UpdateTrustManagerProvider(config.BasicAuthContext(ctx, r.providerConfig), plan.Name.ValueString())

	// Determine what update operations are necessary
	ops := createBlindTrustManagerProviderOperations(plan, state)
	if len(ops) > 0 {
		updateRequest = updateRequest.UpdateRequest(*client.NewUpdateRequest(ops))
		// Log operations
		operations.LogUpdateOperations(ctx, ops)

		trustManagerResponse, httpResp, err := r.apiClient.TrustManagerProviderApi.UpdateTrustManagerProviderExecute(updateRequest)
		if err != nil {
			config.ReportHttpError(ctx, &resp.Diagnostics, "An error occurred while updating the Trust Manager Provider", err, httpResp)
			return
		}

		// Log response JSON
		responseJson, err := trustManagerResponse.MarshalJSON()
		if err == nil {
			tflog.Debug(ctx, "Update response: "+string(responseJson))
		}

		// Read the response
		readBlindTrustManagerProviderResponse(ctx, trustManagerResponse.BlindTrustManagerProviderResponse, &plan)
		// Update computed values
		plan.LastUpdated = types.StringValue(string(time.Now().Format(time.RFC850)))
	} else {
		tflog.Warn(ctx, "No configuration API operations created for update")
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *blindTrustManagerProviderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state blindTrustManagerProviderResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := r.apiClient.TrustManagerProviderApi.DeleteTrustManagerProviderExecute(
		r.apiClient.TrustManagerProviderApi.DeleteTrustManagerProvider(config.BasicAuthContext(ctx, r.providerConfig), state.Name.ValueString()))
	if err != nil {
		config.ReportHttpError(ctx, &resp.Diagnostics, "An error occurred while deleting the Trust Manager Provider", err, httpResp)
		return
	}
}

func (r *blindTrustManagerProviderResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to Name attribute
	resource.ImportStatePassthroughID(ctx, path.Root("name"), req, resp)
}
