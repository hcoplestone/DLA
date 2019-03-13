package main

import (
	"fmt"
	"sync"
)

func runSystem(seed int64, wg *sync.WaitGroup) {
	dla := NewDLASystem(480, 1.2, 1.7, 10000, seed)
	// dla.DisplayGrid()
	dla.isRunning = true

	// f, err := os.OpenFile("data.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	i := 1
	for {
		// time.Sleep(1 * time.Millisecond)
		dla.Update()

		if i%100 == 0 {
			fmt.Println(seed)
		}

		// if i%1000 == 0 {
		// dla.DisplayGrid()
		// }
		i++

		// if _, err := f.Write([]byte(fmt.Sprintf("%d,%f\n", dla.numberOfParticles, dla.clusterRadius))); err != nil {
		// 	log.Fatal(err)
		// }

		if dla.isRunning == false {
			break
		}
	}

	fmt.Println("\nFinished!")
	wg.Done()

	// if err := f.Close(); err != nil {
	// 	log.Fatal(err)
	// }

	// dla.DisplayGrid()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go runSystem(123, &wg)

	wg.Add(1)
	go runSystem(789, &wg)

	wg.Wait()
}
