---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingdirectory_conjur_external_server Resource - terraform-provider-pingdirectory"
subcategory: ""
description: |-
  Manages a Conjur External Server.
---

# pingdirectory_conjur_external_server (Resource)

Manages a Conjur External Server.

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

# Use "pingdirectory_default_conjur_external_server" if you are adopting existing configuration from the PingDirectory server into Terraform
resource "pingdirectory_conjur_external_server" "myConjurExternalServer" {
  id                           = "MyConjurExternalServer"
  conjur_server_base_uri       = "example.com"
  conjur_authentication_method = "myAuthMethod"
  conjur_account_name          = "user"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `conjur_account_name` (String) The name of the account with which the desired secrets are associated.
- `conjur_authentication_method` (String) The mechanism used to authenticate to the Conjur server.
- `conjur_server_base_uri` (Set of String) The base URL needed to access the CyberArk Conjur server. The base URL should consist of the protocol ("http" or "https"), the server address (resolvable name or IP address), and the port number. For example, "https://conjur.example.com:8443/".
- `id` (String) Name of this object.

### Optional

- `description` (String) A description for this External Server
- `trust_store_file` (String) The path to a file containing the information needed to trust the certificate presented by the Conjur servers.
- `trust_store_pin` (String, Sensitive) The PIN needed to access the contents of the trust store. This is only required if a trust store file is required, and if that trust store requires a PIN to access its contents.
- `trust_store_type` (String) The store type for the specified trust store file. The value should likely be one of "JKS", "PKCS12", or "BCFKS".

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
# "conjurExternalServerId" should be the id of the Conjur External Server to be imported
terraform import pingdirectory_conjur_external_server.myConjurExternalServer conjurExternalServerId
```
