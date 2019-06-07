package travian

import (
	"encoding/json"
	"fmt"
)

type Coordinate struct {
	world *World
	X     int
	Y     int
}

func (c Coordinate) absX() int {
	x := c.X + (c.world.Size / 2)
	if x < 0 {
		x = x + c.world.Size
	}
	return x % c.world.Size
}

func (c Coordinate) absY() int {
	y := c.Y + (c.world.Size / 2)
	if y < 0 {
		y = y + c.world.Size
	}
	return y % c.world.Size
}

func (c Coordinate) Id() int {
	x := c.absX()
	return (x * c.world.Size) + c.absY()
}

func (c Coordinate) North(distance int) int {
	c.X = c.X + distance
	return c.Id()
}

func (c Coordinate) South(distance int) int {
	c.X = c.X - distance
	return c.Id()
}

func (c Coordinate) West(distance int) int {
	c.Y = c.Y + distance
	return c.Id()
}

func (c Coordinate) East(distance int) int {
	c.Y = c.Y - distance
	return c.Id()
}

func NewCoordinate(world World, id int) *Coordinate {

	i := &Coordinate{
		//id >> 8, id % 128,
	}
	return i
}

type Tile interface {
	Image() string
	Id() int
	Coordinate() Coordinate
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

func (t UnoccupiedVillage) Id() int {
	return t.coordinate.Id()
}

func (t UnoccupiedVillage) Coordinate() Coordinate {
	return t.coordinate
}

func (t UnoccupiedVillage) MarshalJSON() ([]byte, error) {

	data := make([]interface{}, 8)

	//'x', 'y', 'nr', 'typ', 'querystring', 'img', 'dname', 'name', 'ew', 'ally', 'vid', 'atyp', 'atime'
	data[0] = t.coordinate.X
	data[1] = t.coordinate.Y
	data[2] = t.coordinate.Id()
	data[3] = t.tileType
	data[4] = fmt.Sprintf("d=%d", t.coordinate.Id())
	data[5] = t.icon
	data[6] = "test"
	data[7] = "test"

	return json.Marshal(data)

	//json.NewEncoder().Encode()
	//[[20,27,3,0,"d=14794&c=03","t7",""]
	//return json.Marshal([]interface{}{t.coordinate.X, t.coordinate.Y, t.tileType, t.tileType, fmt.Sprintf("d=%d", t.coordinate.Id()), "Test"})
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

func (t Oasis) Id() int {
	return t.coordinate.Id()
}

func (t Oasis) Coordinate() Coordinate {
	return t.coordinate
}

func (t Oasis) MarshalJSON() ([]byte, error) {

	data := make([]interface{}, 8)

	//'x', 'y', 'nr', 'typ', 'querystring', 'img', 'dname', 'name', 'ew', 'ally', 'vid', 'atyp', 'atime'
	data[0] = t.coordinate.X
	data[1] = t.coordinate.Y
	data[2] = t.coordinate.Id()
	data[3] = t.tileType
	data[4] = fmt.Sprintf("d=%d", t.coordinate.Id())
	data[5] = t.icon
	data[6] = "test"
	data[7] = "test"

	return json.Marshal(data)
}
