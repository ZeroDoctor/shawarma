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
	engine.Use(gin.Recovery(), gin.Logger())

	api.controllerV1(engine.Group("/v1"))

	return engine.Run(address...)
}

func (api *API) controllerV1(router *gin.RouterGroup) {
	router.POST("/register/user", api.registerUser)
	router.GET("/user/:name", api.getUser)
	router.GET("/user", api.getUserByState)
	router.GET("/repos", setHeaderAllowOrigin, userContext, api.getAllRepos)
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

// userContext checks if the shawarma_user cookie exists
func userContext(ctx *gin.Context) {
	cookie, err := ctx.Cookie("shawarma_user")
	if err != nil {
		log.Errorf("failed to find cookie [error=%s]", err.Error())
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	log.Debugf("Cookie Token: %+v", cookie)
	user, err := service.VerifyJWTToken(cookie)
	if err != nil {
		log.Errorf("failed to verify jwt token [error=%s]", err.Error())
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	ctx.Set("user", user)
}

func setHeaderAllowOrigin(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*") // TODO: change this to domain required
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type")
}
