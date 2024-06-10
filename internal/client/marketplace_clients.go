// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_marketplace "github.com/oracle/oci-go-sdk/v65/marketplace"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_marketplace.MarketplaceClient", &OracleClient{InitClientFn: initMarketplaceMarketplaceClient})
}

func initMarketplaceMarketplaceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_marketplace.NewMarketplaceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) MarketplaceClient() *oci_marketplace.MarketplaceClient {
	return m.GetClient("oci_marketplace.MarketplaceClient").(*oci_marketplace.MarketplaceClient)
}
