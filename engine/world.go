package travian

import (
	"encoding/json"
	"fmt"
)

type Coordinate struct {
	world *Map
	X     int
	Y     int
}

func (c Coordinate) Id() int {

	return (c.X * c.world.Size) + c.Y
}

func (c Coordinate) North(distance int) int {

	return ((c.X + distance) * c.world.Size) + c.Y
}

func (c Coordinate) South(distance int) int {

	return ((c.X - distance) * c.world.Size) + c.Y
}

func (c Coordinate) West(distance int) int {

	return (c.X * c.world.Size) + c.Y + distance
}

func (c Coordinate) East(distance int) int {

	return (c.X * c.world.Size) + c.Y - distance
}

func NewCoordinate(world Map, id int) *Coordinate {

	i := &Coordinate{
		//id >> 8, id % 128,
	}
	return i
}

type Tile interface {
	Image() string
	Id() int
}

//Villages

type UnoccupiedVillage struct {
	Tile
	coordinate Coordinate
	tileType   int
	icon       string
}

func (t UnoccupiedVillage) Image() string {
	return t.icon
}

func (t UnoccupiedVillage) ID() int {
	return t.coordinate.Id()
}

func (t UnoccupiedVillage) MarshalJSON() ([]byte, error) {

	//[[20,27,3,0,"d=14794&c=03","t7",""]
	return json.Marshal([]interface{}{t.coordinate.X, t.coordinate.Y, t.tileType, t.tileType, fmt.Sprintf("d=%d", t.coordinate.Id()), "Test"})
}

//OASIS

type Oasis struct {
	Tile
	coordinate Coordinate
	tileType   int
	icon       string
}

func (t Oasis) Image() string {
	return t.icon
}

func (t Oasis) ID() int {
	return t.coordinate.Id()
}
