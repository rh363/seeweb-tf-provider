package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/rh363/terraform-provider-seeweb/pkg/seewebapi"
)

var (
	_ datasource.DataSource              = &regionsDataSource{}
	_ datasource.DataSourceWithConfigure = &regionsDataSource{}
)

func NewRegionsDataSource() datasource.DataSource {
	return &regionsDataSource{}
}

type regionsDataSource struct {
	client *seewebapi.SeewebApi
}

func (r *regionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_regions"
}

func (r *regionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"regions": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"location":    schema.StringAttribute{Computed: true},
						"description": schema.StringAttribute{Computed: true},
					},
				},
			},
		},
	}
}

func (r *regionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state regionDataDourceModel
	tflog.Debug(ctx, "Fetching api regions")
	regions, err := r.client.Ecs().GetRegions(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Regions",
			err.Error())
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("Fetched %d regions by api", len(regions)))

	for _, region := range regions {
		state.Regions = append(state.Regions, regionModel{
			Location:    types.StringValue(region.Location),
			Description: types.StringValue(region.Description),
		})
	}
	tflog.Debug(ctx, "Finished Reading Regions")

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Finished Setting Regions State")
}

func (r *regionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*seewebapi.SeewebApi)
	if !ok {
		resp.Diagnostics.AddError(
			"Invalid Datasource Provider Data ",
			fmt.Sprintf("Expected *seewebapi.SeewebApi, got: %T. Please report this issue to the ", req.ProviderData),
		)
		return
	}

	r.client = client
}

type regionDataDourceModel struct {
	Regions []regionModel `tfsdk:"regions"`
}

type regionModel struct {
	Location    types.String `tfsdk:"location"`
	Description types.String `tfsdk:"description"`
}
