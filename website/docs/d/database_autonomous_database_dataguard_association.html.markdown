---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_database_dataguard_association"
sidebar_current: "docs-oci-datasource-database-autonomous_database_dataguard_association"
description: |-
  Provides details about a specific Autonomous Database Dataguard Association in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_database_dataguard_association
This data source provides details about a specific Autonomous Database Dataguard Association resource in Oracle Cloud Infrastructure Database service.

Gets an Autonomous Database dataguard assocation for the specified Autonomous Database.


## Example Usage

```hcl
data "oci_database_autonomous_database_dataguard_association" "test_autonomous_database_dataguard_association" {
	#Required
	autonomous_database_dataguard_association_id = oci_database_autonomous_database_dataguard_association.test_autonomous_database_dataguard_association.id
	autonomous_database_id = oci_database_autonomous_database.test_autonomous_database.id
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_database_dataguard_association_id` - (Required) The Autonomous Database Dataguard Association [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `autonomous_database_id` - (Required) The database [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `apply_lag` - The lag time between updates to the primary database and application of the redo data on the standby database, as computed by the reporting database.  Example: `9 seconds` 
* `apply_rate` - The rate at which redo logs are synced between the associated databases.  Example: `180 Mb per second` 
* `autonomous_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database that has a relationship with the peer Autonomous Database. 
* `id` - The OCID of the Autonomous Dataguard created for Autonomous Container Database where given Autonomous Database resides in.
* `lifecycle_details` - Additional information about the current lifecycleState, if available. 
* `peer_autonomous_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peer Autonomous Database. 
* `peer_autonomous_database_life_cycle_state` - The current state of the Autonomous Dataguard.
* `peer_role` - The role of the Autonomous Dataguard enabled Autonomous Container Database.
* `protection_mode` - The protection mode of this Data Guard association. For more information, see [Oracle Data Guard Protection Modes](http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000) in the Oracle Data Guard documentation. 
* `role` - The role of the Autonomous Dataguard enabled Autonomous Container Database.
* `state` - The current state of the Autonomous Dataguard.
* `time_created` - The date and time the Data Guard association was created.
* `time_last_role_changed` - The date and time when the last role change action happened.

