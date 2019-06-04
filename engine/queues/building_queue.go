package queues

import (
	ycq "github.com/jetbasrawi/go.cqrs"
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
	status      chan bool
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

//// Handle handles events and build the projection
//func (v *BuildingQueue) Handle(message ycq.EventMessage) {
//	switch event := message.Event().(type) {
//	case travian.VillageEstablished:
//		//depending on the village set the concurrency
//		v.villages[message.AggregateID()] = VillageQueue{concurrency: 1, status: make(chan bool)}
//	case travian.EnqueuedBuilding:
//		v.villages[message.AggregateID()].queue[event.ID] = QueuedBuilding{}
//		v.scheduleBuildingUpgrade(QueuedBuilding{TTP: 500000}, v.villages[message.AggregateID()].status)
//		fmt.Printf("upgraded scheduled %d", event.index)
//	case travian.CompletedBuilding:
//		fmt.Printf("upgraded completed %d", event.index)
//	case travian.AbortedBuilding:
//		fmt.Printf("upgraded aborted %d", event.index)
//	case travian.DestroyedBuilding:
//		fmt.Printf("building Destroyed", event.index)
//	}
//}
//
//// Returns a string after longer pause
//func (b *BuildingQueue) scheduleBuildingUpgrade(upgrade QueuedBuilding, status chan bool) {
//	select {
//	case <-status:
//		//pending task is canceld
//		break
//	case <-time.After(upgrade.TTP * time.Millisecond):
//		//task completed
//		em := ycq.NewCommandMessage(id, &UpgradeField{
//			X:     0,
//			Y:     0,
//			Owner: r.Form.Get("name"),
//		})
//		err = b.dispatcher.Dispatch(em)
//		break;
//
//	}
//}
//
