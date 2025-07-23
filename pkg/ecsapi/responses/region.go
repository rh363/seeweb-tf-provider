package response

import "github.com/rh363/terraform-provider-seeweb/pkg/ecsapi/models"

type GetRegionsResponse struct {
	Status  string          `json:"status"`
	Regions []models.Region `json:"regions"`
}
