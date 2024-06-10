// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreatePublicIpDetails The representation of CreatePublicIpDetails
type CreatePublicIpDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the public IP. For ephemeral public IPs,
	// you must set this to the private IP's compartment OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Defines when the public IP is deleted and released back to the Oracle Cloud
	// Infrastructure public IP pool. For more information, see
	// Public IP Addresses (https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingpublicIPs.htm).
	Lifetime CreatePublicIpDetailsLifetimeEnum `mandatory:"true" json:"lifetime"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private IP to assign the public IP to.
	// Required for an ephemeral public IP because it must always be assigned to a private IP
	// (specifically a *primary* private IP).
	// Optional for a reserved public IP. If you don't provide it, the public IP is created but not
	// assigned to a private IP. You can later assign the public IP with
	// UpdatePublicIp.
	PrivateIpId *string `mandatory:"false" json:"privateIpId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the public IP pool.
	PublicIpPoolId *string `mandatory:"false" json:"publicIpPoolId"`
}

func (m CreatePublicIpDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreatePublicIpDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreatePublicIpDetailsLifetimeEnum(string(m.Lifetime)); !ok && m.Lifetime != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Lifetime: %s. Supported values are: %s.", m.Lifetime, strings.Join(GetCreatePublicIpDetailsLifetimeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreatePublicIpDetailsLifetimeEnum Enum with underlying type: string
type CreatePublicIpDetailsLifetimeEnum string

// Set of constants representing the allowable values for CreatePublicIpDetailsLifetimeEnum
const (
	CreatePublicIpDetailsLifetimeEphemeral CreatePublicIpDetailsLifetimeEnum = "EPHEMERAL"
	CreatePublicIpDetailsLifetimeReserved  CreatePublicIpDetailsLifetimeEnum = "RESERVED"
)

var mappingCreatePublicIpDetailsLifetimeEnum = map[string]CreatePublicIpDetailsLifetimeEnum{
	"EPHEMERAL": CreatePublicIpDetailsLifetimeEphemeral,
	"RESERVED":  CreatePublicIpDetailsLifetimeReserved,
}

var mappingCreatePublicIpDetailsLifetimeEnumLowerCase = map[string]CreatePublicIpDetailsLifetimeEnum{
	"ephemeral": CreatePublicIpDetailsLifetimeEphemeral,
	"reserved":  CreatePublicIpDetailsLifetimeReserved,
}

// GetCreatePublicIpDetailsLifetimeEnumValues Enumerates the set of values for CreatePublicIpDetailsLifetimeEnum
func GetCreatePublicIpDetailsLifetimeEnumValues() []CreatePublicIpDetailsLifetimeEnum {
	values := make([]CreatePublicIpDetailsLifetimeEnum, 0)
	for _, v := range mappingCreatePublicIpDetailsLifetimeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreatePublicIpDetailsLifetimeEnumStringValues Enumerates the set of values in String for CreatePublicIpDetailsLifetimeEnum
func GetCreatePublicIpDetailsLifetimeEnumStringValues() []string {
	return []string{
		"EPHEMERAL",
		"RESERVED",
	}
}

// GetMappingCreatePublicIpDetailsLifetimeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreatePublicIpDetailsLifetimeEnum(val string) (CreatePublicIpDetailsLifetimeEnum, bool) {
	enum, ok := mappingCreatePublicIpDetailsLifetimeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
