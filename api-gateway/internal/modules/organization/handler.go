package organization

import (
	"context"
	"net/http"

	"github.com/IKHINtech/erp_api_gateway/internal/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
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
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
		return
	}
	md := metadata.New(map[string]string{"authorization": token})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	data, err := h.client.GetOrganizations(ctx)

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
