package pkg

import "fmt"

type Position struct {
	coordinateNorthSouth int
	coordinateEastWest   int
	positions            map[string]bool
}

func NewPosition() Position {
	p := Position{
		coordinateNorthSouth: 0,
		coordinateEastWest:   0,
		positions:            map[string]bool{},
	}
	return p
}

func (p *Position) MoveToNorth() {
	p.coordinateNorthSouth++
}

func (p *Position) MoveToSouth() {
	p.coordinateNorthSouth--
}

func (p *Position) MoveToEast() {
	p.coordinateEastWest--
}

func (p *Position) MoveToWest() {
	p.coordinateEastWest++
}

func (p *Position) RegisterPosition() {
	ok := p.IsAlreadyVisited()
	if !ok {
		key := p.getCoordinatesKey()
		p.positions[key] = true
	}
}

func (p *Position) IsAlreadyVisited() bool {
	key := p.getCoordinatesKey()
	_, ok := p.positions[key]
	return ok
}

func (p *Position) getCoordinatesKey() string {
	return fmt.Sprintf("%d,%d", p.coordinateNorthSouth, p.coordinateEastWest)
}
