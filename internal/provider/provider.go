package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/seeweb/seeweb-tf-provider/internal/data"
	"github.com/seeweb/seeweb-tf-provider/internal/resource"
)

// New returns a new instance of the Seeweb provider.
func New() provider.Provider {
	return &seewebProvider{}
}

type seewebProvider struct{}

func (p *seewebProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "seeweb"
	resp.Version = "0.1.0"
}

func (p *seewebProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	// No provider-level configuration yet.
}

func (p *seewebProvider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {
	// No-op configuration for now.
}

func (p *seewebProvider) Resources(_ context.Context) []func() provider.Resource {
	return []func() provider.Resource{
		func() provider.Resource { return &resource.DummyResource{} },
	}
}

func (p *seewebProvider) DataSources(_ context.Context) []func() provider.DataSource {
	return []func() provider.DataSource{
		func() provider.DataSource { return &data.DummyDataSource{} },
	}
}

var _ provider.Provider = (*seewebProvider)(nil)

var _ providerserver.ProviderServer = &providerserver.SimpleProviderServer{Provider: New()}
