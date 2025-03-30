package organization

import (
	"net/http"

	"github.com/IKHINtech/erp_api_gateway/internal/utils"
	"github.com/gin-gonic/gin"
)

type OrganizationHandler struct {
	client *OrganizationClient
}

func NewOrganizationHandler(client *OrganizationClient) *OrganizationHandler {
	return &OrganizationHandler{client: client}
}

func (h *OrganizationHandler) GetOrganization(c *gin.Context) {
	var data GetOrganizationRequest
	res, err := h.client.GetOrganization(c, data.ID.String())

	if err != nil {
		utils.RespondError(
			c,
			http.StatusBadRequest,
			"BAD_REQUEST",
			err,
			nil,
		)
		return
	}

	utils.RespondSuccess(c, http.StatusOK, res, nil)

}

func (h *OrganizationHandler) CreateOrganization(c *gin.Context) {
}

func (h *OrganizationHandler) UpdateOrganization(c *gin.Context) {
}

func (h *OrganizationHandler) DeleteOrganization(c *gin.Context) {
}

func (h *OrganizationHandler) ListOrganizations(c *gin.Context) {
	data, err := h.client.GetOrganizations(c)

	if err != nil {
		utils.RespondError(
			c,
			http.StatusBadRequest,
			"BAD_REQUEST",
			err,
			nil,
		)
		return
	}

	utils.RespondSuccess(c, http.StatusOK, data, nil)

}
