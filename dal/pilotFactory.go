package dal

import (
	"log"
	"sync"
	"sync/atomic"

	"github.com/dominictracey/rugby-scores/models"
	"github.com/go-openapi/errors"
)

var itemsLock = &sync.Mutex{}

// PilotFactory manages lifecyle and persistence for Pilot model type
type PilotFactory struct {
	items  map[int64]*models.Pilot
	lastID int64
}

var instance *PilotFactory
var once sync.Once

// GetPilotFactoryInstance for singleton
func GetPilotFactoryInstance() *PilotFactory {
	once.Do(func() {
		instance = &PilotFactory{}
		instance.items = make(map[int64]*models.Pilot)
		instance.lastID = 0
		log.Printf("Created PilotFactory instance %v", instance)
	})

	return instance
}

func (pf *PilotFactory) newPilotID() int64 {
	return atomic.AddInt64(&pf.lastID, 1)
}

// AddPilot inserts a new Pilot
func (pf *PilotFactory) AddPilot(item *models.Pilot) error {
	if item == nil {
		return errors.New(500, "item must be present")
	}

	//log.Printf("Current lastID: %v for %v", pf.lastID, &pf)
	itemsLock.Lock()
	defer itemsLock.Unlock()

	newID := pf.newPilotID()
	item.ID = newID
	pf.lastID = newID
	pf.items[newID] = item
	//log.Printf("New lastID: %v for %v", pf.lastID, &pf)
	return nil
}

// UpdatePilot updates an existing pilot
func (pf *PilotFactory) UpdatePilot(id int64, item *models.Pilot) error {
	if item == nil {
		return errors.New(500, "item must be present")
	}

	itemsLock.Lock()
	defer itemsLock.Unlock()

	_, exists := pf.items[id]
	if !exists {
		return errors.NotFound("not found: item %d", id)
	}

	item.ID = id
	pf.items[id] = item
	return nil
}

// DeletePilot deletes a pilot
func (pf *PilotFactory) DeletePilot(id int64) error {
	itemsLock.Lock()
	defer itemsLock.Unlock()

	_, exists := pf.items[id]
	if !exists {
		return errors.NotFound("not found: item %d", id)
	}

	delete(pf.items, id)
	return nil
}

// AllPilots allows paging through all pilots
func (pf *PilotFactory) AllPilots(since int64, limit int32) (result []*models.Pilot) {
	result = make([]*models.Pilot, 0)
	for id, item := range pf.items {
		if len(result) >= int(limit) {
			return
		}
		if since == 0 || id > since {
			result = append(result, item)
		}
	}
	return
}
