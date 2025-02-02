package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zerodoctor/shawarma/pkg/service"
)

// TODO: allow this to be set through env var
const JWT_EXPIRATION_TIME = 72 * time.Hour

var (
	ErrRemoteTypeNotFound  error = errors.New("cannot find 'type' field in request")
	ErrInvalidRemoteType   error = errors.New("cannot find 'type' field as string in request")
	ErrRemoteStateNotFound error = errors.New("cannot find 'state' field in request")
	ErrInvalidRemoteState  error = errors.New("cannot find 'state' field as string in request")
)

func (api *API) registerUser(ctx *gin.Context) {
	registerDetails := make(map[string]interface{})
	if err := bindMap(ctx, registerDetails); err != nil {
		log.Warnf("failed to bind json to github user [bad_request=%s]", err.Error())
		badRequestError(ctx, err)
		return
	}

	remoteType, remoteState := getRemoteDetails(ctx, registerDetails)
	if remoteType == "" || remoteState == "" {
		return
	}

	user, err := api.service.RegisterUser(remoteType, registerDetails)
	if err != nil {
		log.Errorf("failed to register user [internal_error=%s]", err.Error())
		internalError(ctx, err)
		return
	}

	token, err := service.CreateJWTToken(user, JWT_EXPIRATION_TIME)
	if err != nil {
		log.Errorf("failed to create jwt token [internal_error=%s]", err.Error())
		internalError(ctx, err)
		return
	}
	api.service.LocalCacheMap[remoteType+remoteState] = token

	log.Infof("[user=%s] successfully registered with [remote=%s]", user.Name, remoteType)
	ctx.JSON(http.StatusAccepted, gin.H{
		"token": token,
	})
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

func (api *API) getUserByState(ctx *gin.Context) {
	remoteType := ctx.Query("type")
	remoteState := ctx.Query("state")

	token, ok := api.service.LocalCacheMap[remoteType+remoteState]
	if !ok {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	delete(api.service.LocalCacheMap, remoteType+remoteState)

	ctx.JSON(http.StatusAccepted, gin.H{
		"token": token,
	})
}

func getRemoteDetails(ctx *gin.Context, details map[string]interface{}) (string, string) {
	iRemoteType, ok := details["type"]
	if !ok {
		log.Warnf("failed to register user [bad_request=%s]", ErrRemoteTypeNotFound.Error())
		badRequestError(ctx, ErrRemoteTypeNotFound)
		return "", ""
	}

	remoteType, ok := iRemoteType.(string)
	if !ok {
		log.Warnf("failed to register user [bad_request=%s]", ErrInvalidRemoteType.Error())
		badRequestError(ctx, ErrInvalidRemoteType)
		return "", ""
	}

	iRemoteState, ok := details["state"]
	if !ok {
		log.Warnf("failed to register user [bad_request=%s]", ErrRemoteStateNotFound.Error())
		badRequestError(ctx, ErrRemoteStateNotFound)
		return "", ""
	}

	remoteState, ok := iRemoteState.(string)
	if !ok {
		log.Warnf("failed to register user [bad_request=%s]", ErrInvalidRemoteState.Error())
		badRequestError(ctx, ErrInvalidRemoteState)
		return "", ""
	}

	return remoteType, remoteState
}
