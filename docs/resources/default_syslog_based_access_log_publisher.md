---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingdirectory_default_syslog_based_access_log_publisher Resource - terraform-provider-pingdirectory"
subcategory: ""
description: |-
  Manages a Syslog Based Access Log Publisher.
---

# pingdirectory_default_syslog_based_access_log_publisher (Resource)

Manages a Syslog Based Access Log Publisher.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Name of this object.

### Optional

- `asynchronous` (Boolean) Indicates whether the Writer Based Access Log Publisher will publish records asynchronously.
- `auto_flush` (Boolean) Specifies whether to flush the writer after every log record.
- `connection_criteria` (String) Specifies a set of connection criteria that must match the associated client connection in order for a connect, disconnect, request, or result message to be logged.
- `correlate_requests_and_results` (Boolean) Indicates whether to automatically log result messages for any operation in which the corresponding request was logged. In such cases, the result, entry, and reference criteria will be ignored, although the log-responses, log-search-entries, and log-search-references properties will be honored.
- `description` (String) A description for this Log Publisher
- `enabled` (Boolean) Indicates whether the Syslog Based Access Log Publisher is enabled for use.
- `generify_message_strings_when_possible` (Boolean) Indicates whether to use generified version of certain message strings, including diagnostic messages, additional information messages, authentication failure reasons, and disconnect messages. Generified versions of those strings may use placeholders (like %s for a string or %d for an integer) rather than the version of the string with those placeholders replaced with specific values.
- `include_add_attribute_names` (Boolean) Indicates whether log messages for add requests should include a list of the names of the attributes included in the entry to add.
- `include_extended_search_request_details` (Boolean) Indicates whether log messages for search requests should include extended information from the request, including the requested size limit, time limit, alias dereferencing behavior, and types only behavior.
- `include_instance_name` (Boolean) Indicates whether log messages should include the instance name for the Directory Server.
- `include_modify_attribute_names` (Boolean) Indicates whether log messages for modify requests should include a list of the names of the attributes to be modified.
- `include_product_name` (Boolean) Indicates whether log messages should include the product name for the Directory Server.
- `include_replication_change_id` (Boolean) Indicates whether to log information about the replication change ID.
- `include_request_controls` (Boolean) Indicates whether log messages for operation requests should include a list of the OIDs of any controls included in the request.
- `include_request_details_in_intermediate_response_messages` (Boolean) Indicates whether log messages for intermediate responses should include information about the associated operation request.
- `include_request_details_in_result_messages` (Boolean) Indicates whether log messages for operation results should include information about both the request and the result.
- `include_request_details_in_search_entry_messages` (Boolean) Indicates whether log messages for search result entries should include information about the associated search request.
- `include_request_details_in_search_reference_messages` (Boolean) Indicates whether log messages for search result references should include information about the associated search request.
- `include_requester_dn` (Boolean) Indicates whether log messages for operation requests should include the DN of the authenticated user for the client connection on which the operation was requested.
- `include_requester_ip_address` (Boolean) Indicates whether log messages for operation requests should include the IP address of the client that requested the operation.
- `include_response_controls` (Boolean) Indicates whether log messages for operation results should include a list of the OIDs of any controls included in the result.
- `include_result_code_names` (Boolean) Indicates whether result log messages should include human-readable names for result codes in addition to their numeric values.
- `include_search_entry_attribute_names` (Boolean) Indicates whether log messages for search result entries should include a list of the names of the attributes included in the entry that was returned.
- `include_startup_id` (Boolean) Indicates whether log messages should include the startup ID for the Directory Server, which is a value assigned to the server instance at startup and may be used to identify when the server has been restarted.
- `include_thread_id` (Boolean) Indicates whether log messages should include the thread ID for the Directory Server in each log message. This ID can be used to correlate log messages from the same thread within a single log as well as generated by the same thread across different types of log files. More information about the thread with a specific ID can be obtained using the cn=JVM Stack Trace,cn=monitor entry.
- `log_assurance_completed` (Boolean) Indicates whether to log information about the result of replication assurance processing.
- `log_client_certificates` (Boolean) Indicates whether to log information about any client certificates presented to the server.
- `log_connects` (Boolean) Indicates whether to log information about connections established to the server.
- `log_disconnects` (Boolean) Indicates whether to log information about connections that have been closed by the client or terminated by the server.
- `log_field_behavior` (String) The behavior to use for determining which fields to log and whether to transform the values of those fields in any way.
- `log_intermediate_responses` (Boolean) Indicates whether to log information about intermediate responses sent to the client.
- `log_requests` (Boolean) Indicates whether to log information about requests received from clients.
- `log_results` (Boolean) Indicates whether to log information about the results of client requests.
- `log_search_entries` (Boolean) Indicates whether to log information about search result entries sent to the client.
- `log_search_references` (Boolean) Indicates whether to log information about search result references sent to the client.
- `log_security_negotiation` (Boolean) Indicates whether to log information about the result of any security negotiation (e.g., SSL handshake) processing that has been performed.
- `logging_error_behavior` (String) Specifies the behavior that the server should exhibit if an error occurs during logging processing.
- `max_string_length` (Number) Specifies the maximum number of characters that may be included in any string in a log message before that string is truncated and replaced with a placeholder indicating the number of characters that were omitted. This can help prevent extremely long log messages from being written.
- `queue_size` (Number) The maximum number of log records that can be stored in the asynchronous queue.
- `request_criteria` (String) Specifies a set of request criteria that must match the associated operation request in order for a request or result to be logged by this Access Log Publisher.
- `result_criteria` (String) Specifies a set of result criteria that must match the associated operation result in order for that result to be logged by this Access Log Publisher.
- `search_entry_criteria` (String) Specifies a set of search entry criteria that must match the associated search result entry in order for that it to be logged by this Access Log Publisher.
- `search_reference_criteria` (String) Specifies a set of search reference criteria that must match the associated search result reference in order for that it to be logged by this Access Log Publisher.
- `server_host_name` (String) Specifies the hostname or IP address of the syslogd host to log to. It is highly recommend to use localhost.
- `server_port` (Number) Specifies the port number of the syslogd host to log to.
- `suppress_internal_operations` (Boolean) Indicates whether internal operations (for example, operations that are initiated by plugins) should be logged along with the operations that are requested by users.
- `suppress_replication_operations` (Boolean) Indicates whether access messages that are generated by replication operations should be suppressed.
- `syslog_facility` (Number) Specifies the syslog facility to use for this Syslog Based Access Log Publisher

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

