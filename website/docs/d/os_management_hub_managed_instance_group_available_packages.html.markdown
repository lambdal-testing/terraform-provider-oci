---
subcategory: "Os Management Hub"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_os_management_hub_managed_instance_group_available_packages"
sidebar_current: "docs-oci-datasource-os_management_hub-managed_instance_group_available_packages"
description: |-
  Provides the list of Managed Instance Group Available Packages in Oracle Cloud Infrastructure Os Management Hub service
---

# Data Source: oci_os_management_hub_managed_instance_group_available_packages
This data source provides the list of Managed Instance Group Available Packages in Oracle Cloud Infrastructure Os Management Hub service.

Lists available packages on the specified managed instances group. Filter the list against a variety 
of criteria including but not limited to the package name.


## Example Usage

```hcl
data "oci_os_management_hub_managed_instance_group_available_packages" "test_managed_instance_group_available_packages" {
	#Required
	managed_instance_group_id = oci_os_management_hub_managed_instance_group.test_managed_instance_group.id

	#Optional
	compartment_id = var.compartment_id
	display_name = var.managed_instance_group_available_package_display_name
	display_name_contains = var.managed_instance_group_available_package_display_name_contains
	is_latest = var.managed_instance_group_available_package_is_latest
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment that contains the resources to list. This filter returns only resources contained within the specified compartment.
* `display_name` - (Optional) A filter to return resources that match the given display names.
* `display_name_contains` - (Optional) A filter to return resources that may partially match the given display name.
* `is_latest` - (Optional) Indicates whether to list only the latest versions of packages, module streams, and stream profiles.
* `managed_instance_group_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance group.


## Attributes Reference

The following attributes are exported:

* `managed_instance_group_available_package_collection` - The list of managed_instance_group_available_package_collection.

### ManagedInstanceGroupAvailablePackage Reference

The following attributes are exported:

* `items` - List of available packages.
	* `architecture` - The architecture for which this package was built.
	* `display_name` - Package name.
	* `is_latest` - Indicates whether this is the latest package version.
	* `name` - Unique identifier for the package. Note that this is not an OCID.
	* `software_sources` - List of software sources that provide the software package.
		* `description` - Software source description.
		* `display_name` - Software source name.
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the software source.
		* `is_mandatory_for_autonomous_linux` - Indicates whether this is a required software source for Autonomous Linux instances. If true, the user can't unselect it.
		* `software_source_type` - Type of the software source.
	* `type` - Type of the package.
	* `version` - Version of the available package.

