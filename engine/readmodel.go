package travian

var cache *Cache

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

	return cache.villages[uuid]
}

func (m Map) GetVillages() (villages []Village) {

	for i, _ := range cache.villages {
		villages = append(villages, *cache.villages[i])
	}
	return villages

}

//
//// newCache constructs a new cache
func NewMap(seed int, size int) *Map {
	return &Map{
		Seed:  seed,
		Tiles: [][]*Tile{},
	}
}

// cache is a simple in memory repository
type Cache struct {
	villages map[string]*Village
}

// newCache constructs a new cache
func newCache() *Cache {
	return &Cache{
		villages: make(map[string]*Village),
	}
}
