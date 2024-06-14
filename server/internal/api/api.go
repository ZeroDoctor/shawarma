package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zerodoctor/shawarma/internal/db"
	"github.com/zerodoctor/shawarma/internal/logger"
)

var log *logrus.Logger = logger.Log

type API struct {
	db db.DB
}

func NewAPI(db db.DB) *API {
	return &API{
		db: db,
	}
}

func (api *API) Run(ctx context.Context, address ...string) error {
	engine := gin.New()

	router := engine.Group("/v1")
	api.controllerV1(router)

	return engine.Run(address...)
}

func (api *API) controllerV1(router *gin.RouterGroup) {
	router.POST("/register/github/user", api.registerGithubUser)
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
