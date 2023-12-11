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

type GeneratorOption func(*rand.Rand)

// WithSeed is a functional option to set the seed for random number generation.
func WithSeed(seed int64) GeneratorOption {
	return func(rng *rand.Rand) {
		rng.Seed(seed)
	}
}

func RandomPositiveIntegersUpTo(max int, options ...GeneratorOption) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		seen := make(map[int]bool) // Hash to track seen values

		// Default seed with current time for randomness
		randSrc := rand.NewSource(time.Now().UnixNano())
		rng := rand.New(randSrc)

		// Apply functional options
		for _, option := range options {
			option(rng)
		}

		for {
			randomNumber := rng.Intn(max + 1)
			// Check if the number is already seen, if yes, generate a new one
			for seen[randomNumber] {
				randomNumber = rng.Intn(max + 1)
			}

			// Mark the number as seen
			seen[randomNumber] = true

			ch <- randomNumber // Generate a random integer in the range [0, max]

			// If all possible numbers are seen, exit the loop
			if len(seen) == max+1 {
				break
			}
		}
	}()

	return ch
}

func run() {
	max := 10
	randomIntGenerator := RandomPositiveIntegersUpTo(max, WithSeed(int64(42)))

	// Consume the generated random integers
	for randomInt := range randomIntGenerator {
		fmt.Printf("%d ", randomInt)
	}
	fmt.Println()
}
