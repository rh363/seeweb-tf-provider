package resource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// DummyResource is a placeholder Terraform resource.
type DummyResource struct{}

func (r *DummyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "dummy"
}

func (r *DummyResource) Schema(_ context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	// Resource schema will be defined here.
}

func (r *DummyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Create implementation.
}

func (r *DummyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Read implementation.
}

func (r *DummyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Update implementation.
}

func (r *DummyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Delete implementation.
}

var _ resource.Resource = (*DummyResource)(nil)
