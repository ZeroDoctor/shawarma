package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zerodoctor/shawarma/pkg/model"
)

// @Summary		Get all repositories
// @Description	Retrieve all users' repositories
// @Tags			git
// @Accept			json
// @Produce		json
// @Success		200	{object}	[]model.Repository
// @Failure		500	{object}	map[string]interface{}
// @Router			/v1/repos    [get]
func (api *API) getAllRepos(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		badRequestError(ctx, ErrUserNotFound)
	}

	repos, err := api.db.GetAllUserRepos(user.(model.User))
	if err != nil {
		log.Errorf("failed to get all repos [error=%s]", err.Error())
		internalError(ctx, err)
		return
	}

	log.Infof("found %d repos", len(repos))
	ctx.JSON(http.StatusOK, repos)
}

func (api *API) getBranches(ctx *gin.Context) {
	repoID := ctx.Param(":repoID")
	log.Infof("get all branches [repoID=%s]", repoID)

	// TODO: check if user has access to repo

	repoUUID, err := uuid.Parse(repoID)
	if err != nil {
		log.Errorf("failed to parse repoID [error=%s]", err.Error())
		internalError(ctx, err)
		return
	}

	branches, err := api.db.GetBranches(model.UUID(repoUUID))
	if err != nil {
		log.Errorf("failed to get branches [error=%s]", err.Error())
		internalError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, branches)
}

func (api *API) branchUpdateEvent(ctx *gin.Context) {
	var branch model.Branch
	if err := ctx.BindJSON(&branch); err != nil {
		log.Warnf("failed to bind json to update branch event [error=%s]", err.Error())
		badRequestError(ctx, err)
		return
	}

	// TODO: check if user has access to repo

	if _, err := api.db.SaveBranch(branch); err != nil {
		log.Errorf("failed to save branch [error=%s]", err.Error())
		internalError(ctx, err)
		return
	}

	ctx.JSON(http.StatusAccepted, branch)
}
