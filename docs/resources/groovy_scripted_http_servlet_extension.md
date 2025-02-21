---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingdirectory_groovy_scripted_http_servlet_extension Resource - terraform-provider-pingdirectory"
subcategory: ""
description: |-
  Manages a Groovy Scripted Http Servlet Extension.
---

# pingdirectory_groovy_scripted_http_servlet_extension (Resource)

Manages a Groovy Scripted Http Servlet Extension.

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

# Use "pingdirectory_default_groovy_scripted_http_servlet_extension" if you are adopting existing configuration from the PingDirectory server into Terraform
resource "pingdirectory_groovy_scripted_http_servlet_extension" "myGroovyScriptedHttpServletExtension" {
  id           = "MyGroovyScriptedHttpServletExtension"
  script_class = "com.example.myscriptclass"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Name of this object.
- `script_class` (String) The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted HTTP Servlet Extension.

### Optional

- `correlation_id_response_header` (String) Specifies the name of the HTTP response header that will contain a correlation ID value. Example values are "Correlation-Id", "X-Amzn-Trace-Id", and "X-Request-Id".
- `cross_origin_policy` (String) The cross-origin request policy to use for the HTTP Servlet Extension.
- `description` (String) A description for this HTTP Servlet Extension
- `response_header` (Set of String) Specifies HTTP header fields and values added to response headers for all requests.
- `script_argument` (Set of String) The set of arguments used to customize the behavior for the Scripted HTTP Servlet Extension. Each configuration property should be given in the form 'name=value'.

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
# "groovyScriptedHttpServletExtensionId" should be the id of the Groovy Scripted Http Servlet Extension to be imported
terraform import pingdirectory_groovy_scripted_http_servlet_extension.myGroovyScriptedHttpServletExtension groovyScriptedHttpServletExtensionId
```
