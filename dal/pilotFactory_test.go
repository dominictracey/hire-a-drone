package dal

import (
	"testing"

	"github.com/dominictracey/rugby-scores/models"
	"github.com/go-openapi/strfmt"
)

func TestGetPilotFactory(t *testing.T) {
	pf := GetPilotFactory(true)
	if pf == nil {
		t.Error("No factory returned")
	}
}

func TestAddNilPilot(t *testing.T) {
	pf := GetPilotFactory(true)

	var pilot *models.Pilot
	err := pf.AddPilot(pilot)

	if err == nil {
		t.Error("Didn't get error adding nil pilot")
	}

	if err.Error() != "pilot must be present" {
		t.Error("Should have error message for adding nil pilot")
	}
}

func TestAddPilot(t *testing.T) {
	pf := GetPilotFactory(true)

	pilot := &(models.Pilot{FirstName: "Dominic", LastName: "Tracey"})
	if err := pf.AddPilot(pilot); err != nil {
		t.Errorf("Error adding pilot %v", err)
	}
}

func TestDeletePilot(t *testing.T) {
	pf := GetPilotFactory(true)

	//pilot := &(models.Pilot{FirstName: "Dominic", LastName: "Tracey"})
	if err := pf.DeletePilot(1); err != nil {
		t.Errorf("Error deleting pilot %v", err)
	}
}

func TestGetPilots(t *testing.T) {
	pf := GetPilotFactory(true)

	result := pf.AllPilots(0, 100)

	if result == nil {
		t.Errorf("Got empty response from AllPilots")
	}

	if result[0].CreatedAt == strfmt.NewDateTime() {
		t.Errorf("Got pilot with bad CreatedAt timestamp: %v", result[0].CreatedAt)
	}

	// if len(result) != 2 {
	// 	t.Errorf("Got AllPilots length of %v", len(result))
	// }

}
