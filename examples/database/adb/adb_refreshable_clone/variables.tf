// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

variable "tenancy_ocid" {
}

variable "user_ocid" {
}

variable "fingerprint" {
}

variable "private_key_path" {
}

variable "region" {
}

variable "compartment_id" {
  default = "ocid1.tenancy.oc1..aaaaaaaazxdmffivtoe32kvio5e2dcgz24re5rqbkis3452yi2e7tc3x2erq"
}

variable "autonomous_database_backup_display_name" {
  default = "Monthly Backup"
}

variable "autonomous_database_db_workload" {
  default = "OLTP"
}

variable "autonomous_data_warehouse_db_workload" {
  default = "DW"
}

variable "autonomous_database_defined_tags_value" {
  default = "value"
}

variable "autonomous_database_freeform_tags" {
  default = {
    "Department" = "Finance"
  }
}

variable "autonomous_database_license_model" {
  default = "LICENSE_INCLUDED"
}

variable "autonomous_database_is_dedicated" {
  default = false
}

variable "autonomous_database_autonomous_maintenance_schedule_type" {
  default = "EARLY"
}

variable "autonomous_database_character_set" {
  default = "AL32UTF8"
}

variable "autonomous_database_ncharacter_set" {
  default = "AL16UTF16"
}

variable "autonomous_database_character_set_character_set_type" {
  default = "DATABASE"
}

variable "autonomous_database_character_set_is_shared" {
  default = true

}

variable "autonomous_database_character_set_is_dedicated" {
  default = false
}