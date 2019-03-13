package main

import "fmt"

// Particle is a particle on our grid
type Particle struct {
	dimensions int
	position   [2]int
}

// The directions a particle can move
const (
	DirectionRight int = 0
	DirectionDown  int = 1
	DirectionLeft  int = 2
	DirectionUp    int = 3
)

// NewParticle initialises a new particle
func NewParticle(position [2]int) *Particle {
	p := new(Particle)
	p.dimensions = 2
	p.position = position
	return p
}

// DisplayPosition prints out the current grid position of a particle
func (p *Particle) DisplayPosition() {
	fmt.Println(p.position)
}

// GetAllowedDirections returns the directions a particle can move
func (p *Particle) GetAllowedDirections() []int {
	return []int{DirectionRight, DirectionDown, DirectionLeft, DirectionUp}
}

// DeterminePositionOfNeighbouringCell determines the grid coordinate position
// of a neighbouring cell.
// NOTE: there is no check if the neighbour is inside the grid.
// This check must be performed separately.
func (p *Particle) DeterminePositionOfNeighbouringCell(neighbourDirection int) [2]int {
	var neighbourPosition [2]int

	neighbourPosition[0] = p.position[0]
	neighbourPosition[1] = p.position[1]

	if neighbourDirection == DirectionRight {
		neighbourPosition[0]++
	}

	if neighbourDirection == DirectionLeft {
		neighbourPosition[0]--
	}

	if neighbourDirection == DirectionUp {
		neighbourPosition[1]++
	}

	if neighbourDirection == DirectionDown {
		neighbourPosition[1]--
	}

	return neighbourPosition
}
