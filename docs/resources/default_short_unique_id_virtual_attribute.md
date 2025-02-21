---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pingdirectory_default_short_unique_id_virtual_attribute Resource - terraform-provider-pingdirectory"
subcategory: ""
description: |-
  Manages a Short Unique Id Virtual Attribute.
---

# pingdirectory_default_short_unique_id_virtual_attribute (Resource)

Manages a Short Unique Id Virtual Attribute.

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

resource "pingdirectory_default_short_unique_id_virtual_attribute" "myShortUniqueIdVirtualAttribute" {
  id                        = "Short Unique ID"
  sequence_number_attribute = "ds-sequence-number"
  enabled                   = true
  attribute_type            = "ds-short-unique-id"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Name of this object.

### Optional

- `allow_index_conflicts` (Boolean) Indicates whether the server should allow creating or altering this virtual attribute definition even if it conflicts with one or more indexes defined in the server.
- `attribute_type` (String) Specifies the attribute type for the attribute whose values are to be dynamically assigned by the virtual attribute.
- `base_dn` (Set of String) Specifies the base DNs for the branches containing entries that are eligible to use this virtual attribute.
- `client_connection_policy` (Set of String) Specifies a set of client connection policies for which this Virtual Attribute should be generated. If this is undefined, then this Virtual Attribute will always be generated. If it is associated with one or more client connection policies, then this Virtual Attribute will be generated only for operations requested by clients assigned to one of those client connection policies.
- `description` (String) A description for this Virtual Attribute
- `enabled` (Boolean) Indicates whether the Virtual Attribute is enabled for use.
- `filter` (Set of String) Specifies the search filters to be applied against entries to determine if the virtual attribute is to be generated for those entries.
- `group_dn` (Set of String) Specifies the DNs of the groups whose members can be eligible to use this virtual attribute.
- `multiple_virtual_attribute_evaluation_order_index` (Number) Specifies the order in which virtual attribute definitions for the same attribute type will be evaluated when generating values for an entry.
- `multiple_virtual_attribute_merge_behavior` (String) Specifies the behavior that will be exhibited for cases in which multiple virtual attribute definitions apply to the same multivalued attribute type. This will be ignored for single-valued attribute types.
- `sequence_number_attribute` (String) Specifies the name or OID of the attribute which contains the sequence number from which unique identifiers are generated. The attribute should have Integer syntax or a String syntax permitting integer values. If this property is modified then the filter property should be updated accordingly so that only entries containing the sequence number attribute are eligible to have a value generated for this virtual attribute.

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
# "shortUniqueIdVirtualAttributeId" should be the id of the Short Unique Id Virtual Attribute to be imported
terraform import pingdirectory_default_short_unique_id_virtual_attribute.myShortUniqueIdVirtualAttribute shortUniqueIdVirtualAttributeId
```
