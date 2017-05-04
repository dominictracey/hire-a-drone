package dal

import (
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/dominictracey/rugby-scores/models"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
)

var itemsLock = &sync.Mutex{}

// InMemPilotFactory manages lifecyle and persistence for Pilot model type
type InMemPilotFactory struct {
	items  map[int64]*models.Pilot
	lastID int64
}

var inMemInstance *InMemPilotFactory
var inMemOnce sync.Once

// GetPilotFactoryInMemInstance for singleton
func GetPilotFactoryInMemInstance() *InMemPilotFactory {
	inMemOnce.Do(func() {
		inMemInstance = &InMemPilotFactory{}
		inMemInstance.items = make(map[int64]*models.Pilot)
		inMemInstance.lastID = 0
		createdAt := strfmt.DateTime(time.Now())
		inMemInstance.items[1] = &(models.Pilot{FirstName: "Dominic", LastName: "Tracey", Licensed: true, CreatedAt: createdAt, LastModified: createdAt})

		log.Printf("Created InMemPilotFactory instance %v", inMemInstance)
	})
	return inMemInstance
}

func (pf *InMemPilotFactory) newPilotID() int64 {
	return atomic.AddInt64(&pf.lastID, 1)
}

// AddPilot inserts a new Pilot
func (pf *InMemPilotFactory) AddPilot(item *models.Pilot) error {
	if item == nil {
		return errors.New(500, "pilot must be present")
	}

	//log.Printf("Current lastID: %v for %v", pf.lastID, &pf)
	itemsLock.Lock()
	defer itemsLock.Unlock()

	newID := pf.newPilotID()
	item.ID = newID
	pf.lastID = newID
	item.CreatedAt = strfmt.DateTime(time.Now())
	pf.items[newID] = item

	for id, item := range pf.items {
		log.Printf("Inmem now has: %v %v", id, item)
	}
	//log.Printf("New lastID: %v for %v", pf.lastID, &pf)
	return nil
}

// UpdatePilot updates an existing pilot
func (pf *InMemPilotFactory) UpdatePilot(id int64, item *models.Pilot) error {
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
func (pf *InMemPilotFactory) DeletePilot(id int64) error {
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
func (pf *InMemPilotFactory) AllPilots(since int64, limit int32) (result []*models.Pilot) {
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
