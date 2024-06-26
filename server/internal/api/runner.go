package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zerodoctor/shawarma/pkg/model"
)

func (api *API) registerRunner(ctx *gin.Context) {
	var runner model.Runner
	if err := ctx.BindJSON(runner); err != nil {
		badRequestError(ctx, err)
		return
	}

	if _, err := api.db.SaveRunner(runner); err != nil {
		internalError(ctx, err)
		return
	}

	ctx.JSON(http.StatusAccepted, runner)
}
