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

	var filenameComponents []string
	var csvFilename, gridFilename string

	filenameComponents = []string{"results/ensemble", strconv.Itoa(systemID), ".csv"}
	csvFilename = strings.Join(filenameComponents, "")

	filenameComponents = []string{"results/ensemble", strconv.Itoa(systemID), ".dat"}
	gridFilename = strings.Join(filenameComponents, "")

	dla := NewDLASystem(30, 1.2, 1.7, 5, seed, false)
	dla.isRunning = true

	f, err := os.OpenFile(csvFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	i := 1
	for {
		// time.Sleep(1 * time.Millisecond)
		dla.Update()

		if i%100 == 0 {
			if _, err := f.Write([]byte(fmt.Sprintf("%d,%f\n", dla.numberOfParticles, dla.clusterRadius))); err != nil {
				log.Fatal(err)
			}
		}

		// if i%1000 == 0 {
		// dla.DisplayGrid()
		// }
		i++

		if dla.isRunning == false {
			break
		}
	}

	fmt.Println("System " + strconv.Itoa(systemID) + " finished!")

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	dla.PersistGridToFile(gridFilename)
	// dla.DisplayGrid()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup

	i := 0
	for i < 100 {
		wg.Add(1)
		fmt.Printf("Starting system %d\n", i)
		go runSystem(int64(i), &wg, i)
		i++
	}

	wg.Wait()
}
