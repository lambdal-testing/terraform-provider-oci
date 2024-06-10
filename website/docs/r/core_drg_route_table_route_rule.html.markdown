---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_drg_route_table_route_rule"
sidebar_current: "docs-oci-resource-core-drg_route_table_route_rule"
description: |-
  Provides the Drg Route Table Route Rule resource in Oracle Cloud Infrastructure Core service
---

# oci_core_drg_route_table_route_rule
This resource provides the Drg Route Table Route Rule resource in Oracle Cloud Infrastructure Core service.

Adds one static route rule to the specified DRG route table.


## Example Usage

```hcl
resource "oci_core_drg_route_table_route_rule" "test_drg_route_table_route_rule" {
	#Required
	drg_route_table_id = oci_core_drg_route_table.test_drg_route_table.id
	destination = var.drg_route_table_route_rule_route_rules_destination
	destination_type = var.drg_route_table_route_rule_route_rules_destination_type
	next_hop_drg_attachment_id = oci_core_drg_attachment.test_drg_attachment.id

}
```

## Argument Reference

The following arguments are supported:

* `drg_route_table_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DRG route table.

		Potential values:
		* IP address range in CIDR notation. This can be an IPv4 CIDR block or IPv6 prefix. For example: `192.168.1.0/24` or `2001:0db8:0123:45::/56`. 
	* `destination_type` - (Required) Type of destination for the rule. Allowed values:
		* `CIDR_BLOCK`: If the rule's `destination` is an IP address range in CIDR notation. 
	* `next_hop_drg_attachment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the next hop DRG attachment. The next hop DRG attachment is responsible for reaching the network destination. 
* `destination` - (Required) (Updatable) This is the range of IP addresses used for matching when routing traffic. Only CIDR_BLOCK values are allowed.

	Potential values:
	* IP address range in CIDR notation. This can be an IPv4 or IPv6 CIDR. For example: `192.168.1.0/24` or `2001:0db8:0123:45::/56`.
* `destination_type` - (Required) Type of destination for the rule. Required if `direction` = `EGRESS`. Allowed values:
	* `CIDR_BLOCK`: If the rule's `destination` is an IP address range in CIDR notation. 
* `next_hop_drg_attachment_id` - (Required)  The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the next hop DRG attachment. The next hop DRG attachment is responsible for reaching the network destination.

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `attributes` - Additional properties for the route, computed by the service.
* `destination` - Represents the range of IP addresses to match against when routing traffic.

	Potential values:
	* An IP address range (IPv4 or IPv6) in CIDR notation. For example: `192.168.1.0/24` or `2001:0db8:0123:45::/56`.
	* When you're setting up a security rule for traffic destined for a particular `Service` through a service gateway, this is the `cidrBlock` value associated with that [Service](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Service/). For example: `oci-phx-objectstorage`. 
* `destination_type` - The type of destination for the rule.

	Allowed values:
	* `CIDR_BLOCK`: If the rule's `destination` is an IP address range in CIDR notation.
	* `SERVICE_CIDR_BLOCK`: If the rule's `destination` is the `cidrBlock` value for a [Service](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Service/) (the rule is for traffic destined for a particular `Service` through a service gateway). 
* `id` - The Oracle-assigned ID of the DRG route rule. 
* `is_blackhole` - Indicates that if the next hop attachment does not exist, so traffic for this route is discarded without notification. 
* `is_conflict` - Indicates that the route was not imported due to a conflict between route rules. 
* `next_hop_drg_attachment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the next hop DRG attachment responsible for reaching the network destination.

	A value of `BLACKHOLE` means traffic for this route is discarded without notification. 
* `route_provenance` - The earliest origin of a route. If a route is advertised to a DRG through an IPsec tunnel attachment, and is propagated to peered DRGs via RPC attachments, the route's provenance in the peered DRGs remains `IPSEC_TUNNEL`, because that is the earliest origin.

	No routes with a provenance `IPSEC_TUNNEL` or `VIRTUAL_CIRCUIT` will be exported to IPsec tunnel or virtual circuit attachments, regardless of the attachment's export distribution. 
* `route_type` - You can specify static routes for the DRG route table using the API. The DRG learns dynamic routes from the DRG attachments using various routing protocols. 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Drg Route Table Route Rule
	* `update` - (Defaults to 20 minutes), when updating the Drg Route Table Route Rule
	* `delete` - (Defaults to 20 minutes), when destroying the Drg Route Table Route Rule


## Import

DrgRouteTableRouteRule can be imported using the `id`, e.g.

```
$ terraform import oci_core_drg_route_table_route_rule.test_drg_route_table_route_rule "drgRouteTables/{drgRouteTableId}/routeRules/{id}" 
```

