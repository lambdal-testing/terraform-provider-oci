// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_apm "github.com/oracle/oci-go-sdk/v65/apmcontrolplane"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_apm.ApmDomainClient", &OracleClient{InitClientFn: initApmcontrolplaneApmDomainClient})
}

func initApmcontrolplaneApmDomainClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_apm.NewApmDomainClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) ApmDomainClient() *oci_apm.ApmDomainClient {
	return m.GetClient("oci_apm.ApmDomainClient").(*oci_apm.ApmDomainClient)
}
