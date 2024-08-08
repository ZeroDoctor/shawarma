package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zerodoctor/shawarma/internal/db"
	"github.com/zerodoctor/shawarma/internal/logger"
	"github.com/zerodoctor/shawarma/pkg/service"
)

var log *logrus.Logger = logger.Log

type API struct {
	db      db.DB
	service *service.Service
}

func NewAPI(db db.DB) *API {
	return &API{
		db:      db,
		service: service.NewService(db),
	}
}

func (api *API) Run(ctx context.Context, address ...string) error {
	engine := gin.New()

	api.controllerV1(engine.Group("/v1"))

	return engine.Run(address...)
}

func (api *API) controllerV1(router *gin.RouterGroup) {
	router.POST("/register/user", api.registerUser)
	router.GET("/user/:name", api.getUser)
	// router.POST("/register/runner", api.registerRunner)
	// router.POST("/event/branch", api.branchUpdateEvent)
	// router.POST("/pipeline/webhook", api.webhookPipeline)

	// NOTE: private apis
	// router.PUT("/internal/pipeline/status", api.setPipelineStatus)
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

func bindMap(ctx *gin.Context, m map[string]interface{}) error {
	data, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return err
	}
	defer ctx.Request.Body.Close()

	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	return nil
}
