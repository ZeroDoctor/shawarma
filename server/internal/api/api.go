package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zerodoctor/shawarma/internal/db"
)

type API struct {
	db db.DB
}

func (api *API) Run(ctx context.Context, address ...string) {
	engine := gin.New()

	router := engine.Group("/v1")
	api.controllerV1(router)

	engine.Run(address...)
}

func (api *API) controllerV1(router *gin.RouterGroup) {
	router.POST("/register/org", api.registerOrganization)
	router.POST("/register/runner", api.registerRunner)

	router.POST("/event/branch", api.branchUpdateEvent)

	router.POST("/pipeline/webhook", api.webhookPipeline)

	// NOTE: private apis

	router.PUT("/internal/pipeline/status", api.setPipelineStatus)
}

func internalError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}

func badRequestError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}
