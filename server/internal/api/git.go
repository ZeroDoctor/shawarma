package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zerodoctor/shawarma/internal/model"
)

func (api *API) registerGithubUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.BindJSON(&user); err != nil {
		log.Warnf("failed to bind json to github user [error=%s]", err.Error())
		badRequestError(ctx, err)
		return
	}

	var err error
	user, err = api.service.SaveGithubAuthUser(user)
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

	if _, err := api.db.InsertBranch(branch); err != nil {
		log.Errorf("failed to save branch [error=%s]", err.Error())
		internalError(ctx, err)
		return
	}

	ctx.JSON(http.StatusAccepted, branch)
}
