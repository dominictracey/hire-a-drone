package restapi

import (
	"crypto/tls"
	"net/http"
	"sync"
	"sync/atomic"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	graceful "github.com/tylerb/graceful"

	"github.com/dominictracey/rugby-scores/models"
	"github.com/dominictracey/rugby-scores/restapi/operations"
	"github.com/dominictracey/rugby-scores/restapi/operations/pilots"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name hire-a-drone --spec ../openapi.yaml
var items = make(map[int64]*models.Pilot)
var lastID int64

var itemsLock = &sync.Mutex{}

func newItemID() int64 {
	return atomic.AddInt64(&lastID, 1)
}

func addItem(item *models.Pilot) error {
	if item == nil {
		return errors.New(500, "item must be present")
	}

	itemsLock.Lock()
	defer itemsLock.Unlock()

	newID := newItemID()
	item.ID = newID
	items[newID] = item

	return nil
}

func updateItem(id int64, item *models.Pilot) error {
	if item == nil {
		return errors.New(500, "item must be present")
	}

	itemsLock.Lock()
	defer itemsLock.Unlock()

	_, exists := items[id]
	if !exists {
		return errors.NotFound("not found: item %d", id)
	}

	item.ID = id
	items[id] = item
	return nil
}

func deleteItem(id int64) error {
	itemsLock.Lock()
	defer itemsLock.Unlock()

	_, exists := items[id]
	if !exists {
		return errors.NotFound("not found: item %d", id)
	}

	delete(items, id)
	return nil
}

func allItems(since int64, limit int32) (result []*models.Pilot) {
	result = make([]*models.Pilot, 0)
	for id, item := range items {
		if len(result) >= int(limit) {
			return
		}
		if since == 0 || id > since {
			result = append(result, item)
		}
	}
	return
}

func configureFlags(api *operations.HireADroneAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

//var lastID int64

func configureAPI(api *operations.HireADroneAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.GoogleIDTokenAuth = func(token string, scopes []string) (interface{}, error) {
		return nil, errors.NotImplemented("oauth2 bearer auth (google_id_token) has not yet been implemented")
	}

	api.FirebaseAuth = func(token string, scopes []string) (interface{}, error) {
		return nil, errors.NotImplemented("oauth2 bearer auth (firebase) has not yet been implemented")
	}

	api.GoogleJwtAuth = func(token string, scopes []string) (interface{}, error) {
		return nil, errors.NotImplemented("oauth2 bearer auth (google_jwt) has not yet been implemented")
	}

	api.Auth0JwkAuth = func(token string, scopes []string) (interface{}, error) {
		return nil, errors.NotImplemented("oauth2 bearer auth (auth0_jwk) has not yet been implemented")
	}

	// Applies when the "key" query is set
	api.APIKeyAuth = func(token string) (interface{}, error) {
		return nil, errors.NotImplemented("api key auth (api_key) key from query param [key] has not yet been implemented")
	}

	api.PilotsAddOneHandler = pilots.AddOneHandlerFunc(func(params pilots.AddOneParams) middleware.Responder {
		if err := addItem(params.Body); err != nil {
			return pilots.NewAddOneDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return pilots.NewAddOneCreated().WithPayload(params.Body)
	})
	api.AuthInfoFirebaseHandler = operations.AuthInfoFirebaseHandlerFunc(func(params operations.AuthInfoFirebaseParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation .AuthInfoFirebase has not yet been implemented")
	})
	api.AuthInfoGoogleIDTokenHandler = operations.AuthInfoGoogleIDTokenHandlerFunc(func(params operations.AuthInfoGoogleIDTokenParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation .AuthInfoGoogleIDToken has not yet been implemented")
	})
	api.AuthInfoAuth0JwkHandler = operations.AuthInfoAuth0JwkHandlerFunc(func(params operations.AuthInfoAuth0JwkParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation .AuthInfoAuth0Jwk has not yet been implemented")
	})
	api.AuthInfoGoogleJwtHandler = operations.AuthInfoGoogleJwtHandlerFunc(func(params operations.AuthInfoGoogleJwtParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation .AuthInfoGoogleJwt has not yet been implemented")
	})
	api.PilotsDestroyOneHandler = pilots.DestroyOneHandlerFunc(func(params pilots.DestroyOneParams) middleware.Responder {
		if err := deleteItem(params.ID); err != nil {
			return pilots.NewDestroyOneDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return pilots.NewDestroyOneNoContent()
	})
	api.EchoHandler = operations.EchoHandlerFunc(func(params operations.EchoParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation .Echo has not yet been implemented")
	})
	api.PilotsFindPilotsHandler = pilots.FindPilotsHandlerFunc(func(params pilots.FindPilotsParams) middleware.Responder {
		mergedParams := pilots.NewFindPilotsParams()
		mergedParams.Since = swag.Int64(0)
		if params.Since != nil {
			mergedParams.Since = params.Since
		}
		if params.Limit != nil {
			mergedParams.Limit = params.Limit
		}
		return pilots.NewFindPilotsOK().WithPayload(allItems(*mergedParams.Since, *mergedParams.Limit))
	})
	api.PilotsUpdateOneHandler = pilots.UpdateOneHandlerFunc(func(params pilots.UpdateOneParams) middleware.Responder {
		if err := updateItem(params.ID, params.Body); err != nil {
			return pilots.NewUpdateOneDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return pilots.NewUpdateOneOK().WithPayload(params.Body)
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
