package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mvrilo/go-redoc"
	ginredoc "github.com/mvrilo/go-redoc/gin"
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

//go:generate swag init

// @title			Shawarma API
// @version		1.0
// @description	Shawarma Server API
// @host			localhost:4000
// @BasePath		/
func (api *API) Run(ctx context.Context, address ...string) error {
	doc := redoc.Redoc{
		Title:       "Shawarma API",
		Description: "Shawarma Server API",
		SpecFile:    "./server/docs/swagger.yaml",
		SpecPath:    "/server/docs/swagger.yaml",
		DocsPath:    "./server/docs",
	}

	engine := gin.New()
	engine.Use(gin.Recovery(), gin.Logger())
	engine.Use(ginredoc.New(doc))

	engine.StaticFile("favicon.ico", "./server/resources/favicon.ico")
	engine.LoadHTMLGlob("./server/resources/*.html")
	engine.GET("/docs", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

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

// userContext checks if the shawarma_user cookie exists
func userContext(ctx *gin.Context) {
	cookie, err := ctx.Cookie("shawarma_user")
	if err != nil {
		log.Errorf("failed to find cookie [error=%s]", err.Error())
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

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
