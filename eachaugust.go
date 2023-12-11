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

		// Default seed with current time for randomness
		randSrc := rand.NewSource(time.Now().UnixNano())
		rng := rand.New(randSrc)

		// Apply functional options
		for _, option := range options {
			option(rng)
		}

		for i := 0; i <= max; i++ {
			ch <- rng.Intn(max + 1) // Generate a random integer in the range [0, max]
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
