package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrRemoteTypeNotFound error = errors.New("cannot find 'type' field in request")
	ErrInvalidRemoteType  error = errors.New("cannot find 'type' field as string in request")
)

func (api *API) registerUser(ctx *gin.Context) {
	registerDetails := make(map[string]interface{})
	if err := bindMap(ctx, registerDetails); err != nil {
		log.Warnf("failed to bind json to github user [bad_request=%s]", err.Error())
		badRequestError(ctx, err)
		return
	}
	iRemoteType, ok := registerDetails["type"]
	if !ok {
		log.Warnf("failed to register user [bad_request=%s]", ErrRemoteTypeNotFound.Error())
		badRequestError(ctx, ErrRemoteTypeNotFound)
		return
	}

	remoteType, ok := iRemoteType.(string)
	if !ok {
		log.Warnf("failed to register user [bad_request=%s]", ErrInvalidRemoteType.Error())
		badRequestError(ctx, ErrInvalidRemoteType)
		return
	}

	user, err := api.service.RegisterUser(remoteType, registerDetails)
	if err != nil {
		log.Errorf("failed to register user [internal_error=%s]", err.Error())
		internalError(ctx, err)
		return
	}

	log.Infof("[user=%s] successfully registered with [remote=%s]", user.Name, remoteType)
	ctx.JSON(http.StatusAccepted, user)
}

func (api *API) getUser(ctx *gin.Context) {
	name := ctx.Param("name")

	user, err := api.service.GetUser(name)
	if err != nil {
		log.Errorf("failed to fetch user [internal_error=%s]", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, user)
}
