---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingdirectory_delay_recurring_task Resource - terraform-provider-pingdirectory"
subcategory: ""
description: |-
  Manages a Delay Recurring Task.
---

# pingdirectory_delay_recurring_task (Resource)

Manages a Delay Recurring Task.

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

# Use "pingdirectory_default_delay_recurring_task" if you are adopting existing configuration from the PingDirectory server into Terraform
resource "pingdirectory_delay_recurring_task" "myDelayRecurringTask" {
  id = "MyDelayRecurringTask"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Name of this object.

### Optional

- `alert_on_failure` (Boolean) Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task fails to complete successfully.
- `alert_on_start` (Boolean) Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task starts running.
- `alert_on_success` (Boolean) Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task completes successfully.
- `cancel_on_task_dependency_failure` (Boolean) Indicates whether an instance of this Recurring Task should be canceled if the task immediately before it in the recurring task chain fails to complete successfully (including if it is canceled by an administrator before it starts or while it is running).
- `description` (String) A description for this Recurring Task
- `duration_to_wait_for_search_to_return_entries` (String) The maximum length of time that the server will continue to perform internal searches using the criteria from the ldap-url-for-search-expected-to-return-entries property.
- `duration_to_wait_for_work_queue_idle` (String) Indicates that task should wait for up to the specified length of time for the work queue to report that all worker threads are idle and there are no pending operations. Note that this primarily monitors operations that use worker threads, which does not include internal operations (for example, those invoked by extensions), and may not include requests from non-LDAP clients (for example, HTTP-based clients).
- `email_on_failure` (Set of String) The email addresses to which a message should be sent if an instance of this Recurring Task fails to complete successfully. If this option is used, then at least one smtp-server must be configured in the global configuration.
- `email_on_start` (Set of String) The email addresses to which a message should be sent whenever an instance of this Recurring Task starts running. If this option is used, then at least one smtp-server must be configured in the global configuration.
- `email_on_success` (Set of String) The email addresses to which a message should be sent whenever an instance of this Recurring Task completes successfully. If this option is used, then at least one smtp-server must be configured in the global configuration.
- `ldap_url_for_search_expected_to_return_entries` (Set of String) An LDAP URL that provides the criteria for a search request that is expected to return at least one entry. The search will be performed internally, and only the base DN, scope, and filter from the URL will be used; any host, port, or requested attributes included in the URL will be ignored.
- `search_interval` (String) The length of time the server should sleep between searches performed using the criteria from the ldap-url-for-search-expected-to-return-entries property.
- `search_time_limit` (String) The length of time that the server will wait for a response to each internal search performed using the criteria from the ldap-url-for-search-expected-to-return-entries property.
- `sleep_duration` (String) The length of time to sleep before the task completes.
- `task_return_state_if_timeout_is_encountered` (String) The return state to use if a timeout is encountered while waiting for the server work queue to become idle (if the duration-to-wait-for-work-queue-idle property has a value), or if the time specified by the duration-to-wait-for-search-to-return-entries elapses without the associated search returning any entries.

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
# "delayRecurringTaskId" should be the id of the Delay Recurring Task to be imported
terraform import pingdirectory_delay_recurring_task.myDelayRecurringTask delayRecurringTaskId
```
