package travian

import (
	ycq "github.com/jetbasrawi/go.cqrs"
	"log"
	"reflect"
)

//command
type EstablishVillage struct {
	X     int8
	Y     int8
	Owner string
}

type UpgradeField struct {
	ID    string
	Index int8
}

// repos
type VillageRepository interface {
	Load(string, string) (*Village, error)
	Save(ycq.AggregateRoot, *int) error
}

// InventoryCommandHandlers provides methods for processing commands related
// to inventory items.
type VillageCommandHandlers struct {
	repo VillageRepository
}

// NewInventoryCommandHandlers contructs a new InventoryCommandHandlers
func NewVillageCommandHandlers(repo VillageRepository) *VillageCommandHandlers {
	return &VillageCommandHandlers{
		repo: repo,
	}
}

// Handle processes inventory item commands.
func (h *VillageCommandHandlers) Handle(message ycq.CommandMessage) error {
	var village *Village

	switch cmd := message.Command().(type) {

	case *EstablishVillage:
		village = NewVillage(message.AggregateID())
		if err := village.Establish(cmd.X, cmd.Y, cmd.Owner); err != nil {
			return &ycq.ErrCommandExecution{Command: message, Reason: err.Error()}
		}
		return h.repo.Save(village, ycq.Int(village.OriginalVersion()))

	case *UpgradeField:

		village, _ = h.repo.Load(reflect.TypeOf(&Village{}).Elem().Name(), message.AggregateID())
		if err := village.UpgradeField(cmd.Index); err != nil {
			return &ycq.ErrCommandExecution{Command: message, Reason: err.Error()}
		}
		return h.repo.Save(village, ycq.Int(village.OriginalVersion()))
	default:
		log.Fatalf("InventoryCommandHandlers has received a command that it is does not know how to handle, %#v", cmd)
	}

	return nil
}
