---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingdirectory_default_trust_store_backend Resource - terraform-provider-pingdirectory"
subcategory: ""
description: |-
  Manages a Trust Store Backend.
---

# pingdirectory_default_trust_store_backend (Resource)

Manages a Trust Store Backend.

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

resource "pingdirectory_default_trust_store_backend" "myTrustStoreBackend" {
  enabled = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `backend_id` (String) Specifies a name to identify the associated backend.

### Optional

- `backup_file_permissions` (String) Specifies the permissions that should be applied to files and directories created by a backup of the backend.
- `base_dn` (Set of String) Specifies the base DN(s) for the data that the backend handles.
- `description` (String) A description for this Backend
- `enabled` (Boolean) Indicates whether the backend is enabled in the server.
- `notification_manager` (String) Specifies a notification manager for changes resulting from operations processed through this Backend
- `return_unavailable_when_disabled` (Boolean) Determines whether any LDAP operation that would use this Backend is to return UNAVAILABLE when this Backend is disabled.
- `set_degraded_alert_when_disabled` (Boolean) Determines whether the Directory Server enters a DEGRADED state (and sends a corresponding alert) when this Backend is disabled.
- `trust_store_file` (String) Specifies the path to the file that stores the trust information.
- `trust_store_pin` (String, Sensitive) Specifies the clear-text PIN needed to access the Trust Store Backend.
- `trust_store_pin_file` (String) Specifies the path to the text file whose only contents should be a single line containing the clear-text PIN needed to access the Trust Store Backend.
- `trust_store_pin_passphrase_provider` (String) The passphrase provider to use to obtain the clear-text PIN needed to access the Trust Store Backend.
- `trust_store_type` (String) Specifies the format for the data in the key store file.
- `writability_mode` (String) Specifies the behavior that the backend should use when processing write operations.

### Read-Only

- `id` (String) Placeholder name of this object required by Terraform.
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
# "trustStoreBackendId" should be the backend_id of the Trust Store Backend to be imported
terraform import pingdirectory_default_trust_store_backend.myTrustStoreBackend trustStoreBackendId
```
