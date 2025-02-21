---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingdirectory_quickstart_http_servlet_extension Resource - terraform-provider-pingdirectory"
subcategory: ""
description: |-
  Manages a Quickstart Http Servlet Extension.
---

# pingdirectory_quickstart_http_servlet_extension (Resource)

Manages a Quickstart Http Servlet Extension.

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

# Use "pingdirectory_default_quickstart_http_servlet_extension" if you are adopting existing configuration from the PingDirectory server into Terraform
resource "pingdirectory_quickstart_http_servlet_extension" "myQuickstartHttpServletExtension" {
  id          = "MyQuickstartHttpServletExtension"
  description = "Example Quickstart Http Servlet Extension"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Name of this object.

### Optional

- `correlation_id_response_header` (String) Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are "Correlation-Id", "X-Amzn-Trace-Id", and "X-Request-Id".
- `cross_origin_policy` (String) The cross-origin request policy to use for the HTTP Servlet Extension.
- `description` (String) A description for this HTTP Servlet Extension
- `response_header` (Set of String) Specifies HTTP header fields and values added to response headers for all requests.
- `server` (String) Specifies the PingFederate server to be configured.

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
# "quickstartHttpServletExtensionId" should be the id of the Quickstart Http Servlet Extension to be imported
terraform import pingdirectory_quickstart_http_servlet_extension.myQuickstartHttpServletExtension quickstartHttpServletExtensionId
```
