---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_vm_cluster_recommended_network"
sidebar_current: "docs-oci-datasource-database-vm_cluster_recommended_network"
description: |-
  Provides details about a specific Vm Cluster Recommended Network in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_vm_cluster_recommended_network
This data source provides details about a specific Vm Cluster Recommended Network resource in Oracle Cloud Infrastructure Database service.

Generates a recommended Cloud@Customer VM cluster network configuration.


## Example Usage

```hcl
data "oci_database_vm_cluster_recommended_network" "test_vm_cluster_recommended_network" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.vm_cluster_recommended_network_display_name
	exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
	networks {
		#Required
		cidr = var.vm_cluster_recommended_network_networks_cidr
		domain = var.vm_cluster_recommended_network_networks_domain
		gateway = var.vm_cluster_recommended_network_networks_gateway
		netmask = var.vm_cluster_recommended_network_networks_netmask
		network_type = var.vm_cluster_recommended_network_networks_network_type
		prefix = var.vm_cluster_recommended_network_networks_prefix
		vlan_id = oci_core_vlan.test_vlan.id
	}

	#Optional
	db_servers = var.vm_cluster_recommended_network_db_servers
	defined_tags = map(oci_identity_tag_namespace.tag-namespace1.name.oci_identity_tag.tag1.name, var.vm_cluster_recommended_network_defined_tags_value)
	dns = var.vm_cluster_recommended_network_dns
	dr_scan_listener_port_tcp = var.vm_cluster_recommended_network_dr_scan_listener_port_tcp
	freeform_tags = var.vm_cluster_recommended_network_freeform_tags
	ntp = var.vm_cluster_recommended_network_ntp
	scan_listener_port_tcp = var.vm_cluster_recommended_network_scan_listener_port_tcp
	scan_listener_port_tcp_ssl = var.vm_cluster_recommended_network_scan_listener_port_tcp_ssl
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `db_servers` - (Optional) The list of Db server Ids to configure network.
* `defined_tags` - (Optional) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Required) The user-friendly name for the VM cluster network. The name does not need to be unique.
* `dns` - (Optional) The list of DNS server IP addresses. Maximum of 3 allowed.
* `dr_scan_listener_port_tcp` - (Optional) The DR SCAN TCPIP port. Default is 1521.
* `exadata_infrastructure_id` - (Required) The Exadata infrastructure [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `freeform_tags` - (Optional) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `networks` - (Required) List of parameters for generation of the client and backup networks.
	* `cidr` - (Required) The cidr for the network.
	* `domain` - (Required) The network domain name.
	* `gateway` - (Required) The network gateway.
	* `netmask` - (Required) The network netmask.
	* `network_type` - (Required) The network type.
	* `prefix` - (Required) The network domain name.
	* `vlan_id` - (Required) The network VLAN ID.
* `ntp` - (Optional) The list of NTP server IP addresses. Maximum of 3 allowed.
* `scan_listener_port_tcp` - (Optional) The SCAN TCPIP port. Default is 1521.
* `scan_listener_port_tcp_ssl` - (Optional) The SCAN TCPIP SSL port. Default is 2484.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the Exadata Cloud@Customer VM cluster network. The name does not need to be unique.
* `dns` - The list of DNS server IP addresses. Maximum of 3 allowed.
* `dr_scans` - The SCAN details for DR network
	* `hostname` - The Disaster recovery SCAN hostname.
	* `ips` - The list of Disaster recovery SCAN IP addresses. Three addresses should be provided.
	* `scan_listener_port_tcp` - The Disaster recovery SCAN TCPIP port. Default is 1521.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `ntp` - The list of NTP server IP addresses. Maximum of 3 allowed.
* `scans` - The SCAN details.
	* `hostname` - The SCAN hostname.
	* `ips` - The list of SCAN IP addresses. Three addresses should be provided.
	* `port` - **Deprecated.** This field is deprecated. You may use 'scanListenerPortTcp' to specify the port. The SCAN TCPIP port. Default is 1521. 
	* `scan_listener_port_tcp` - The SCAN TCPIP port. Default is 1521.
	* `scan_listener_port_tcp_ssl` - The SCAN TCPIP SSL port. Default is 2484.
* `vm_networks` - Details of the client and backup networks.
	* `domain_name` - The network domain name.
	* `gateway` - The network gateway.
	* `netmask` - The network netmask.
	* `network_type` - The network type.
	* `nodes` - The list of node details.
		* `db_server_id` - The Db server associated with the node.
		* `hostname` - The node host name.
		* `ip` - The node IP address.
		* `state` - The current state of the VM cluster network nodes. CREATING - The resource is being created REQUIRES_VALIDATION - The resource is created and may not be usable until it is validated. VALIDATING - The resource is being validated and not available to use. VALIDATED - The resource is validated and is available for consumption by VM cluster. VALIDATION_FAILED - The resource validation has failed and might require user input to be corrected. UPDATING - The resource is being updated and not available to use. ALLOCATED - The resource is currently being used by VM cluster. TERMINATING - The resource is being deleted and not available to use. TERMINATED - The resource is deleted and unavailable. FAILED - The resource is in a failed state due to validation or other errors. 
		* `vip` - The node virtual IP (VIP) address.
		* `vip_hostname` - The node virtual IP (VIP) host name.
	* `vlan_id` - The network VLAN ID.

