package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

func runSystem(seed int64, wg *sync.WaitGroup, systemID int, stickingProbability float64) {
	defer wg.Done()

	dla := NewDLASystem(500, 1.2, 1.7, 2000, seed, stickingProbability, false)
	dla.isRunning = true

	var filenameComponents []string
	var csvFilename string
	// var gridFilename string

	filenameComponents = []string{"results/stick4/ensemble-p", strconv.Itoa(int(stickingProbability * 10)), "-#", strconv.Itoa(systemID), ".csv"}
	csvFilename = strings.Join(filenameComponents, "")

	// filenameComponents = []string{"results/stick/ensemble-p", strconv.Itoa(int(stickingProbability * 10)), "-#", strconv.Itoa(systemID), ".dat"}
	// gridFilename = strings.Join(filenameComponents, "")

	f, err := os.OpenFile(csvFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	var lastNumberOfParticlesWrittenAt int
	i := 1
	for {
		// time.Sleep(40 * time.Millisecond)

		dla.Update()

		if dla.numberOfParticles%100 == 0 && dla.lastParticleIsActive == false &&
			lastNumberOfParticlesWrittenAt != dla.numberOfParticles {
			if _, err := f.Write([]byte(fmt.Sprintf("%d,%f\n", dla.numberOfParticles, dla.clusterRadius))); err != nil {
				log.Fatal(err)
			}
			if err == nil {
				lastNumberOfParticlesWrittenAt = dla.numberOfParticles
			}
		}

		// if i%10 == 0 {
		// dla.DisplayGrid()
		// }
		i++

		if dla.isRunning == false {
			break
		}
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Writing grid " + strconv.Itoa(systemID) + " to disc...")
	// dla.PersistGridToFile(gridFilename)

	fmt.Println("System " + strconv.Itoa(systemID) + " finished!")
}

// func main() {
// 	numberOfCores := runtime.NumCPU()
// 	fmt.Printf("Using " + strconv.Itoa(numberOfCores) + " cores...\n")
// 	runtime.GOMAXPROCS(numberOfCores)

// 	var wg sync.WaitGroup

// 	i := 0
// 	for i < 1 {
// 		wg.Add(1)
// 		fmt.Printf("Starting system %d\n", i)
// 		go runSystem(int64(i), &wg, i, 1.0)
// 		i++
// 	}
// 	fmt.Println("")

// 	wg.Wait()
// 	fmt.Println("\nAll systems complete!!!")
// }

func main() {
	numberOfCores := runtime.NumCPU()
	fmt.Printf("Using " + strconv.Itoa(numberOfCores) + " cores...\n")
	runtime.GOMAXPROCS(numberOfCores)

	var wg sync.WaitGroup

	i := 0
	var stickingProbability float64
	for i < 1000 {
		j := 10
		// for j <= 10 {
		wg.Add(1)
		stickingProbability = float64(j) / 10.0
		fmt.Printf("Starting system #%d, pstick = %f\n", i, stickingProbability)
		go runSystem(int64(i), &wg, i, stickingProbability)
		// j = j + 1
		// }
		i = i + 1
	}
	fmt.Println("")

	wg.Wait()
	fmt.Println("\nAll systems complete!!!")
}
