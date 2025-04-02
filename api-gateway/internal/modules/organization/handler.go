package organization

import (
	"context"
	"errors"
	"net/http"

	"github.com/IKHINtech/erp_api_gateway/internal/utils"
	"github.com/gin-gonic/gin"

	orgClient "github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/client"
	organizationv1 "github.com/IKHINtech/erp_microservice_go_grpc/organization-microservice/pb/v1"
)

type OrganizationHandler struct {
	client *orgClient.OrganizationClient
}

func NewOrganizationHandler(client *orgClient.OrganizationClient) *OrganizationHandler {
	return &OrganizationHandler{client: client}
}

// @Summary Get Organization
// @Description Get Organization
// @Tags Organization
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Organization ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /org/organization/{id} [get]
func (h *OrganizationHandler) GetOrganization(c *gin.Context) {
	// baca param :id
	id := c.Param("id")

	if id == "" {
		// jika id tidak ada
		utils.RespondError(c, http.StatusBadRequest, "BAD_REQUEST", errors.New("id is required"), nil)
	}

	data := organizationv1.GetOrganizationRequest{Id: id}
	res, err := h.client.GetOrganization(c, &data)

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

// @Summary Create Organization
// @Description Create Organization
// @Tags Organization
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body CreateOrganizationRequest true "Organization"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /org/organization [post]
func (h *OrganizationHandler) CreateOrganization(c *gin.Context) {

	// cek konteks
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
	data := new(CreateOrganizationRequest)

	if err := c.ShouldBindJSON(data); err != nil {
		utils.RespondError(
			c,
			http.StatusBadRequest,
			"BAD_REQUEST",
			err,
			nil,
		)
		return
	}

	payload := organizationv1.CreateOrganizationRequest{
		Name:        data.Name,
		Description: data.Description,
	}
	res, err := h.client.CreateOrganization(ctx.(context.Context), &payload)
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

func (h *OrganizationHandler) UpdateOrganization(c *gin.Context) {
}

func (h *OrganizationHandler) DeleteOrganization(c *gin.Context) {
}

// @Summary List Organization
// @Description List Organization
// @Tags Organization
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /org/organization [get]
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
	data, err := h.client.ListOrganization(ctx.(context.Context))

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

// @Summary Create Departments
// @Description Create Department
// @Tags Organization
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} utils.SuccessResponse
// @Param data body CreateDepartmentRequest true "Department"
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /org/department [post]
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

	// buat payload
	payload := organizationv1.CreateDepartmentRequest{
		Name:        data.Name,
		Description: data.Description,
	}

	res, err := h.client.CreateDepartment(ctx.(context.Context), &payload)

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

// @Summary List Departments
// @Description List Departments
// @Tags Organization
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /org/department [get]
func (h *OrganizationHandler) GetDepartments(c *gin.Context) {
	// cek konteks
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

	res, err := h.client.ListDepartment(ctx.(context.Context))
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
