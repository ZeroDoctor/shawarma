package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zerodoctor/shawarma/pkg/model"
)

func (api *API) registerUser(ctx *gin.Context) {
	var registerDetails map[string]interface{}
	if err := ctx.BindJSON(registerDetails); err != nil {
		log.Warnf("failed to bind json to github user [error=%s]", err.Error())
		badRequestError(ctx, err)
		return
	}

	user, err := api.gitRemote.RegisterUser(registerDetails)
	if err != nil {
		internalError(ctx, err)
		return
	}

	ctx.JSON(http.StatusAccepted, user)
}

func (api *API) branchUpdateEvent(ctx *gin.Context) {
	var branch model.Branch
	if err := ctx.BindJSON(&branch); err != nil {
		log.Warnf("failed to bind json to update branch event [error=%s]", err.Error())
		badRequestError(ctx, err)
		return
	}

	if _, err := api.db.SaveBranch(branch); err != nil {
		log.Errorf("failed to save branch [error=%s]", err.Error())
		internalError(ctx, err)
		return
	}

	ctx.JSON(http.StatusAccepted, branch)
}
