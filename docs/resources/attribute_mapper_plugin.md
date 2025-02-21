---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingdirectory_attribute_mapper_plugin Resource - terraform-provider-pingdirectory"
subcategory: ""
description: |-
  Manages a Attribute Mapper Plugin.
---

# pingdirectory_attribute_mapper_plugin (Resource)

Manages a Attribute Mapper Plugin.

## Example Usage

```terraform
terraform {
  required_version = ">=1.1"
  required_providers {
    pingdirectory = {
      source = "pingidentity/pingdirectory"
    }
  }
}

provider "pingdirectory" {
  username   = "cn=administrator"
  password   = "2FederateM0re"
  https_host = "https://localhost:1443"
  # Warning: The insecure_trust_all_tls attribute configures the provider to trust any certificate presented by the PingDirectory server.
  # It should not be used in production. If you need to specify trusted CA certificates, use the
  # ca_certificate_pem_files attribute to point to any number of trusted CA certificate files
  # in PEM format. If you do not specify certificates, the host's default root CA set will be used.
  # Example:
  # ca_certificate_pem_files = ["/example/path/to/cacert1.pem", "/example/path/to/cacert2.pem"]
  insecure_trust_all_tls = true
  product_version        = "9.2.0.0"
}

# Use "pingdirectory_default_attribute_mapper_plugin" if you are adopting existing configuration from the PingDirectory server into Terraform
resource "pingdirectory_attribute_mapper_plugin" "myAttributeMapperPlugin" {
  id               = "MyAttributeMapperPlugin"
  source_attribute = "cn"
  target_attribute = "sn"
  enabled          = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `enabled` (Boolean) Indicates whether the plug-in is enabled for use.
- `id` (String) Name of this object.
- `source_attribute` (String) Specifies the source attribute type that may appear in client requests which should be remapped to the target attribute. Note that the source attribute type must be defined in the server schema and must not be equal to the target attribute type.
- `target_attribute` (String) Specifies the target attribute type to which the source attribute type should be mapped. Note that the target attribute type must be defined in the server schema and must not be equal to the source attribute type.

### Optional

- `always_map_responses` (Boolean) Indicates whether the target attribute in response messages should always be remapped back to the source attribute. If this is "false", then the mapping will be performed for a response message only if one or more elements of the associated request are mapped. Otherwise, the mapping will be performed for all responses regardless of whether the mapping was applied to the request.
- `description` (String) A description for this Plugin
- `enable_control_mapping` (Boolean) Indicates whether mapping should be applied to attribute types that may be present in specific controls. If enabled, attribute mapping will only be applied for control types which are specifically supported by the attribute mapper plugin.
- `invoke_for_internal_operations` (Boolean) Indicates whether the plug-in should be invoked for internal operations.
- `plugin_type` (Set of String) Specifies the set of plug-in types for the plug-in, which specifies the times at which the plug-in is invoked.

### Read-Only

- `last_updated` (String) Timestamp of the last Terraform update of this resource.
- `notifications` (Set of String) Notifications returned by the PingDirectory Configuration API.
- `required_actions` (Set of Object) Required actions returned by the PingDirectory Configuration API. (see [below for nested schema](#nestedatt--required_actions))

<a id="nestedatt--required_actions"></a>
### Nested Schema for `required_actions`

Read-Only:

- `property` (String)
- `synopsis` (String)
- `type` (String)

## Import

Import is supported using the following syntax:

```shell
# "attributeMapperPluginId" should be the id of the Attribute Mapper Plugin to be imported
terraform import pingdirectory_attribute_mapper_plugin.myAttributeMapperPlugin attributeMapperPluginId
```
