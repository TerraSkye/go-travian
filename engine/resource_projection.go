package travian

import (
	"fmt"
	ycq "github.com/jetbasrawi/go.cqrs"
	"time"
)

type Production struct {
	Lumber           int
	Clay             int
	Iron             int
	Crop             int
	WarehouseStorage int
	GranaryStorage   int
}

// InventoryItemDetailView handles messages related to inventory and builds an
// in memory read model of inventory item details.
type ResourceProjection struct {
}

// NewInventoryItemDetailView constructs a new InventoryItemDetailView
func NewResourceProjection() *ResourceProjection {
	if cache == nil {
		cache = newCache()
	}

	return &ResourceProjection{}
}

// Handle handles events and build the projection
func (v *ResourceProjection) Handle(message ycq.EventMessage) {

	switch event := message.Event().(type) {

	case *VillageEstablished:
		fmt.Print("hallo")
		cache.villages[message.AggregateID()] = &Village{ID: event.ID, Name: event.Owner}
		//delay task for EA upgrades
		go verySlowFunction()
		go ticker()
	case *FieldUpgraded:
		fmt.Printf("index upgraded %d", event.index)

	}
}

// Returns a string after longer pause
func verySlowFunction() interface{} {
	time.Sleep(time.Second * 4)
	fmt.Println("very slow function")
	return "I'm ready"
}

// Returns a string after longer pause
func ticker() interface{} {
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// Tickers can be stopped like timers. Once a ticker
	// is stopped it won't receive any more values on its
	// channel. We'll stop ours after 1600ms.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
	return "Done"
}
