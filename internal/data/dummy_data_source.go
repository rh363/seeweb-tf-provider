package data

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

// DummyDataSource is a placeholder Terraform data source.
type DummyDataSource struct{}

func (d *DummyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "dummy"
}

func (d *DummyDataSource) Schema(_ context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	// Data source schema will be defined here.
}

func (d *DummyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// Read implementation.
}

var _ datasource.DataSource = (*DummyDataSource)(nil)
