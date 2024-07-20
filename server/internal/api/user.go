package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrRemoteTypeNotFound error = errors.New("failed to find 'type' field in request")
	ErrInvalidRemoteType  error = errors.New("failed to find 'type' field as string in request")
)

func (api *API) registerUser(ctx *gin.Context) {
	registerDetails := make(map[string]interface{})
	if err := bindMap(ctx, registerDetails); err != nil {
		log.Warnf("failed to bind json to github user [error=%s]", err.Error())
		badRequestError(ctx, err)
		return
	}
	iRemoteType, ok := registerDetails["type"]
	if !ok {
		log.Warnf(ErrRemoteTypeNotFound.Error())
		badRequestError(ctx, ErrRemoteTypeNotFound)
	}

	remoteType, ok := iRemoteType.(string)
	if !ok {
		log.Warnf(ErrInvalidRemoteType.Error())
		badRequestError(ctx, ErrInvalidRemoteType)
	}

	user, err := api.service.RegisterUser(remoteType, registerDetails)
	if err != nil {
		internalError(ctx, err)
		return
	}

	user, err = api.db.SaveUser(user)
	if err != nil {
		internalError(ctx, err)
		return
	}

	ctx.JSON(http.StatusAccepted, user)
}

func (api *API) getUser(ctx *gin.Context) {
	name := ctx.Param("name")

	user, err := api.service.GetUser(name)
	if err != nil {
		internalError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
