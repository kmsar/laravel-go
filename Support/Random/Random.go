package Random

import (
	"math/rand"
	"sync"
	"time"
)

func Random() *_Random {
	return &_Random{
		pseudo: rand.New(&source{src: rand.NewSource(time.Now().UnixNano())}),
		mr:     &sync.Mutex{},
	}
}

func (r *_Random) RandoFloat() float64 {
	return r.pseudo.Float64()
}

// Int // Int returns a non-negative pseudo-random int.
func (r *_Random) Int() int {
	return r.pseudo.Int()
}

// Duration D returns a random time.Duration between min and max: [min, max].
func (r *_Random) Duration(min, max time.Duration) time.Duration {
	multiple := int64(1)
	if min != 0 {
		for min%10 == 0 {
			multiple *= 10
			min /= 10
			max /= 10
		}
	}
	n := int64(r.Between(int(min), int(max)))
	return time.Duration(n * multiple)
}

// Between N returns a random int between min and max: [min, max].
// The `min` and `max` also support negative numbers.
func (r *_Random) Between(min, max int) int {
	if min >= max {
		return min
	}
	if min >= 0 {
		return r.Intn(max-min+1) + min
	}
	// As `Intn` dose not support negative number,
	// so we should first shift the value to right,
	// then call `Intn` to produce the random number,
	// and finally shift the result back to left.
	return r.Intn(max+(0-min)+1) - (0 - min)
}

// Intn returns, as an int, a non-negative pseudo-random number in [0,n).
func (r *_Random) Intn(n int) int {
	return r.pseudo.Intn(n)
}

// Int63n returns, as an int64, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func (r *_Random) int64(n int64) int64 {
	return r.pseudo.Int63n(n)
}

// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers [0,n).
func (r *_Random) Perm(n int) []int { return r.pseudo.Perm(n) }

// Seed uses the provided seed value to initialize the default Source to a
// deterministic state. If Seed is not called, the generator behaves as if
// seeded by Seed(1).
func (r *_Random) Seed(n int64) { r.pseudo.Seed(n) }

// Shuffle pseudo-randomizes the order of elements.
// n is the number of elements.
// swap swaps the elements with indexes i and j.
func (r *_Random) Shuffle(n int, swap func(i, j int)) { r.pseudo.Shuffle(n, swap) }

type _Random struct {
	pseudo *rand.Rand
	mr     *sync.Mutex
}

type source struct {
	src rand.Source
	mu  sync.Mutex
}

func (s *source) Int63() int64 {
	s.mu.Lock()
	n := s.src.Int63()
	s.mu.Unlock()
	return n
}

func (s *source) Seed(seed int64) {
	s.mu.Lock()
	s.src.Seed(seed)
	s.mu.Unlock()
}
