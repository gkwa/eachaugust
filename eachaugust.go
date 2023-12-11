package eachaugust

import (
	"fmt"
	"log/slog"
	"math/rand"
	"time"
)

func Main() int {
	slog.Debug("eachaugust", "test", true)

	run()
	return 0
}

func RandomPositiveIntegersUpTo(max int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		randSeed := time.Now().UnixNano() // Seed with current time for randomness
		randSrc := rand.NewSource(randSeed)
		rng := rand.New(randSrc)

		for i := 0; i <= max; i++ {
			ch <- rng.Intn(max + 1) // Generate a random integer in the range [0, max]
		}
	}()

	return ch
}

func run() {
	max := 10
	randomIntGenerator := RandomPositiveIntegersUpTo(max)

	// Consume the generated random integers
	for randomInt := range randomIntGenerator {
		fmt.Printf("%d ", randomInt)
	}
	fmt.Println()
}
