---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingdirectory_third_party_plugin Resource - terraform-provider-pingdirectory"
subcategory: ""
description: |-
  Manages a Third Party Plugin.
---

# pingdirectory_third_party_plugin (Resource)

Manages a Third Party Plugin.

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

# Use "pingdirectory_default_third_party_plugin" if you are adopting existing configuration from the PingDirectory server into Terraform
resource "pingdirectory_third_party_plugin" "myThirdPartyPlugin" {
  id              = "MyThirdPartyPlugin"
  extension_class = "com.example.myclass"
  enabled         = false
  plugin_type     = ["postoperationadd"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `enabled` (Boolean) Indicates whether the plug-in is enabled for use.
- `extension_class` (String) The fully-qualified name of the Java class providing the logic for the Third Party Plugin.
- `id` (String) Name of this object.
- `plugin_type` (Set of String) Specifies the set of plug-in types for the plug-in, which specifies the times at which the plug-in is invoked.

### Optional

- `description` (String) A description for this Plugin
- `extension_argument` (Set of String) The set of arguments used to customize the behavior for the Third Party Plugin. Each configuration property should be given in the form 'name=value'.
- `invoke_for_internal_operations` (Boolean) Indicates whether the plug-in should be invoked for internal operations.
- `request_criteria` (String) Specifies a set of request criteria that may be used to indicate that this Third Party Plugin should only be invoked for operations in which the associated request matches this criteria.

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
# "thirdPartyPluginId" should be the id of the Third Party Plugin to be imported
terraform import pingdirectory_third_party_plugin.myThirdPartyPlugin thirdPartyPluginId
```
