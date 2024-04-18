package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zerodoctor/shawarma/internal/model"
)

func (api *API) registerOrganization(ctx *gin.Context) {
	var org model.Organization
	if err := ctx.BindJSON(org); err != nil {
		badRequestError(ctx, err)
		return
	}

	if _, err := api.db.InsertOrganization(org); err != nil {
		internalError(ctx, err)
		return
	}

	ctx.JSON(http.StatusAccepted, org)
}

func (api *API) branchUpdateEvent(ctx *gin.Context) {
	var branch model.Branch
	if err := ctx.BindJSON(branch); err != nil {
		badRequestError(ctx, err)
		return
	}

	if _, err := api.db.InsertBranch(branch); err != nil {
		internalError(ctx, err)
		return
	}
}
