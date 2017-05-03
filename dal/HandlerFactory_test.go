package dal

import (
	"log"
	"testing"

	"github.com/dominictracey/rugby-scores/models"
	"github.com/dominictracey/rugby-scores/restapi/operations/pilot"
	"github.com/go-openapi/runtime/middleware"
)

func TestGetAddOnePilotHandlerFunc(t *testing.T) {
	handler := GetHandlerFactoryInstance().GetAddOnePilotHandler(true)
	log.Printf("Using Handler %v", handler)
	var params pilot.AddOnePilotParams = (pilot.AddOnePilotParams{Body: &(models.Pilot{FirstName: "Valerie", LastName: "Green", Licensed: false})})
	var principal models.Principal = "boss"

	var res middleware.Responder = handler.Handle(params, &principal)
	log.Printf("Called Handler %v", handler)
	//pilot := &(models.Pilot{FirstName: "Dominic", LastName: "Tracey"})
	if res == nil {
		t.Errorf("Error invoking handler")
	}

}
