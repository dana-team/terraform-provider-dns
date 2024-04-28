package xpprovider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-dns/internal/provider"
)

func GetProvider(_ context.Context) (*schema.Provider, error) {
	return provider.New(), nil
}
