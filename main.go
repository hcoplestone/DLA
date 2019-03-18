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

func runSystem(seed int64, wg *sync.WaitGroup, systemID int) {
	defer wg.Done()

	dla := NewDLASystem(450, 1.2, 1.7, 3000, seed, false)
	dla.isRunning = true

	var filenameComponents []string
	var csvFilename, gridFilename string

	filenameComponents = []string{"results/second/ensemble", strconv.Itoa(systemID), ".csv"}
	csvFilename = strings.Join(filenameComponents, "")

	filenameComponents = []string{"results/second/ensemble", strconv.Itoa(systemID), ".dat"}
	gridFilename = strings.Join(filenameComponents, "")

	f, err := os.OpenFile(csvFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	var lastNumberOfParticlesWrittenAt int
	i := 1
	for {
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

	fmt.Println("Writing grid " + strconv.Itoa(systemID) + " to disc...")
	dla.PersistGridToFile(gridFilename)

	fmt.Println("System " + strconv.Itoa(systemID) + " finished!")
}

func main() {
	numberOfCores := runtime.NumCPU()
	fmt.Printf("Using " + strconv.Itoa(numberOfCores) + " cores...\n")
	runtime.GOMAXPROCS(numberOfCores)

	var wg sync.WaitGroup

	i := 0
	for i < 4 {
		wg.Add(1)
		fmt.Printf("Starting system %d\n", i)
		go runSystem(int64(i), &wg, i)
		i++
	}
	fmt.Println("")

	wg.Wait()
	fmt.Println("\nAll systems complete!!!")
}
