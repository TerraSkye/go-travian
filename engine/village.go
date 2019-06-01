package travian

import (
	ycq "github.com/jetbasrawi/go.cqrs"
)

var (
	production = [...]int{2, 5, 9, 15, 22, 33, 50, 70, 100, 145, 200, 280, 375, 495, 635, 800, 1000, 1300, 1600, 2000, 2450, 3050}
)

type Field struct {
	index int8
	level int8
}

func (f *Field) getProduction() int {
	return production[f.level];
}

func (f *Field) getNextProduction() int {
	return production[f.level+1];
}

type Village struct {
	*ycq.AggregateBase
	ID      string
	Name    string
	loyalty int8
	x       int8
	y       int8
	owner   string
	fields  []Field
}

// NewInventoryItem constructs a new inventory item aggregate.
//
// Importantly it embeds a new AggregateBase.
func NewVillage(id string) *Village {
	i := &Village{
		AggregateBase: ycq.NewAggregateBase(id),
	}
	return i
}

// Apply handles the logic of events on the aggregate.
func (a *Village) Apply(message ycq.EventMessage, isNew bool) {
	if isNew {
		a.TrackChange(message)
	}

	switch ev := message.Event().(type) {

	case *VillageEstablished:
		a.ID = ev.ID
		a.owner = ev.Owner
		a.Name = "village"
		a.fields = make([]Field, 18)
		a.x = ev.X
		a.y = ev.Y

	case *FieldUpgraded:
		a.fields[ev.index].level++
	}
}

type VillageEstablished struct {
	ID    string
	Owner string
	X     int8
	Y     int8
}

// Create raises InventoryItemCreatedEvent
func (a *Village) Establish(x int8, y int8, owner string) error {
	//if name == "" {
	//	return errors.New("the name can not be empty")
	//}

	a.Apply(ycq.NewEventMessage(a.AggregateID(),
		&VillageEstablished{ID: a.AggregateID(), Owner: owner, X: x, Y: y},
		ycq.Int(a.CurrentVersion())), true)

	return nil
}

//Field upgraded
type FieldUpgraded struct {
	ID    string
	index int8
}

func (a *Village) UpgradeField(index int8) error {
	a.Apply(ycq.NewEventMessage(a.AggregateID(),
		&FieldUpgraded{ID: a.AggregateID(), index: index},
		ycq.Int(a.CurrentVersion())), true)
	return nil;
}
