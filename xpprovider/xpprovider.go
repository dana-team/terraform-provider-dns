package xpprovider

import (
	"context"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	fwschema "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-dns/internal/provider"
)

// GetProvider returns new provider instances for both Terraform Plugin Framework provider of type provider.Provider
// and Terraform Plugin SDKv2 provider of type *schema.Provider
// provider
func GetProvider(_ context.Context) (fwprovider.Provider, *schema.Provider) {
	p := provider.New()
	fwProvider := provider.NewFrameworkProvider()

	return fwProvider, p
}

// GetProviderSchema returns the Terraform Plugin SDKv2 provider schema of the provider
func GetProviderSchema(_ context.Context) *schema.Provider {
	return provider.New()
}

// GetFrameworkProviderSchema returns the Terraform Plugin Framework provider schema of the provider
func GetFrameworkProviderSchema(ctx context.Context) (fwschema.Schema, error) {
	fwProvider, _ := GetProvider(ctx)

	schemaReq := fwprovider.SchemaRequest{}
	schemaResp := fwprovider.SchemaResponse{}
	fwProvider.Schema(ctx, schemaReq, &schemaResp)

	return schemaResp.Schema, nil
}
