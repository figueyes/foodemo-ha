package _test_test

import (
	"math"
	"testing"
)

func Test_List(t *testing.T) {
	t.Parallel()
	t.Run("When it tries to got abs, it tries abs", func(t *testing.T) {
		got := math.Abs(-1)
		if got != 1 {
			t.Errorf("Abs(-1) = %v; want 1", got)
		}
	})
}
