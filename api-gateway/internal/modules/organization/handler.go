package organization

import (
	"context"
	"errors"
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

// @Summary List Organization
// @Description List Organization
// @Tags Organization
// @Accept json
// @Produce json
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /organization [get]
func (h *OrganizationHandler) ListOrganizations(c *gin.Context) {
	ctx, exists := c.Get("grpc_ctx")
	if !exists {
		utils.RespondError(
			c,
			http.StatusUnauthorized,
			"UNAUTHORIZED",
			errors.New("missing context"),
			nil,
		)
		return
	}
	data, err := h.client.GetOrganizations(ctx.(context.Context))

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

func (h *OrganizationHandler) CreateDepartment(c *gin.Context) {

	ctx, exists := c.Get("grpc_ctx")
	if !exists {
		utils.RespondError(
			c,
			http.StatusUnauthorized,
			"UNAUTHORIZED",
			errors.New("missing context"),
			nil,
		)
		return
	}

	var data CreateDepartmentRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.RespondError(
			c,
			http.StatusBadRequest,
			"BAD_REQUEST",
			err,
			nil,
		)
		return
	}
	res, err := h.client.CreateDepartment(ctx.(context.Context), data)

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
