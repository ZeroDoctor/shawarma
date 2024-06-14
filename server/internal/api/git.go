package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zerodoctor/shawarma/internal/model"
	"github.com/zerodoctor/shawarma/internal/service"
)

func (api *API) registerGithubUser(ctx *gin.Context) {
	var githubUser model.GithubUser
	if err := ctx.BindJSON(&githubUser); err != nil {
		log.Warnf("failed to bind json to github user [error=%s]", err.Error())
		badRequestError(ctx, err)
		return
	}

	token, err := service.GetGithubToken(githubUser)
	if err != nil {
		log.Errorf("failed to fetch github token [error=%s]", err.Error())
		internalError(ctx, err)
		return
	}
	githubUser.Token = token

	user, err := api.db.InsertGithubUser(githubUser)
	if err != nil {
		log.Errorf("failed to save user [error=%s]", err.Error())
		internalError(ctx, err)
		return
	}
	user.GithubToken = ""

	ctx.JSON(http.StatusAccepted, user)
}

func (api *API) registerOrganization(ctx *gin.Context) {
	var org model.Organization
	if err := ctx.BindJSON(&org); err != nil {
		log.Warnf("failed to bind json to register org [error=%s]", err.Error())
		badRequestError(ctx, err)
		return
	}

	if _, err := api.db.InsertOrganization(org); err != nil {
		log.Errorf("failed to save org [error=%s]", err.Error())
		internalError(ctx, err)
		return
	}

	ctx.JSON(http.StatusAccepted, org)
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
