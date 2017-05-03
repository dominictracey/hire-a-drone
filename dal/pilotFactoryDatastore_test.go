package dal

import (
	"testing"

	"github.com/dominictracey/rugby-scores/models"
)

func TestGetPilotFactory(t *testing.T) {
	pf := GetPilotDBFactoryInstance()
	if pf == nil {
		t.Error("No factory returned")
	}
}

func TestAddNilPilot(t *testing.T) {
	pf := GetPilotDBFactoryInstance()

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
	pf := GetPilotDBFactoryInstance()

	pilot := &(models.Pilot{FirstName: "Dominic", LastName: "Tracey"})
	if err := pf.AddPilot(pilot); err != nil {
		t.Errorf("Error adding pilot %v", err)
	}
}

func TestGetPilots(t *testing.T) {
	pf := GetPilotDBFactoryInstance()

	result := pf.AllPilots(0, 100)

	if result == nil {
		t.Errorf("Go empty response from AllPilots")
	}

	if len(result) != 7 {
		t.Errorf("Got AllPilots length of %v", len(result))
	}

}
