package travian

import (
	"fmt"
	"math"
	"math/rand"
)

var cache *Cache

// ReadModelFacade is an interface for the readmodel facade
type ReadModelFacade interface {
	GetVillage(uuid string) *Village
	GetVillages() []Village
	FetchMapSegment(center int, size int) [][]*Tile
	CoordinateForId(id int) Coordinate
	Coordinate(x int, y int) Coordinate
}

type World struct {
	Seed  int64
	Tiles [][]*Tile
	Size  int
}

func (w World) CoordinateForId(id int) Coordinate {
	quadrant := w.Size / 2
	x := int(math.RoundToEven(float64(id/w.Size))) - quadrant
	y := id%w.Size - quadrant
	return Coordinate{&w, x, y}
}

func (w World) Coordinate(x int, y int) Coordinate {
	return Coordinate{&w, x, y}
}

func (w World) FetchMapSegment(center int, size int) [][]*Tile {
	offset := size / 2
	coordinate := w.CoordinateForId(center)
	tiles := make([][]*Tile, size)
	for i := range tiles {
		tiles[i] = make([]*Tile, size)
		for j := range tiles[i] {
			coordinate := w.Coordinate(coordinate.X-offset+i, coordinate.Y-offset+j)
			tiles[i][j] = w.Tiles[coordinate.absX()][coordinate.absY()]
		}
	}
	return tiles
}

func (w World) GetVillage(uuid string) *Village {

	return cache.villages[uuid]
}

func (w World) GetVillages() (villages []Village) {

	for i, _ := range cache.villages {
		villages = append(villages, *cache.villages[i])
	}
	return villages

}

//
//// newCache constructs a new cache
func NewMap(seed int64, size int) *World {
	rand.Seed(seed)

	world := &World{
		Seed: seed,
		Size: size,
	}
	quandrant := size / 2

	tiles := make([][]*Tile, size)
	for i := range tiles {
		tiles[i] = make([]*Tile, size)
	}
	for x := -quandrant; x <= quandrant; x++ {
		for y := -quandrant; y <= quandrant; y++ {
			var tile Tile
			if x == size || x == 0 || y == size || y == 0 {
				tile = &UnoccupiedVillage{coordinate: Coordinate{world, x, y}, tileType: 3, icon: fmt.Sprintf("t%d", rand.Intn(9))}
			} else {
				switch random := rand.Intn(1000); {

				case random <= 10:
					{
						tile = &UnoccupiedVillage{coordinate: Coordinate{world, x, y}, tileType: 1, icon: fmt.Sprintf("t%d", rand.Intn(9))}

					}
				case random <= 90:
					{
						tile = &UnoccupiedVillage{coordinate: Coordinate{world, x, y}, tileType: 2, icon: fmt.Sprintf("t%d", rand.Intn(9))}
					}
				case random <= 400:
					{
						tile = &UnoccupiedVillage{coordinate: Coordinate{world, x, y}, tileType: 3, icon: fmt.Sprintf("t%d", rand.Intn(9))}

					}
				case random <= 480:
					{
						tile = &UnoccupiedVillage{coordinate: Coordinate{world, x, y}, tileType: 4, icon: fmt.Sprintf("t%d", rand.Intn(9))}

					}
				case random <= 560:
					{
						tile = &UnoccupiedVillage{coordinate: Coordinate{world, x, y}, tileType: 5, icon: fmt.Sprintf("t%d", rand.Intn(9))}
					}
				case random <= 570:
					{
						tile = &UnoccupiedVillage{coordinate: Coordinate{world, x, y}, tileType: 6, icon: fmt.Sprintf("t%d", rand.Intn(9))}
					}
				case random <= 600:
					{
						tile = &UnoccupiedVillage{coordinate: Coordinate{world, x, y}, tileType: 7, icon: fmt.Sprintf("t%d", rand.Intn(9))}
					}
				case random <= 630:
					{
						tile = &UnoccupiedVillage{coordinate: Coordinate{world, x, y}, tileType: 8, icon: fmt.Sprintf("t%d", rand.Intn(9))}
					}
				case random <= 660:
					{
						tile = &UnoccupiedVillage{coordinate: Coordinate{world, x, y}, tileType: 9, icon: fmt.Sprintf("t%d", rand.Intn(9))}
					}
				case random <= 740:
					{
						tile = &UnoccupiedVillage{coordinate: Coordinate{world, x, y}, tileType: 10, icon: fmt.Sprintf("t%d", rand.Intn(9))}
					}
				case random <= 820:
					{
						tile = &UnoccupiedVillage{coordinate: Coordinate{world, x, y}, tileType: 11, icon: fmt.Sprintf("t%d", rand.Intn(9))}
					}
				case random <= 900:
					{
						tile = &UnoccupiedVillage{coordinate: Coordinate{world, x, y}, tileType: 12, icon: fmt.Sprintf("t%d", rand.Intn(9))}
					}
				case random <= 908:
					{
						tile = &Oasis{coordinate: Coordinate{world, x, y}, tileType: 1, icon: fmt.Sprintf("o%d", rand.Intn(9))}
					}
				case random <= 916:
					{
						tile = &Oasis{coordinate: Coordinate{world, x, y}, tileType: 2, icon: fmt.Sprintf("o%d", rand.Intn(9))}
					}
				case random <= 924:
					{
						tile = &Oasis{coordinate: Coordinate{world, x, y}, tileType: 3, icon: fmt.Sprintf("o%d", rand.Intn(9))}
					}
				case random <= 932:
					{
						tile = &Oasis{coordinate: Coordinate{world, x, y}, tileType: 4, icon: fmt.Sprintf("o%d", rand.Intn(9))}
					}
				case random <= 940:
					{
						tile = &Oasis{coordinate: Coordinate{world, x, y}, tileType: 5, icon: fmt.Sprintf("o%d", rand.Intn(9))}
					}
				case random <= 948:
					{
						tile = &Oasis{coordinate: Coordinate{world, x, y}, tileType: 6, icon: fmt.Sprintf("o%d", rand.Intn(9))}
					}
				case random <= 956:
					{
						tile = &Oasis{coordinate: Coordinate{world, x, y}, tileType: 7, icon: fmt.Sprintf("o%d", rand.Intn(9))}
					}
				case random <= 964:
					{
						tile = &Oasis{coordinate: Coordinate{world, x, y}, tileType: 8, icon: fmt.Sprintf("o%d", rand.Intn(9))}
					}
				case random <= 972:
					{
						tile = &Oasis{coordinate: Coordinate{world, x, y}, tileType: 9, icon: fmt.Sprintf("o%d", rand.Intn(9))}
					}
				case random <= 980:
					{
						tile = &Oasis{coordinate: Coordinate{world, x, y}, tileType: 10, icon: fmt.Sprintf("o%d", rand.Intn(9))}
					}
				case random <= 988:
					{
						tile = &Oasis{coordinate: Coordinate{world, x, y}, tileType: 11, icon: fmt.Sprintf("o%d", rand.Intn(9))}
					}
				case random <= 1000:
					{
						tile = &Oasis{coordinate: Coordinate{world, x, y}, tileType: 12, icon: fmt.Sprintf("o%d", rand.Intn(9))}
					}

				}
			}
			//fmt.Printf("(%d - %d)")
			tiles[x+quandrant][y+quandrant] = &tile
		}
	}

	world.Tiles = tiles
	return world
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
