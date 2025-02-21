---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingdirectory_simple_request_criteria Resource - terraform-provider-pingdirectory"
subcategory: ""
description: |-
  Manages a Simple Request Criteria.
---

# pingdirectory_simple_request_criteria (Resource)

Manages a Simple Request Criteria.

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

# Use "pingdirectory_default_simple_request_criteria" if you are adopting existing configuration from the PingDirectory server into Terraform
resource "pingdirectory_simple_request_criteria" "mySimpleRequestCriteria" {
  id = "MySimpleRequestCriteria"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Name of this object.

### Optional

- `all_included_request_control` (Set of String) Specifies the OID of a control that must be present in the request from the client for operations included in this Simple Request Criteria. If any control OIDs are provided, then the request must contain all of those controls.
- `all_included_target_entry_filter` (Set of String) Specifies a search filter that must match the target entry for requests included in this Simple Request Criteria. This will only be taken into account for add, simple bind, compare, delete, modify, modify DN, and search operations. It will be ignored for abandon, SASL bind, extended, and unbind operations. If any filters are provided, then the target entry must match all of those filters.
- `all_included_target_entry_group_dn` (Set of String) Specifies the DN of a group in which the user associated with the target entry must be a member for requests included in this Simple Request Criteria. This will only be taken into account for add, simple bind, compare, delete, modify, modify DN, and search operations. It will be ignored for abandon, SASL bind, extended, and unbind operations. If any group DNs are provided, then the target entry must be a member of all of those groups.
- `any_included_request_control` (Set of String) Specifies the OID of a control that may be present in the request from the client for operations included in this Simple Request Criteria. If any control OIDs are provided, then the request must contain at least one of those controls.
- `any_included_target_entry_filter` (Set of String) Specifies a search filter that may match the target entry for requests included in this Simple Request Criteria. This will only be taken into account for add, simple bind, compare, delete, modify, modify DN, and search operations. It will be ignored for abandon, SASL bind, extended, and unbind operations. If any filters are provided, then the target entry must match at least one of those filters.
- `any_included_target_entry_group_dn` (Set of String) Specifies the DN of a group in which the user associated with the target entry may be a member for requests included in this Simple Request Criteria. This will only be taken into account for add, simple bind, compare, delete, modify, modify DN, and search operations. It will be ignored for abandon, SASL bind, extended, and unbind operations. If any group DNs are provided, then the target entry must be a member of at least one of those groups.
- `connection_criteria` (String) Specifies a connection criteria object that must match the associated client connection for operations included in this Simple Request Criteria.
- `description` (String) A description for this Request Criteria
- `excluded_application_name` (Set of String) Specifies an application name for requests excluded from this Simple Request Criteria.
- `excluded_extended_operation_oid` (Set of String) Specifies the request OID for extended requests excluded from this Simple Request Criteria. This will only be taken into account for extended requests and will be ignored for all other types of requests.
- `excluded_target_attribute` (Set of String) Specifies the name or OID of an attribute type which must not be targeted by requests included in this Simple Request Criteria. This will only be taken into account for add, compare, modify, modify DN, and search operations. It will be ignored for abandon, bind, delete, extended, and unbind operations.
- `excluded_target_entry_dn` (Set of String) Specifies a base DN below which targeted entries may not exist for requests included in this Simple Request Criteria. This will only be taken into account for add, simple bind, compare, delete, modify, modify DN, and search operations. It will be ignored for abandon, SASL bind, extended, and unbind operations.
- `excluded_target_sasl_mechanism` (Set of String) Specifies the name of a SASL mechanism for bind requests excluded from this Simple Request Criteria. This will only be taken into account for SASL bind operations and will be ignored for other types of operations and for bind operations that do not use SASL authentication.
- `included_application_name` (Set of String) Specifies an application name for requests included in this Simple Request Criteria.
- `included_extended_operation_oid` (Set of String) Specifies the request OID for extended requests included in this Simple Request Criteria. This will only be taken into account for extended requests and will be ignored for all other types of requests.
- `included_search_scope` (Set of String) Specifies the search scope values included in this Simple Request Criteria. This will only be taken into account for search requests and will be ignored for all other types of requests.
- `included_target_attribute` (Set of String) Specifies the name or OID of an attribute type which must be targeted by requests included in this Simple Request Criteria. This will only be taken into account for add, compare, modify, modify DN, and search operations. It will be ignored for abandon, bind, delete, extended, and unbind operations.
- `included_target_entry_dn` (Set of String) Specifies a base DN below which targeted entries may exist for requests included in this Simple Request Criteria. This will only be taken into account for add, simple bind, compare, delete, modify, modify DN, and search operations. It will be ignored for abandon, SASL bind, extended, and unbind operations.
- `included_target_sasl_mechanism` (Set of String) Specifies the name of a SASL mechanism for bind requests included in this Simple Request Criteria. This will only be taken into account for SASL bind operations and will be ignored for other types of operations and for bind operations that do not use SASL authentication.
- `none_included_request_control` (Set of String) Specifies the OID of a control that must not be present in the request from the client for operations included in this Simple Request Criteria. If any control OIDs are provided, then the request must not contain any of those controls.
- `none_included_target_entry_filter` (Set of String) Specifies a search filter that must not match the target entry for requests included in this Simple Request Criteria. This will only be taken into account for add, simple bind, compare, delete, modify, modify DN, and search operations. It will be ignored for abandon, SASL bind, extended, and unbind operations. If any filters are provided, then the target entry must not match any of those filters.
- `none_included_target_entry_group_dn` (Set of String) Specifies the DN of a group in which the user associated with the target entry must not be a member for requests included in this Simple Request Criteria. This will only be taken into account for add, simple bind, compare, delete, modify, modify DN, and search operations. It will be ignored for abandon, SASL bind, extended, and unbind operations. If any group DNs are provided, then the target entry must not be a member of any of those groups.
- `not_all_included_request_control` (Set of String) Specifies the OID of a control that should not be present in the request from the client for operations included in this Simple Request Criteria. If any control OIDs are provided, then the request must not contain at least one of those controls (that is, the request may contain zero or more of those controls, but not all of them).
- `not_all_included_target_entry_filter` (Set of String) Specifies a search filter that should not match the target entry for requests included in this Simple Request Criteria. This will only be taken into account for add, simple bind, compare, delete, modify, modify DN, and search operations. It will be ignored for abandon, SASL bind, extended, and unbind operations. If any filters are provided, then the target entry must not match at least one of those filters (that is, the request may match zero or more of those filters, but not of all of them).
- `not_all_included_target_entry_group_dn` (Set of String) Specifies the DN of a group in which the user associated with the target entry should not be a member for requests included in this Simple Request Criteria. This will only be taken into account for add, simple bind, compare, delete, modify, modify DN, and search operations. It will be ignored for abandon, SASL bind, extended, and unbind operations. If any group DNs are provided, then the target entry must not be a member of at least one of those groups (that is, the target entry may be a member of zero or more of those groups, but not all of them).
- `operation_origin` (Set of String) Specifies the origin for operations to be included in this Simple Request Criteria. If no values are provided, then the operation origin will not be taken into consideration when determining whether an operation matches this Simple Request Criteria.
- `operation_type` (Set of String) Specifies the operation type(s) for operations that should be included in this Simple Request Criteria.
- `target_bind_type` (Set of String) Specifies the authentication type for bind requests included in this Simple Request Criteria. This will only be taken into account for bind operations and will be ignored for any other type of operation. If no values are provided, then the authentication type will not be considered when determining whether the request should be included in this Simple Request Criteria.
- `using_administrative_session_worker_thread` (String) Indicates whether operations being processed using a dedicated administrative session worker thread should be included in this Simple Request Criteria.

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
# "simpleRequestCriteriaId" should be the id of the Simple Request Criteria to be imported
terraform import pingdirectory_simple_request_criteria.mySimpleRequestCriteria simpleRequestCriteriaId
```
