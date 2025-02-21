---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingdirectory_admin_alert_account_status_notification_handler Resource - terraform-provider-pingdirectory"
subcategory: ""
description: |-
  Manages a Admin Alert Account Status Notification Handler.
---

# pingdirectory_admin_alert_account_status_notification_handler (Resource)

Manages a Admin Alert Account Status Notification Handler.

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

# Use "pingdirectory_default_admin_alert_account_status_notification_handler" if you are adopting existing configuration from the PingDirectory server into Terraform
resource "pingdirectory_admin_alert_account_status_notification_handler" "myAdminAlertAccountStatusNotificationHandler" {
  id                               = "MyAdminAlertAccountStatusNotificationHandler"
  account_status_notification_type = ["account-created"]
  enabled                          = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_status_notification_type` (Set of String) The types of account status notifications that should result in administrative alerts.
- `enabled` (Boolean) Indicates whether the Account Status Notification Handler is enabled. Only enabled handlers are invoked whenever a related event occurs in the server.
- `id` (String) Name of this object.

### Optional

- `account_creation_notification_request_criteria` (String) A request criteria object that identifies which add requests should result in account creation notifications for this handler.
- `account_update_notification_request_criteria` (String) A request criteria object that identifies which modify and modify DN requests should result in account update notifications for this handler.
- `asynchronous` (Boolean) Indicates whether the server should attempt to invoke this Account Status Notification Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver a message) will not delay processing for the operation that triggered the notification.
- `description` (String) A description for this Account Status Notification Handler

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
# "adminAlertAccountStatusNotificationHandlerId" should be the id of the Admin Alert Account Status Notification Handler to be imported
terraform import pingdirectory_admin_alert_account_status_notification_handler.myAdminAlertAccountStatusNotificationHandler adminAlertAccountStatusNotificationHandlerId
```
