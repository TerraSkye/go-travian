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
	Karte(center int, size int) [][]*Tile
	CoordinateForId(id int) Coordinate
	Coordinate(x int, y int) Coordinate
}

type Map struct {
	Seed  int64
	Tiles [][]*Tile
	Size  int
}

func (m Map) CoordinateForId(id int) Coordinate {

	// get the absolute values
	y := id % m.Size
	x := int(math.Floor(float64(id / m.Size)))
	if y < 0 {
		y = y + (m.Size / 2)
	}
	if x < 0 {
		x = x + (m.Size / 2)
	}

	return Coordinate{&m, x, y}

}

func (m Map) Coordinate(x int, y int) Coordinate {
	y = y % m.Size
	x = y % m.Size
	if y < 0 {
		y = y + (m.Size / 2)
	}
	if x < 0 {
		x = x + (m.Size / 2)
	}
	return Coordinate{&m, x, y}
}

func (m Map) Karte(center int, size int) [][]*Tile {

	coordinate := m.CoordinateForId(center)
	x := coordinate.X
	y := coordinate.Y

	fmt.Printf("Center @  (%d | %d) with a size %d", x, y, size)
	tiles := make([][]*Tile, size)
	for i := range tiles {
		tiles[i] = make([]*Tile, size)
		for j := range tiles[i] {
			retrievingY := y + j
			retrievingX := x + i
			if retrievingX > m.Size {
				retrievingX = retrievingX - m.Size
			}
			if retrievingY > m.Size {
				retrievingY = retrievingY - m.Size
			}
			fmt.Printf("[X=%d|Y=%d]\n", retrievingY, retrievingX)
			tiles[i][j] = m.Tiles[x+i][y+i]
		}
	}

	return tiles
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
func NewMap(seed int64, size int) *Map {

	//fmt.Print(tiles)
	world := &Map{
		Seed: seed,
		Size: size,
	}

	rand.Seed(seed)

	tiles := make([][]*Tile, size)
	for i := range tiles {
		tiles[i] = make([]*Tile, size)
	}

	for i := -(size / 2); i < size/2; i++ {
		for j := -(size / 2); j < size/2; j++ {
			var tile Tile
			if i == size || i == 0 || j == size || j == 0 {
				tile = &UnoccupiedVillage{coordinate: Coordinate{world, i, j}, tileType: 3, icon: fmt.Sprintf("t%d", rand.Intn(9))}
			} else {
				switch random := rand.Intn(1000); {

				case random <= 10:
					{
						tile = &UnoccupiedVillage{coordinate: Coordinate{world, i, j}, tileType: 1, icon: fmt.Sprintf("t%d", rand.Intn(9))}

					}
				case random <= 90:
					{
						tile = &UnoccupiedVillage{coordinate: Coordinate{world, i, j}, tileType: 2, icon: fmt.Sprintf("t%d", rand.Intn(9))}
					}
				case random <= 400:
					{
						tile = &UnoccupiedVillage{coordinate: Coordinate{world, i, j}, tileType: 3, icon: fmt.Sprintf("t%d", rand.Intn(9))}

					}
				case random <= 480:
					{
						tile = &UnoccupiedVillage{coordinate: Coordinate{world, i, j}, tileType: 4, icon: fmt.Sprintf("t%d", rand.Intn(9))}

					}
				case random <= 560:
					{
						tile = &UnoccupiedVillage{coordinate: Coordinate{world, i, j}, tileType: 5, icon: fmt.Sprintf("t%d", rand.Intn(9))}
					}
				case random <= 570:
					{
						tile = &UnoccupiedVillage{coordinate: Coordinate{world, i, j}, tileType: 6, icon: fmt.Sprintf("t%d", rand.Intn(9))}
					}
				case random <= 600:
					{
						tile = &UnoccupiedVillage{coordinate: Coordinate{world, i, j}, tileType: 7, icon: fmt.Sprintf("t%d", rand.Intn(9))}
					}
				case random <= 630:
					{
						tile = &UnoccupiedVillage{coordinate: Coordinate{world, i, j}, tileType: 8, icon: fmt.Sprintf("t%d", rand.Intn(9))}
					}
				case random <= 660:
					{
						tile = &UnoccupiedVillage{coordinate: Coordinate{world, i, j}, tileType: 9, icon: fmt.Sprintf("t%d", rand.Intn(9))}
					}
				case random <= 740:
					{
						tile = &UnoccupiedVillage{coordinate: Coordinate{world, i, j}, tileType: 10, icon: fmt.Sprintf("t%d", rand.Intn(9))}
					}
				case random <= 820:
					{
						tile = &UnoccupiedVillage{coordinate: Coordinate{world, i, j}, tileType: 11, icon: fmt.Sprintf("t%d", rand.Intn(9))}
					}
				case random <= 900:
					{
						tile = &UnoccupiedVillage{coordinate: Coordinate{world, i, j}, tileType: 12, icon: fmt.Sprintf("t%d", rand.Intn(9))}
					}
				case random <= 908:
					{
						tile = &Oasis{coordinate: Coordinate{world, i, j}, tileType: 1, icon: fmt.Sprintf("o%d", rand.Intn(9))}
					}
				case random <= 916:
					{
						tile = &Oasis{coordinate: Coordinate{world, i, j}, tileType: 2, icon: fmt.Sprintf("o%d", rand.Intn(9))}
					}
				case random <= 924:
					{
						tile = &Oasis{coordinate: Coordinate{world, i, j}, tileType: 3, icon: fmt.Sprintf("o%d", rand.Intn(9))}
					}
				case random <= 932:
					{
						tile = &Oasis{coordinate: Coordinate{world, i, j}, tileType: 4, icon: fmt.Sprintf("o%d", rand.Intn(9))}
					}
				case random <= 940:
					{
						tile = &Oasis{coordinate: Coordinate{world, i, j}, tileType: 5, icon: fmt.Sprintf("o%d", rand.Intn(9))}
					}
				case random <= 948:
					{
						tile = &Oasis{coordinate: Coordinate{world, i, j}, tileType: 6, icon: fmt.Sprintf("o%d", rand.Intn(9))}
					}
				case random <= 956:
					{
						tile = &Oasis{coordinate: Coordinate{world, i, j}, tileType: 7, icon: fmt.Sprintf("o%d", rand.Intn(9))}
					}
				case random <= 964:
					{
						tile = &Oasis{coordinate: Coordinate{world, i, j}, tileType: 8, icon: fmt.Sprintf("o%d", rand.Intn(9))}
					}
				case random <= 972:
					{
						tile = &Oasis{coordinate: Coordinate{world, i, j}, tileType: 9, icon: fmt.Sprintf("o%d", rand.Intn(9))}
					}
				case random <= 980:
					{
						tile = &Oasis{coordinate: Coordinate{world, i, j}, tileType: 10, icon: fmt.Sprintf("o%d", rand.Intn(9))}
					}
				case random <= 988:
					{
						tile = &Oasis{coordinate: Coordinate{world, i, j}, tileType: 11, icon: fmt.Sprintf("o%d", rand.Intn(9))}
					}
				case random <= 1000:
					{
						tile = &Oasis{coordinate: Coordinate{world, i, j}, tileType: 12, icon: fmt.Sprintf("o%d", rand.Intn(9))}
					}

				}
			}
			//fmt.Printf("(%d - %d)")
			tiles[i+size/2][j+size/2] = &tile
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
