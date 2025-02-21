---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingdirectory_syslog_text_error_log_publisher Resource - terraform-provider-pingdirectory"
subcategory: ""
description: |-
  Manages a Syslog Text Error Log Publisher.
---

# pingdirectory_syslog_text_error_log_publisher (Resource)

Manages a Syslog Text Error Log Publisher.

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

# Use "pingdirectory_default_syslog_text_error_log_publisher" if you are adopting existing configuration from the PingDirectory server into Terraform
resource "pingdirectory_syslog_text_error_log_publisher" "mySyslogTextErrorLogPublisher" {
  id                     = "MySyslogTextErrorLogPublisher"
  syslog_external_server = ["example.com"]
  enabled                = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `enabled` (Boolean) Indicates whether the Log Publisher is enabled for use.
- `id` (String) Name of this object.
- `syslog_external_server` (Set of String) The syslog server to which messages should be sent.

### Optional

- `default_severity` (Set of String) Specifies the default severity levels for the logger.
- `description` (String) A description for this Log Publisher
- `generify_message_strings_when_possible` (Boolean) Indicates whether to use the generified version of the log message string (which may use placeholders like %s for a string or %d for an integer), rather than the version of the message with those placeholders replaced with specific values that would normally be written to the log.
- `include_instance_name` (Boolean) Indicates whether log messages should include the instance name for the Directory Server.
- `include_product_name` (Boolean) Indicates whether log messages should include the product name for the Directory Server.
- `include_startup_id` (Boolean) Indicates whether log messages should include the startup ID for the Directory Server, which is a value assigned to the server instance at startup and may be used to identify when the server has been restarted.
- `include_thread_id` (Boolean) Indicates whether log messages should include the thread ID for the Directory Server in each log message. This ID can be used to correlate log messages from the same thread within a single log as well as generated by the same thread across different types of log files. More information about the thread with a specific ID can be obtained using the cn=JVM Stack Trace,cn=monitor entry.
- `logging_error_behavior` (String) Specifies the behavior that the server should exhibit if an error occurs during logging processing.
- `override_severity` (Set of String) Specifies the override severity levels for the logger based on the category of the messages.
- `queue_size` (Number) The maximum number of log records that can be stored in the asynchronous queue.
- `syslog_facility` (String) The syslog facility to use for the messages that are logged by this Syslog Text Error Log Publisher.
- `syslog_message_application_name` (String) The application name that will be included in syslog messages that are logged by this Syslog Text Error Log Publisher.
- `syslog_message_host_name` (String) The local host name that will be included in syslog messages that are logged by this Syslog Text Error Log Publisher.
- `syslog_severity` (String) The syslog severity to use for the messages that are logged by this Syslog Text Error Log Publisher. If this is not specified, then the severity for each syslog message will be automatically based on the severity for the associated log message.
- `timestamp_precision` (String) Specifies the smallest time unit to be included in timestamps.

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
# "syslogTextErrorLogPublisherId" should be the id of the Syslog Text Error Log Publisher to be imported
terraform import pingdirectory_syslog_text_error_log_publisher.mySyslogTextErrorLogPublisher syslogTextErrorLogPublisherId
```
