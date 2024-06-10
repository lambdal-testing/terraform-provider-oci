resource "oci_service_mesh_virtual_deployment" "test_virtual_deployment" {
  #Required
  compartment_id = var.compartment_ocid

  #Optional
  listeners {
    #Required
    port     = local.virtual_deployment_listeners_port
    protocol = local.virtual_deployment_listeners_protocol
    #Optional
    request_timeout_in_ms = local.virtual_deployment_listeners_request_timeout_in_ms
    idle_timeout_in_ms    = local.virtual_deployment_listeners_idle_timeout_in_ms
  }
  name = local.virtual_deployment_name
  service_discovery {
    #Required if type is DNS.
    hostname = local.virtual_deployment_service_discovery_hostname
    type     = local.virtual_deployment_service_discovery_type
  }
  virtual_service_id = oci_service_mesh_virtual_service.test_virtual_service.id
  access_logging {

    #Optional
    is_enabled = local.virtual_deployment_access_logging_is_enabled
  }
  description   = local.virtual_deployment_description
  freeform_tags = { "bar-key" = "value" }
}

data "oci_service_mesh_virtual_deployment" "test_virtual_deployment" {
  #Required
  virtual_deployment_id = oci_service_mesh_virtual_deployment.test_virtual_deployment.id
}

locals {
  virtual_deployment_name                            = "test-virtual-deployment"
  virtual_deployment_description                     = "test virtual deployment description"
  virtual_deployment_listeners_port                  = "8080"
  virtual_deployment_listeners_request_timeout_in_ms = 100
  virtual_deployment_listeners_idle_timeout_in_ms    = 500
  virtual_deployment_listeners_protocol              = "HTTP"
  // allowed values for above for above are ("HTTP", "TLS_PASSTHROUGH", "TCP", "HTTP2", "GRPC")

  virtual_deployment_service_discovery_hostname = "test.com"
  virtual_deployment_service_discovery_type     = "DNS"
  // Allowed values for above are ("DNS", "DISABLED"). Default is "DISABLED"
  virtual_deployment_access_logging_is_enabled  = true
}
