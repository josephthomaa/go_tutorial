package main

import (
    "fmt"
    "time"
)

func sleep() {
    time.Sleep(1 * time.Second)
    fmt.Println("Slept for 1 second")
}

func main() {
	start := time.Now()

    // Create a channel to synchronize the goroutines
    done := make(chan bool)

    // Start 10 goroutines to call the `sleep` function
    for i := 0; i < 20; i++ {
        go func() {
            sleep()
            // Signal that the goroutine is done
            done <- true
        }()
    }

    // Wait for all goroutines to finish
    for i := 0; i < 20; i++ {
        <-done
    }

    fmt.Println("All goroutines finished")
	elapsed := time.Since(start)
fmt.Printf("Total time taken: %s\n", elapsed)
}
