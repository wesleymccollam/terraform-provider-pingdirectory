---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingdirectory_http_proxy_external_server Resource - terraform-provider-pingdirectory"
subcategory: ""
description: |-
  Manages a Http Proxy External Server. Supported in PingDirectory product version 9.2.0.0+.
---

# pingdirectory_http_proxy_external_server (Resource)

Manages a Http Proxy External Server. Supported in PingDirectory product version 9.2.0.0+.

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

resource "pingdirectory_http_proxy_external_server" "myHttpProxyExternalServer" {
  id               = "MyHttpProxyExternalServer"
  server_host_name = "example.com"
  server_port      = 1234
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Name of this object.
- `server_host_name` (String) The host name or IP address of the HTTP Proxy External Server.
- `server_port` (Number) The port on which the HTTP Proxy External Server is listening for connections.

### Optional

- `basic_authentication_passphrase_provider` (String) A passphrase provider that provides access to the password to use to authenticate to the HTTP Proxy External Server.
- `basic_authentication_username` (String) The username to use to authenticate to the HTTP Proxy External Server.
- `description` (String) A description for this External Server

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
# "httpProxyExternalServerId" should be the id of the Http Proxy External Server to be imported
terraform import pingdirectory_http_proxy_external_server.myHttpProxyExternalServer httpProxyExternalServerId
```
