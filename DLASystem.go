package main

import "fmt"
import "math"
import "os"
import "os/exec"

// DLASystem is a contained system for modelling Diffusion Limited Aggregation
type DLASystem struct {
	particleList []*Particle

	numberOfParticles int
	clusterRadius     float64

	addCircle  float64
	killCircle float64

	gridSize int
	grid     [][]bool

	endNumberOfParticles int

	isRunning   bool
	slowNotFast bool

	lastParticleIsActive bool

	addRatio  float64
	killRatio float64

	randomGenerator *RandomGenerator

	verbose bool
}

// NewDLASystem initialises a new DLA system
func NewDLASystem(gridSize int, addRatio float64, killRatio float64, endNumberOfParticles int, seed int64) *DLASystem {
	fmt.Printf("Creating system, grid size %d.\n", gridSize)

	dla := new(DLASystem)
	dla.verbose = false
	dla.gridSize = gridSize
	dla.addRatio = addRatio
	dla.killRatio = killRatio
	dla.endNumberOfParticles = endNumberOfParticles

	dla.particleList = make([]*Particle, endNumberOfParticles)
	dla.numberOfParticles = 0

	dla.grid = make([][]bool, gridSize)
	for i := 0; i < gridSize; i++ {
		dla.grid[i] = make([]bool, gridSize)
	}

	dla.randomGenerator = NewRandomGenerator(seed)
	dla.Reset()

	return dla
}

// DisplayGrid prints grid to stdout
func (dla *DLASystem) DisplayGrid() {
	// fmt.Printf("\n")
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
	for _, column := range dla.grid {
		for _, row := range column {
			if row {
				fmt.Printf("● ")
			} else {
				fmt.Printf("◌ ")
			}
		}
		fmt.Printf("\n")
	}
}

// Reset removes any existing particles and sets back to initial conditions
func (dla *DLASystem) Reset() {
	dla.isRunning = false
	dla.ClearParticles()
	dla.lastParticleIsActive = false

	// Reset grid
	for i, column := range dla.grid {
		for j := range column {
			dla.grid[i][j] = false
		}
	}

	dla.addCircle = 5.0
	dla.killCircle = dla.killRatio * dla.addCircle
	dla.clusterRadius = 0

	// Add a single particle at the origin
	dla.AddParticle([2]int{0, 0})
}

// ClearParticles deletes particle and the particle list
func (dla *DLASystem) ClearParticles() {
	dla.particleList = make([]*Particle, dla.endNumberOfParticles)
	dla.numberOfParticles = 0
}

// Update updates the system - if there is an active particle
// then move it, create a new particle on the adding circle otherwise
func (dla *DLASystem) Update() {
	if dla.lastParticleIsActive {
		dla.MoveLastParticle()
	} else if dla.numberOfParticles < dla.endNumberOfParticles {
		dla.AddParticleOnAddCircle()
		dla.SetParticleActive()
	}
}

// SetParticleActive marks the latest added particle as not yet stuck
func (dla *DLASystem) SetParticleActive() {
	dla.lastParticleIsActive = true
}

// SetParticleInactive marks latest added particle as now stuck
func (dla *DLASystem) SetParticleInactive() {
	dla.lastParticleIsActive = false
}

// DeterminePositionDistanceFromOrigin determines the radius of a coordinate position (x,y)
func (dla *DLASystem) DeterminePositionDistanceFromOrigin(position [2]int) float64 {

	radiusSquared := float64(position[0]*position[0] + position[1]*position[1])
	return math.Sqrt(radiusSquared)
}

// MoveLastParticle moves the last added particle
func (dla *DLASystem) MoveLastParticle() {
	if dla.verbose {
		fmt.Println("Moving particle...")
	}

	var lastParticle *Particle
	var neighbourDirection int
	var distanceOfNewPositionFromOrigin float64
	var positionOfNewCell [2]int

	lastParticle = dla.particleList[dla.numberOfParticles-1]
	neighbourDirection = dla.randomGenerator.RandomInt(4)
	positionOfNewCell = lastParticle.DeterminePositionOfNeighbouringCell(neighbourDirection)

	distanceOfNewPositionFromOrigin = dla.DeterminePositionDistanceFromOrigin(positionOfNewCell)
	// fmt.Printf("Distance is %f\nKill circle is %f", distanceOfNewPositionFromOrigin, dla.killCircle)
	if distanceOfNewPositionFromOrigin > dla.killCircle {
		if dla.verbose {
			fmt.Println("KILLING PARTICLE")
		}
		dla.SetGrid(lastParticle.position, false)
		dla.particleList[dla.numberOfParticles-1] = nil
		dla.numberOfParticles--
		dla.SetParticleInactive()
	} else if dla.ReadGrid(positionOfNewCell) == false {
		dla.SetGrid(lastParticle.position, false)
		lastParticle.position = positionOfNewCell
		dla.SetGrid(positionOfNewCell, true)

		if dla.ShouldParticleStick(lastParticle) {
			dla.SetParticleInactive()
			dla.UpdateClusterRadius(positionOfNewCell)
		}
	} else {
		// If we get here then we are trying to move to an occupied site
		// This should never happen as long as sticking probability = 1
		fmt.Printf("Move to (%d, %d) REJECTED", positionOfNewCell[0], positionOfNewCell[1])
	}
}

// UpdateClusterRadius sets the radius of the system
func (dla *DLASystem) UpdateClusterRadius(lastParticlePosition [2]int) {
	distanceOfNewParticleFromOrigin := dla.DeterminePositionDistanceFromOrigin(lastParticlePosition)
	if distanceOfNewParticleFromOrigin > dla.clusterRadius {
		dla.clusterRadius = distanceOfNewParticleFromOrigin
		// Add circle is supposed to be:
		// Either 20% bigger than cluster radius or at least 2 bigger
		check := dla.clusterRadius * dla.addRatio
		if check < (dla.clusterRadius + 2) {
			check = dla.clusterRadius + 2
		}
		if dla.addCircle < check {
			dla.addCircle = check
			dla.killCircle = dla.killRatio * dla.addCircle
		}
		dla.CheckIfShouldStop()
	}
}

// CheckIfShouldStop stops the simulation if the cluster is big enough
// To be safe, we need the kill circle to be at leat 2 less than the
// edge of the grid
func (dla *DLASystem) CheckIfShouldStop() bool {
	if dla.killCircle+2 >= float64(dla.gridSize)/2 {
		dla.PauseSimulation()
		if dla.verbose {
			fmt.Println("Simulation finishing...")
		}
		return true
	}
	return false
}

// PauseSimulation stops the simulation
func (dla *DLASystem) PauseSimulation() {
	dla.isRunning = false
}

// ShouldParticleStick determines if a particle should stick to a neighbour
// when it moves to position
func (dla *DLASystem) ShouldParticleStick(particle *Particle) bool {
	var shouldStick = false
	var positionOfNeighbouringCell [2]int
	for _, direction := range particle.GetAllowedDirections() {
		positionOfNeighbouringCell = particle.DeterminePositionOfNeighbouringCell(direction)
		if dla.ReadGrid(positionOfNeighbouringCell) {
			shouldStick = true
		}
	}
	return shouldStick
}

// AddParticleOnAddCircle adds a new particle to the system!
func (dla *DLASystem) AddParticleOnAddCircle() {
	if dla.verbose {
		fmt.Println("\nAdding new particle on add circle")
	}

	var position [2]int
	var theta, sinTheta, cosTheta float64

	theta = dla.randomGenerator.rand.Float64() * 2 * math.Pi
	sinTheta, cosTheta = math.Sincos(theta)

	position[0] = int(math.Floor(dla.addCircle * cosTheta))
	position[1] = int(math.Floor(dla.addCircle * sinTheta))

	if dla.ReadGrid(position) == false {
		dla.AddParticle(position)
	} else {
		fmt.Printf("Failed adding particle at (%d, %d)", position[0], position[1])
	}
}

// AddParticle adds a particle to a grid
func (dla *DLASystem) AddParticle(position [2]int) {
	var particle *Particle
	particle = NewParticle(position)
	dla.particleList[dla.numberOfParticles] = particle
	dla.numberOfParticles++
	dla.SetGrid(position, true)
}

// SetGrid sets the value of a grid cell for a particular position
// Note position of initial particle is (0,0)
// This correponds to the middle of the grid array: dla.grid[halfGrid][halfGrid]
func (dla *DLASystem) SetGrid(position [2]int, isOccupied bool) {
	halfGrid := dla.gridSize / 2
	dla.grid[position[0]+halfGrid][position[1]+halfGrid] = isOccupied
}

// ReadGrid determines if a particle exists at a given (x, y) position
func (dla *DLASystem) ReadGrid(position [2]int) bool {
	halfGrid := dla.gridSize / 2
	return dla.grid[position[0]+halfGrid][position[1]+halfGrid]
}
