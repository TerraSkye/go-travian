package travian

var bullShitDatabase *BullShitDatabase

// ReadModelFacade is an interface for the readmodel facade
type ReadModelFacade interface {
	GetVillage(uuid string) *Village
	GetVillages() []Village
}

type Tile interface {
}

type Map struct {
	Seed  int
	Tiles [][]*Tile
}

func (m Map) GetVillage(uuid string) *Village {

	return &Village{}
	//return m.Tiles[0][0].(Village);
}

func (m Map) GetVillages() (villages []Village) {

	for i, _ := range bullShitDatabase.Details {
		villages = append(villages, *bullShitDatabase.Details[i])
	}
	return villages

}

//
//// NewBullShitDatabase constructs a new BullShitDatabase
func NewMap(seed int, size int) *Map {
	return &Map{
		Seed:  seed,
		Tiles: [][]*Tile{},
	}
}

// BullShitDatabase is a simple in memory repository
type BullShitDatabase struct {
	Details map[string]*Village
}

// NewBullShitDatabase constructs a new BullShitDatabase
func NewBullShitDatabase() *BullShitDatabase {
	return &BullShitDatabase{

		Details: make(map[string]*Village),
	}
}
