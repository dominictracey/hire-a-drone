package dal

import (
	"log"
	"sync"

	"github.com/dominictracey/rugby-scores/models"
	"github.com/dominictracey/rugby-scores/restapi/operations/pilot"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

// HandlerFactory gives back Handlers which can be inserted into the api
type HandlerFactory struct {
}

var hfInstance *HandlerFactory
var hfOnce sync.Once

// GetHandlerFactoryInstance for singleton
func GetHandlerFactoryInstance() *HandlerFactory {
	hfOnce.Do(func() {
		hfInstance = &HandlerFactory{}
		log.Printf("Created HandlerFactory instance %v", hfInstance)
	})

	return hfInstance
}

// GetAddOnePilotHandler returns an appropriate Handler function, for inmem testing or database
func (hf *HandlerFactory) GetAddOnePilotHandler(test bool) pilot.AddOnePilotHandler {

	return pilot.AddOnePilotHandlerFunc(func(params pilot.AddOnePilotParams, principal *models.Principal) middleware.Responder {
		pf := GetPilotFactory(test)
		log.Printf("Using pilotFactory %v", pf)
		if err := pf.AddPilot(params.Body); err != nil {
			return pilot.NewAddOnePilotDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		//api.Logger("Pilot added: %s %s %v %t", params.Body.FirstName, params.Body.LastName, params.Body.ID, params.Body.Licensed)
		return pilot.NewAddOnePilotCreated().WithPayload(params.Body)
	})
}

//GetDestroyOnePilotNoContentCodeHandler returns handler for api instance
func (hf *HandlerFactory) GetDestroyOnePilotNoContentCodeHandler(test bool) pilot.DestroyOnePilotHandler {
	return pilot.DestroyOnePilotHandlerFunc(func(params pilot.DestroyOnePilotParams, principal *models.Principal) middleware.Responder {
		pf := GetPilotFactory(test)
		if err := pf.DeletePilot(params.ID); err != nil {
			return pilot.NewDestroyOnePilotDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return pilot.NewDestroyOnePilotNoContent()
	})
}

//GetFindPilotsHandler returns handler for api instance
func (hf *HandlerFactory) GetFindPilotsHandler(test bool) pilot.FindPilotsHandler {
	return pilot.FindPilotsHandlerFunc(func(params pilot.FindPilotsParams, principal *models.Principal) middleware.Responder {
		pf := GetPilotFactory(test)
		mergedParams := pilot.NewFindPilotsParams()
		mergedParams.Since = swag.Int64(0)
		if params.Since != nil {
			mergedParams.Since = params.Since
		}
		if params.Limit != nil {
			mergedParams.Limit = params.Limit
		}

		return pilot.NewFindPilotsOK().WithPayload(pf.AllPilots(*mergedParams.Since, *mergedParams.Limit))
	})
}

//GetUpdateOnePilotHandler returns handler for api instance
func (hf *HandlerFactory) GetUpdateOnePilotHandler(test bool) pilot.UpdateOnePilotHandler {
	return pilot.UpdateOnePilotHandlerFunc(func(params pilot.UpdateOnePilotParams, principal *models.Principal) middleware.Responder {
		pf := GetPilotFactory(test)
		if err := pf.UpdatePilot(params.ID, params.Body); err != nil {
			return pilot.NewUpdateOnePilotDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return pilot.NewUpdateOnePilotOK().WithPayload(params.Body)
	})
}
