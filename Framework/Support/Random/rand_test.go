package Random

import (
	"fmt"

	"testing"
)

func TestRandInt(t *testing.T) {
	for i := 0; i < 50; i++ {
		fmt.Println(RandInt(0, 3))
	}
}

func TestRandIntArray(t *testing.T) {
	for i := 0; i < 50; i++ {
		fmt.Println(RandIntArray(0, 2, 10))
	}
}
