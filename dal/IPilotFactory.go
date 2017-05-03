package dal

import "github.com/dominictracey/rugby-scores/models"

// IPilotFactory defines a data access layer for pilots
type IPilotFactory interface {
	AddPilot(item *models.Pilot) error
	UpdatePilot(id int64, item *models.Pilot) error
	DeletePilot(id int64) error
	AllPilots(since int64, limit int32) []*models.Pilot
}

// GetPilotFactory returns an appropriate ORM-ish manager for Pilot types
func GetPilotFactory(test bool) IPilotFactory {
	if test {
		return GetPilotFactoryInMemInstance()
	}
	return GetPilotFactoryDatabaseInstance()

}
