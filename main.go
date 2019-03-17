package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func runSystem(seed int64, wg *sync.WaitGroup, filename string) {
	dla := NewDLASystem(480, 1.2, 1.7, 10000, seed)
	// dla.DisplayGrid()
	dla.isRunning = true

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	i := 1
	for {
		// time.Sleep(1 * time.Millisecond)
		dla.Update()

		if i%10 == 0 {
			fmt.Printf(".")
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

	fmt.Println("\nFinished!")

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	wg.Done()

	// dla.DisplayGrid()
}

func main() {
	var wg sync.WaitGroup

	var filenameComponents []string
	var filename string

	i := 0
	for i < 100 {
		filenameComponents = []string{"results/ensemble", strconv.Itoa(i), ".csv"}
		wg.Add(1)
		fmt.Printf("Starting system %d\n", i)
		filename = strings.Join(filenameComponents, "")
		go runSystem(int64(i), &wg, filename)
		i++
	}

	wg.Wait()
}
