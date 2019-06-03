package travian

import (
	"fmt"
	ycq "github.com/jetbasrawi/go.cqrs"
	"time"
)

// InventoryItemDetailView handles messages related to inventory and builds an
// in memory read model of inventory item details.
type BuildingQueue struct {
	dispatcher ycq.Dispatcher
	villages   map[string]VillageQueue
}

type VillageQueue struct {
	ID          string
	queue       map[string]QueuedBuilding
	concurrency int8
}
type QueuedBuilding struct {
	ID       string
	Position int8
	Status   string
	TTP      int8
}

// In memory Queue for handling actions
func NewBuildingQueue(dispatcher ycq.Dispatcher) *BuildingQueue {
	queue := BuildingQueue{
		dispatcher,
		map[string]VillageQueue{},
	}
	return &queue
}

// Handle handles events and build the projection
func (v *BuildingQueue) Handle(message ycq.EventMessage) {
	switch event := message.Event().(type) {
	case *VillageEstablished:
		//depending on the village set the concurrency
		v.villages[message.AggregateID()] = VillageQueue{concurrency: 1}
	case *EnqueuedBuilding:
		fmt.Printf("upgraded scheduled %d", event.index)
	case *CompletedBuilding:
		fmt.Printf("upgraded completed %d", event.index)

	case *AbortedBuilding:
		fmt.Printf("upgraded aborted %d", event.index)

	case *DestroyedBuilding:
		fmt.Printf("building Destroyed", event.index)
	}
}

// Returns a string after longer pause
func (b *BuildingQueue) scheduleBuildingUpgrade(upgrade QueuedBuilding) {
	time.Sleep(time.Second * 4)
	fmt.Println("Building upgraded")

}

// Returns a string after longer pause
//func ticker() interface{} {
//	ticker := time.NewTicker(500 * time.Millisecond)
//	go func() {
//		for t := range ticker.C {
//			fmt.Println("Tick at", t)
//		}
//	}()
//
//	// Tickers can be stopped like timers. Once a ticker
//	// is stopped it won't receive any more values on its
//	// channel. We'll stop ours after 1600ms.
//	time.Sleep(1600 * time.Millisecond)
//	ticker.Stop()
//	fmt.Println("Ticker stopped")
//	return "Done"
//}
