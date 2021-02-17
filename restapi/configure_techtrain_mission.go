// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/GuiltyMorishita/techtrain-mission/restapi/operations"
	"github.com/GuiltyMorishita/techtrain-mission/restapi/operations/character"
	"github.com/GuiltyMorishita/techtrain-mission/restapi/operations/gacha"
	"github.com/GuiltyMorishita/techtrain-mission/restapi/operations/user"
)

//go:generate swagger generate server --target ../../techtrain-mission --name TechtrainMission --spec ../api-document.yaml --principal interface{}

func configureFlags(api *operations.TechtrainMissionAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TechtrainMissionAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.CharacterGetCharacterListHandler == nil {
		api.CharacterGetCharacterListHandler = character.GetCharacterListHandlerFunc(func(params character.GetCharacterListParams) middleware.Responder {
			return middleware.NotImplemented("operation character.GetCharacterList has not yet been implemented")
		})
	}
	if api.UserGetUserGetHandler == nil {
		api.UserGetUserGetHandler = user.GetUserGetHandlerFunc(func(params user.GetUserGetParams) middleware.Responder {
			return middleware.NotImplemented("operation user.GetUserGet has not yet been implemented")
		})
	}
	if api.GachaPostGachaDrawHandler == nil {
		api.GachaPostGachaDrawHandler = gacha.PostGachaDrawHandlerFunc(func(params gacha.PostGachaDrawParams) middleware.Responder {
			return middleware.NotImplemented("operation gacha.PostGachaDraw has not yet been implemented")
		})
	}
	if api.UserPostUserCreateHandler == nil {
		api.UserPostUserCreateHandler = user.PostUserCreateHandlerFunc(func(params user.PostUserCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation user.PostUserCreate has not yet been implemented")
		})
	}
	if api.UserPutUserUpdateHandler == nil {
		api.UserPutUserUpdateHandler = user.PutUserUpdateHandlerFunc(func(params user.PutUserUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation user.PutUserUpdate has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
