package quicksort_test

import (
	"testing"

	"github.com/scaule/DataStructures/algo/quicksort"
	"github.com/stretchr/testify/assert"
)

func TestQuicksearch(t *testing.T) {
	t.Run("it should be able to sort simple example", func(t *testing.T) {
		have := quicksort.Do([]int{5, 4, 3, 2, 1})
		assert.Len(t, have, 5)
		assert.Equal(t, have[0], 1)
		assert.Equal(t, have[1], 2)
		assert.Equal(t, have[2], 3)
		assert.Equal(t, have[3], 4)
		assert.Equal(t, have[4], 5)
	})

	t.Run("it should be able to sort ordered example", func(t *testing.T) {
		have := quicksort.Do([]int{1, 2, 3, 4, 5})
		assert.Len(t, have, 5)
		assert.Equal(t, have[0], 1)
		assert.Equal(t, have[1], 2)
		assert.Equal(t, have[2], 3)
		assert.Equal(t, have[3], 4)
		assert.Equal(t, have[4], 5)
	})

	t.Run("it should be able to sort complex example", func(t *testing.T) {
		have := quicksort.Do([]int{2, 3, 1, 10, 4})
		assert.Len(t, have, 5)
		assert.Equal(t, have[0], 1)
		assert.Equal(t, have[1], 2)
		assert.Equal(t, have[2], 3)
		assert.Equal(t, have[3], 4)
		assert.Equal(t, have[4], 10)
	})
}
