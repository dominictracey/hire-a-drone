package dal

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/datastore"

	"github.com/dominictracey/rugby-scores/models"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
)

// PilotDBFactory manages lifecyle and persistence for Pilot model type
type PilotDBFactory struct {
	//lastID int64
}

var dbInstance *PilotDBFactory
var dbOnce sync.Once

// GetPilotFactoryDatabaseInstance for singleton
func GetPilotFactoryDatabaseInstance() *PilotDBFactory {
	dbOnce.Do(func() {
		dbInstance = &PilotDBFactory{}
		//dbInstance.lastID = 7
		log.Printf("Created PilotDBFactory instance %v", dbInstance)
	})

	return dbInstance
}

// AddPilot inserts a new Pilot
func (pf *PilotDBFactory) AddPilot(pilot *models.Pilot) error {
	if pilot == nil {
		return errors.New(500, "pilot must be present")
	}

	ctx := context.Background()

	// Setx your Google Cloud Platform project ID.
	projectID := "rugby-scores-7"

	// Creates a client.
	client, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return err
	}

	// timestamp for creation
	pilot.CreatedAt = strfmt.DateTime(time.Now())

	// Sets the kind for the new entity.
	kind := "Pilot"

	// Creates a Key instance.
	pilotKey := datastore.IncompleteKey(kind, nil)

	// Saves the new entity.
	if pilotKey, err = client.Put(ctx, pilotKey, pilot); err != nil {
		log.Fatalf("Failed to save pilot: %v", err)
		return err
	}

	// save the ID
	// pilot.ID = pilotKey.ID
	// if _, err := client.Put(ctx, pilotKey, pilot); err != nil {
	// 	log.Fatalf("Failed to save pilot's new ID: %v", err)
	// 	return err
	// }

	fmt.Printf("Saved %v at %v\n", pilot, pilotKey)

	//try to get it back
	var pilout models.Pilot
	if err := client.Get(ctx, pilotKey, &pilout); err != nil {
		log.Fatalf("Failed to get new pilot: %v", err)
		return err
	}
	log.Printf("Retrieved new pilot: %v", pilout)

	return nil
}

// UpdatePilot updates an existing pilot
func (pf *PilotDBFactory) UpdatePilot(id int64, pilot *models.Pilot) error {
	if pilot == nil {
		return errors.New(500, "item must be present")
	}

	ctx := context.Background()
	client, _ := datastore.NewClient(ctx, "rugby-scores-7")
	//pilot := &Pilot{} // Populated with appropriate data.
	key := datastore.IncompleteKey("Pilot", nil)
	key.ID = id
	pilot.CreatedAt = strfmt.DateTime(time.Now())
	// [START upsert]
	key, err := client.Put(ctx, key, pilot)
	if err != nil {
		log.Fatalf("Failed to update pilot: %v", err)
	}
	// [END upsert]

	return nil
}

// DeletePilot deletes a pilot
func (pf *PilotDBFactory) DeletePilot(id int64) error {
	ctx := context.Background()
	client, _ := datastore.NewClient(ctx, "rugby-scores-7")
	key := datastore.IncompleteKey("Pilot", nil)
	key.ID = id

	err := client.Delete(ctx, key)
	if err != nil {
		return err
	}

	log.Printf("Successfully deleted pilot with id %v", id)

	return nil
}

// AllPilots allows paging through all pilots
func (pf *PilotDBFactory) AllPilots(since int64, limit int32) (result []*models.Pilot) {
	// result = make([]*models.Pilot, 0)
	// for id, item := range pf.items {
	// 	if len(result) >= int(limit) {
	// 		return
	// 	}
	// 	if since == 0 || id > since {
	// 		result = append(result, item)
	// 	}
	// }
	ctx := context.Background()
	client, _ := datastore.NewClient(ctx, "rugby-scores-7")
	// [START basic_query]
	query := datastore.NewQuery("Pilot")
	// [END basic_query]
	// [START run_query]
	result = make([]*models.Pilot, 0)
	it := client.Run(ctx, query)
	for {
		var pilot models.Pilot
		_, err := it.Next(&pilot)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error fetching next pilot: %v", err)
		}
		//fmt.Printf("Pilot %q %q\n", pilot.FirstName, pilot.LastName)

		result = append(result, &pilot)
	}

	return
}
