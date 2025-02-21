---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingdirectory_default_consent_service Resource - terraform-provider-pingdirectory"
subcategory: ""
description: |-
  Manages a Consent Service.
---

# pingdirectory_default_consent_service (Resource)

Manages a Consent Service.

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

resource "pingdirectory_default_consent_service" "myConsentService" {
  enabled                    = true
  base_dn                    = "ou=consents,dc=example,dc=com"
  bind_dn                    = "cn=consent service account"
  unprivileged_consent_scope = "urn:pingdirectory:consent"
  privileged_consent_scope   = "urn:pingdirectory:consent_admin"
  search_size_limit          = 90
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `audience` (String) A string or URI that identifies the Consent Service in the context of OAuth2 authorization.
- `base_dn` (String) The base DN under which consent records are stored.
- `bind_dn` (String) The DN of an internal service account used by the Consent Service to make internal LDAP requests.
- `consent_record_identity_mapper` (Set of String) If specified, the Identity Mapper(s) that may be used to map consent record subject and actor values to DNs. This is typically only needed if privileged API clients will be used.
- `enabled` (Boolean) Indicates whether the Consent Service is enabled.
- `privileged_consent_scope` (String) The name of a scope that must be present in an access token accepted by the Consent Service if the client is to be considered privileged.
- `search_size_limit` (Number) The maximum number of consent resources that may be returned from a search request.
- `service_account_dn` (Set of String) The set of account DNs that the Consent Service will consider to be privileged.
- `unprivileged_consent_scope` (String) The name of a scope that must be present in an access token accepted by the Consent Service for unprivileged clients.

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
# This resource is singleton, so the value of "id" doesn't matter - it is just a placeholder
terraform import pingdirectory_default_consent_service.myConsentService id
```
