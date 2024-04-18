package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zerodoctor/shawarma/internal/model"
)

func (api *API) RegisterOrganization(ctx *gin.Context) {
	var org model.Organization
	if err := ctx.BindJSON(org); err != nil {
		InternalError(ctx, err)
		return
	}

	if _, err := api.db.InsertOrganization(org); err != nil {
		InternalError(ctx, err)
		return
	}

	ctx.JSON(http.StatusAccepted, org)
}
