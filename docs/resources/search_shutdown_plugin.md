---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingdirectory_search_shutdown_plugin Resource - terraform-provider-pingdirectory"
subcategory: ""
description: |-
  Manages a Search Shutdown Plugin.
---

# pingdirectory_search_shutdown_plugin (Resource)

Manages a Search Shutdown Plugin.

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

# Use "pingdirectory_default_search_shutdown_plugin" if you are adopting existing configuration from the PingDirectory server into Terraform
resource "pingdirectory_search_shutdown_plugin" "mySearchShutdownPlugin" {
  id          = "MySearchShutdownPlugin"
  scope       = "base"
  filter      = "uid=user.1"
  output_file = "outlog"
  enabled     = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `enabled` (Boolean) Indicates whether the plug-in is enabled for use.
- `filter` (String) The filter to use for the search.
- `id` (String) Name of this object.
- `output_file` (String) The path of an LDIF file that should be created with the results of the search.
- `scope` (String) The scope to use for the search.

### Optional

- `base_dn` (String) The base DN to use for the search.
- `description` (String) A description for this Plugin
- `include_attribute` (Set of String) The name of an attribute that should be included in the results. This may include any token which is allowed as a requested attribute in search requests, including the name of an attribute, an asterisk (to indicate all user attributes), a plus sign (to indicate all operational attributes), an object class name preceded with an at symbol (to indicate all attributes associated with that object class), an attribute name preceded by a caret (to indicate that attribute should be excluded), or an object class name preceded by a caret and an at symbol (to indicate that all attributes associated with that object class should be excluded).
- `previous_file_extension` (String) An extension that should be appended to the name of an existing output file rather than deleting it. If a file already exists with the full previous file name, then it will be deleted before the current file is renamed to become the previous file.

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
# "searchShutdownPluginId" should be the id of the Search Shutdown Plugin to be imported
terraform import pingdirectory_search_shutdown_plugin.mySearchShutdownPlugin searchShutdownPluginId
```
