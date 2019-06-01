package travian

import (
	"fmt"
	ycq "github.com/jetbasrawi/go.cqrs"
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
	if bullShitDatabase == nil {
		bullShitDatabase = NewBullShitDatabase()
	}

	return &ResourceProjection{}
}

// Handle handles events and build the projection
func (v *ResourceProjection) Handle(message ycq.EventMessage) {

	switch event := message.Event().(type) {

	case *VillageEstablished:
		fmt.Print("hallo")
		bullShitDatabase.Details[message.AggregateID()] = &Village{ID: event.ID,Name: event.Owner}

	case *FieldUpgraded:
		fmt.Printf("index upgraded %d", event.index)

	}
}

// GetDetailsItem gets an InventoryItemDetailsDto by ID
func (v *ResourceProjection) getWood() (current int, max int) {
	return 0, 800;
}

// GetDetailsItem gets an InventoryItemDetailsDto by ID
func (v *ResourceProjection) getCrop() (current int, max int) {
	return 0, 800;
}

// GetDetailsItem gets an InventoryItemDetailsDto by ID
func (v *ResourceProjection) getIron() (current int, max int) {
	return 0, 800;
}

// GetDetailsItem gets an InventoryItemDetailsDto by ID
func (v *ResourceProjection) getClay() (current int, max int) {
	return 0, 800;
}

func updateResources() {

}
