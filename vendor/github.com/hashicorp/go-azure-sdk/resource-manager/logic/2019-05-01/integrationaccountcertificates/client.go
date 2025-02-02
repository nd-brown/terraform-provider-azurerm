package integrationaccountcertificates

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationAccountCertificatesClient struct {
	Client *resourcemanager.Client
}

func NewIntegrationAccountCertificatesClientWithBaseURI(api environments.Api) (*IntegrationAccountCertificatesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "integrationaccountcertificates", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntegrationAccountCertificatesClient: %+v", err)
	}

	return &IntegrationAccountCertificatesClient{
		Client: client,
	}, nil
}
