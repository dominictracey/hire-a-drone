package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"github.com/dominictracey/rugby-scores/models"
	"github.com/dominictracey/rugby-scores/restapi/operations"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name hire-a-drone --spec ../openApi.yaml --principal models.Principal

func configureFlags(api *operations.HireADroneAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.HireADroneAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	api.Logger = log.Printf

	api.Logger("Logger started for api")

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.FirebaseAuth = func(token string, scopes []string) (*models.Principal, error) {
		return nil, errors.NotImplemented("oauth2 bearer auth (firebase) has not yet been implemented")
	}

	// Applies when the "key" query is set
	api.APIKeyAuth = func(token string) (*models.Principal, error) {
		//api.Logger("Trying auth with key", token)
		if token != "" {
			prin := models.Principal(token)
			if api.Logger != nil {
				api.Logger("Trying auth with key %s", prin)
			}
			return &prin, nil
		}
		api.Logger("Access attempt with incorrect api key auth: %s", token)
		return nil, errors.New(401, "incorrect api key auth")
	}

	api.Logger("Key handler configured for api")

	api.AuthInfoFirebaseHandler = operations.AuthInfoFirebaseHandlerFunc(func(params operations.AuthInfoFirebaseParams, principal *models.Principal) middleware.Responder {
		return middleware.NotImplemented("operation .AuthInfoFirebase has not yet been implemented")
	})
	api.EchoHandler = operations.EchoHandlerFunc(func(params operations.EchoParams, principal *models.Principal) middleware.Responder {
		api.Logger("Trying echo %s", *params.Message)
		return operations.NewEchoOK().WithPayload(params.Message)
	})

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
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
