package dal

import "github.com/dominictracey/rugby-scores/models"

// IPilotFactory defines a data access layer for pilots
type IPilotFactory interface {
	AddPilot(item *models.Pilot) error

	UpdatePilot(id int64, item *models.Pilot) error
	DeletePilot(id int64) error
	AllPilots(since int64, limit int32) []*models.Pilot
}
