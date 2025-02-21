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

# Use "pingdirectory_default_ping_one_pass_through_authentication_plugin" if you are adopting existing configuration from the PingDirectory server into Terraform
resource "pingdirectory_ping_one_pass_through_authentication_plugin" "myPingOnePassThroughAuthenticationPlugin" {
  id                             = "MyPingOnePassThroughAuthenticationPlugin"
  api_url                        = "example.com/api"
  auth_url                       = "example.com/auth"
  oauth_client_id                = "example1"
  environment_id                 = "example2"
  user_mapping_local_attribute   = "cn"
  user_mapping_remote_json_field = "cn"
  enabled                        = false
}
