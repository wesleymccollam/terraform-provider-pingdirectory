---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingdirectory_simple_connection_criteria Resource - terraform-provider-pingdirectory"
subcategory: ""
description: |-
  Manages a Simple Connection Criteria.
---

# pingdirectory_simple_connection_criteria (Resource)

Manages a Simple Connection Criteria.

## Example Usage

```terraform
terraform {
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
}

resource "pingdirectory_simple_connection_criteria" "mySimpleConnectionCriteria" {
  id             = "MySimpleConnectionCriteria"
  description    = "Simple connection example"
  user_auth_type = ["internal", "sasl"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Name of this object.

### Optional

- `all_included_user_filter` (Set of String) Specifies a search filter that must match the entry of the authenticated user for clients included in this Simple Connection Criteria. If any filters are provided, then all of those filters must match the authenticated user entry. This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections.
- `all_included_user_group_dn` (Set of String) Specifies the DN of a group in which authenticated users must exist for clients included in this Simple Connection Criteria. If any group DNs are provided, then the authenticated user must be a member of all of those groups. This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections. Refer to the authz version of this property in Simple Result Criteria if operations are being proxied (performed using proxied authorization), and you need to match the originating user of the operation rather than the proxy user (the user the proxy authenticated as).
- `all_included_user_privilege` (Set of String) Specifies the name of a privilege that must be held by the authenticated user for clients included in this Simple Connection Criteria. If any privilege names are provided, then the authenticated user must have all of those privileges. This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections.
- `any_included_user_filter` (Set of String) Specifies a search filter that may match the entry of the authenticated user for clients included in this Simple Connection Criteria. If any filters are provided, then at least one of those filters must match the authenticated user entry. This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections.
- `any_included_user_group_dn` (Set of String) Specifies the DN of a group in which authenticated users may exist for clients included in this Simple Connection Criteria. If any group DNs are provided, then the authenticated user must be a member of at least one of those groups. This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections. Refer to the authz version of this property in Simple Result Criteria if operations are being proxied (performed using proxied authorization), and you need to match the originating user of the operation rather than the proxy user (the user the proxy authenticated as).
- `any_included_user_privilege` (Set of String) Specifies the name of a privilege that may be held by the authenticated user for clients included in this Simple Connection Criteria. If any privilege names are provided, then the authenticated user must have at least one of those privileges. This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections.
- `authentication_security_level` (String) Indicates whether this Simple Connection Criteria should require or allow clients that authenticated using a secure manner. This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections.
- `communication_security_level` (String) Indicates whether this Simple Connection Criteria should require or allow clients using a secure communication channel.
- `description` (String) A description for this Connection Criteria
- `excluded_client_address` (Set of String) Specifies an address mask that may be used to specify a set of clients that should be excluded from this Simple Connection Criteria.
- `excluded_connection_handler` (Set of String) Specifies a connection handler for clients that should be excluded from this Simple Connection Criteria.
- `excluded_protocol` (Set of String) Specifies the name of a communication protocol that should be used by clients excluded from this Simple Connection Criteria.
- `excluded_user_base_dn` (Set of String) Specifies a base DN below which authenticated user entries may exist for clients excluded from this Simple Connection Criteria. This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections. Refer to the authz version of this property in Simple Result Criteria if operations are being proxied (performed using proxied authorization), and you need to match the originating user of the operation rather than the proxy user (the user the proxy authenticated as).
- `excluded_user_sasl_mechanism` (Set of String) Specifies the name of a SASL mechanism that should be used by clients excluded from this Simple Connection Criteria. This will only be taken into account for client connections that have authenticated to the server using a SASL mechanism and will be ignored for unauthenticated client connections and for client connections that authenticated using some other method (e.g., those performing simple or internal authentication).
- `included_client_address` (Set of String) Specifies an address mask that may be used to specify a set of clients that should be included in this Simple Connection Criteria.
- `included_connection_handler` (Set of String) Specifies a connection handler for clients that should be included in this Simple Connection Criteria.
- `included_protocol` (Set of String) Specifies the name of a communication protocol that should be used by clients included in this Simple Connection Criteria.
- `included_user_base_dn` (Set of String) Specifies a base DN below which authenticated user entries may exist for clients included in this Simple Connection Criteria. This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections. Refer to the authz version of this property in Simple Result Criteria if operations are being proxied (performed using proxied authorization), and you need to match the originating user of the operation rather than the proxy user (the user the proxy authenticated as).
- `included_user_sasl_mechanism` (Set of String) Specifies the name of a SASL mechanism that should be used by clients included in this Simple Connection Criteria. This will only be taken into account for client connections that have authenticated to the server using a SASL mechanism and will be ignored for unauthenticated client connections and for client connections that authenticated using some other method (e.g., those performing simple or internal authentication).
- `none_included_user_filter` (Set of String) Specifies a search filter that must not match the entry of the authenticated user for clients included in this Simple Connection Criteria. If any filters are provided, then none of those filters may match the authenticated user entry. This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections.
- `none_included_user_group_dn` (Set of String) Specifies the DN of a group in which authenticated users must not exist for clients included in this Simple Connection Criteria. If any group DNs are provided, then the authenticated user must not be a member any of those groups. This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections. Refer to the authz version of this property in Simple Result Criteria if operations are being proxied (performed using proxied authorization), and you need to match the originating user of the operation rather than the proxy user (the user the proxy authenticated as).
- `none_included_user_privilege` (Set of String) Specifies the name of a privilege that must not be held by the authenticated user for clients included in this Simple Connection Criteria. If any privilege names are provided, then the authenticated user must not have any of those privileges. This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections.
- `not_all_included_user_filter` (Set of String) Specifies a search filter that should not match the entry of the authenticated user for clients included in this Simple Connection Criteria. If any filters are provided, then at least one of those filters must not match the authenticated user entry (that is, the user entry may match zero or more of those filters, but not all of them). This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections.
- `not_all_included_user_group_dn` (Set of String) Specifies the DN of a group in which authenticated users should not exist for clients included in this Simple Connection Criteria. If any group DNs are provided, then the authenticated user must not be a member of at least one of those groups (that is, the user may be a member of zero or more of those groups, but not of all of them). This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections. Refer to the authz version of this property in Simple Result Criteria if operations are being proxied (performed using proxied authorization), and you need to match the originating user of the operation rather than the proxy user (the user the proxy authenticated as).
- `not_all_included_user_privilege` (Set of String) Specifies the name of a privilege that should not be held by the authenticated user for clients included in this Simple Connection Criteria. If any privilege names are provided, then the authenticated user must not have at least one of those privileges (that is, the user may hold zero or more of those privileges, but not all of them). This will only be taken into account for client connections that have authenticated to the server and will be ignored for unauthenticated client connections.
- `user_auth_type` (Set of String) Specifies the authentication types for client connections that may be included in this Simple Connection Criteria.

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
# "simpleConnectionCriteriaName" should be the name of the Simple Connection Criteria to be imported
terraform import pingdirectory_simple_connection_criteria simpleConnectionCriteriaName
```