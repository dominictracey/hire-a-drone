package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/dominictracey/rugby-scores/restapi/operations/pilots"
)

// NewHireADroneAPI creates a new HireADrone instance
func NewHireADroneAPI(spec *loads.Document) *HireADroneAPI {
	return &HireADroneAPI{
		handlers:        make(map[string]map[string]http.Handler),
		formats:         strfmt.Default,
		defaultConsumes: "application/json",
		defaultProduces: "application/json",
		ServerShutdown:  func() {},
		spec:            spec,
		ServeError:      errors.ServeError,
		JSONConsumer:    runtime.JSONConsumer(),
		JSONProducer:    runtime.JSONProducer(),
		PilotsAddOneHandler: pilots.AddOneHandlerFunc(func(params pilots.AddOneParams) middleware.Responder {
			return middleware.NotImplemented("operation PilotsAddOne has not yet been implemented")
		}),
		AuthInfoFirebaseHandler: AuthInfoFirebaseHandlerFunc(func(params AuthInfoFirebaseParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation AuthInfoFirebase has not yet been implemented")
		}),
		AuthInfoGoogleIDTokenHandler: AuthInfoGoogleIDTokenHandlerFunc(func(params AuthInfoGoogleIDTokenParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation AuthInfoGoogleIDToken has not yet been implemented")
		}),
		AuthInfoAuth0JwkHandler: AuthInfoAuth0JwkHandlerFunc(func(params AuthInfoAuth0JwkParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation AuthInfoAuth0Jwk has not yet been implemented")
		}),
		AuthInfoGoogleJwtHandler: AuthInfoGoogleJwtHandlerFunc(func(params AuthInfoGoogleJwtParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation AuthInfoGoogleJwt has not yet been implemented")
		}),
		PilotsDestroyOneHandler: pilots.DestroyOneHandlerFunc(func(params pilots.DestroyOneParams) middleware.Responder {
			return middleware.NotImplemented("operation PilotsDestroyOne has not yet been implemented")
		}),
		EchoHandler: EchoHandlerFunc(func(params EchoParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation Echo has not yet been implemented")
		}),
		PilotsFindPilotsHandler: pilots.FindPilotsHandlerFunc(func(params pilots.FindPilotsParams) middleware.Responder {
			return middleware.NotImplemented("operation PilotsFindPilots has not yet been implemented")
		}),
		PilotsUpdateOneHandler: pilots.UpdateOneHandlerFunc(func(params pilots.UpdateOneParams) middleware.Responder {
			return middleware.NotImplemented("operation PilotsUpdateOne has not yet been implemented")
		}),

		GoogleJwtAuth: func(token string, scopes []string) (interface{}, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (google_jwt) has not yet been implemented")
		},

		GoogleIDTokenAuth: func(token string, scopes []string) (interface{}, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (google_id_token) has not yet been implemented")
		},

		FirebaseAuth: func(token string, scopes []string) (interface{}, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (firebase) has not yet been implemented")
		},

		Auth0JwkAuth: func(token string, scopes []string) (interface{}, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (auth0_jwk) has not yet been implemented")
		},

		// Applies when the "key" query is set
		APIKeyAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (api_key) key from query param [key] has not yet been implemented")
		},
	}
}

/*HireADroneAPI Submit and view scores */
type HireADroneAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// GoogleJwtAuth registers a function that takes an access token and a collection of required scopes and returns a principal
	// it performs authentication based on an oauth2 bearer token provided in the request
	GoogleJwtAuth func(string, []string) (interface{}, error)

	// GoogleIDTokenAuth registers a function that takes an access token and a collection of required scopes and returns a principal
	// it performs authentication based on an oauth2 bearer token provided in the request
	GoogleIDTokenAuth func(string, []string) (interface{}, error)

	// FirebaseAuth registers a function that takes an access token and a collection of required scopes and returns a principal
	// it performs authentication based on an oauth2 bearer token provided in the request
	FirebaseAuth func(string, []string) (interface{}, error)

	// Auth0JwkAuth registers a function that takes an access token and a collection of required scopes and returns a principal
	// it performs authentication based on an oauth2 bearer token provided in the request
	Auth0JwkAuth func(string, []string) (interface{}, error)

	// APIKeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key key provided in the query
	APIKeyAuth func(string) (interface{}, error)

	// PilotsAddOneHandler sets the operation handler for the add one operation
	PilotsAddOneHandler pilots.AddOneHandler
	// AuthInfoFirebaseHandler sets the operation handler for the auth info firebase operation
	AuthInfoFirebaseHandler AuthInfoFirebaseHandler
	// AuthInfoGoogleIDTokenHandler sets the operation handler for the auth info google Id token operation
	AuthInfoGoogleIDTokenHandler AuthInfoGoogleIDTokenHandler
	// AuthInfoAuth0JwkHandler sets the operation handler for the auth info auth0 jwk operation
	AuthInfoAuth0JwkHandler AuthInfoAuth0JwkHandler
	// AuthInfoGoogleJwtHandler sets the operation handler for the auth info google jwt operation
	AuthInfoGoogleJwtHandler AuthInfoGoogleJwtHandler
	// PilotsDestroyOneHandler sets the operation handler for the destroy one operation
	PilotsDestroyOneHandler pilots.DestroyOneHandler
	// EchoHandler sets the operation handler for the echo operation
	EchoHandler EchoHandler
	// PilotsFindPilotsHandler sets the operation handler for the find pilots operation
	PilotsFindPilotsHandler pilots.FindPilotsHandler
	// PilotsUpdateOneHandler sets the operation handler for the update one operation
	PilotsUpdateOneHandler pilots.UpdateOneHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *HireADroneAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *HireADroneAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *HireADroneAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *HireADroneAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *HireADroneAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *HireADroneAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *HireADroneAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the HireADroneAPI
func (o *HireADroneAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.GoogleJwtAuth == nil {
		unregistered = append(unregistered, "GoogleJwtAuth")
	}

	if o.GoogleIDTokenAuth == nil {
		unregistered = append(unregistered, "GoogleIDTokenAuth")
	}

	if o.FirebaseAuth == nil {
		unregistered = append(unregistered, "FirebaseAuth")
	}

	if o.Auth0JwkAuth == nil {
		unregistered = append(unregistered, "Auth0JwkAuth")
	}

	if o.APIKeyAuth == nil {
		unregistered = append(unregistered, "KeyAuth")
	}

	if o.PilotsAddOneHandler == nil {
		unregistered = append(unregistered, "pilots.AddOneHandler")
	}

	if o.AuthInfoFirebaseHandler == nil {
		unregistered = append(unregistered, "AuthInfoFirebaseHandler")
	}

	if o.AuthInfoGoogleIDTokenHandler == nil {
		unregistered = append(unregistered, "AuthInfoGoogleIDTokenHandler")
	}

	if o.AuthInfoAuth0JwkHandler == nil {
		unregistered = append(unregistered, "AuthInfoAuth0JwkHandler")
	}

	if o.AuthInfoGoogleJwtHandler == nil {
		unregistered = append(unregistered, "AuthInfoGoogleJwtHandler")
	}

	if o.PilotsDestroyOneHandler == nil {
		unregistered = append(unregistered, "pilots.DestroyOneHandler")
	}

	if o.EchoHandler == nil {
		unregistered = append(unregistered, "EchoHandler")
	}

	if o.PilotsFindPilotsHandler == nil {
		unregistered = append(unregistered, "pilots.FindPilotsHandler")
	}

	if o.PilotsUpdateOneHandler == nil {
		unregistered = append(unregistered, "pilots.UpdateOneHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *HireADroneAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *HireADroneAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "google_jwt":

			result[name] = security.BearerAuth(scheme.Name, o.GoogleJwtAuth)

		case "google_id_token":

			result[name] = security.BearerAuth(scheme.Name, o.GoogleIDTokenAuth)

		case "firebase":

			result[name] = security.BearerAuth(scheme.Name, o.FirebaseAuth)

		case "auth0_jwk":

			result[name] = security.BearerAuth(scheme.Name, o.Auth0JwkAuth)

		case "api_key":

			result[name] = security.APIKeyAuth(scheme.Name, scheme.In, o.APIKeyAuth)

		}
	}
	return result

}

// ConsumersFor gets the consumers for the specified media types
func (o *HireADroneAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *HireADroneAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *HireADroneAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the hire a drone API
func (o *HireADroneAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *HireADroneAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"][""] = pilots.NewAddOne(o.context, o.PilotsAddOneHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/auth/info/firebase"] = NewAuthInfoFirebase(o.context, o.AuthInfoFirebaseHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/auth/info/googleidtoken"] = NewAuthInfoGoogleIDToken(o.context, o.AuthInfoGoogleIDTokenHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/auth/info/auth0"] = NewAuthInfoAuth0Jwk(o.context, o.AuthInfoAuth0JwkHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/auth/info/googlejwt"] = NewAuthInfoGoogleJwt(o.context, o.AuthInfoGoogleJwtHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/{id}"] = pilots.NewDestroyOne(o.context, o.PilotsDestroyOneHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/echo"] = NewEcho(o.context, o.EchoHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"][""] = pilots.NewFindPilots(o.context, o.PilotsFindPilotsHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/{id}"] = pilots.NewUpdateOne(o.context, o.PilotsUpdateOneHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *HireADroneAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *HireADroneAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
