package sequencecompare_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/james-d-elliott/challenges/sequencecompare"
)

func TestFindOneMissingArrayValues(t *testing.T) {
	t.Run("Ordered", func(t *testing.T) {
		assert.Equal(t, 23, sequencecompare.FindOneMissingSliceValue(full, missingOneOrdered))
	})

	t.Run("OrderedLarge", func(t *testing.T) {
		assert.Equal(t, 143, sequencecompare.FindOneMissingSliceValue(fullLarge, missingOneLargeOrdered))
	})

	t.Run("Unordered", func(t *testing.T) {
		assert.Equal(t, 23, sequencecompare.FindOneMissingSliceValue(full, missingOneUnordered))
	})

	t.Run("UnorderedLarge", func(t *testing.T) {
		assert.Equal(t, 143, sequencecompare.FindOneMissingSliceValue(fullLarge, missingOneLargeUnordered))
	})
}

func TestFindTwoMissingArrayValues(t *testing.T) {
	t.Run("Ordered", func(t *testing.T) {
		assert.Equal(t, []int{21, 66}, sequencecompare.FindTwoMissingSliceValues(full, missingTwoOrdered))
	})

	t.Run("OrderedLarge", func(t *testing.T) {
		assert.Equal(t, []int{343, 4359}, sequencecompare.FindTwoMissingSliceValues(fullLarge, missingTwoLargeOrdered))
	})

	t.Run("Unordered", func(t *testing.T) {
		assert.Equal(t, []int{21, 66}, sequencecompare.FindTwoMissingSliceValues(full, missingTwoUnordered))
	})

	t.Run("UnorderedLarge", func(t *testing.T) {
		assert.Equal(t, []int{343, 4359}, sequencecompare.FindTwoMissingSliceValues(fullLarge, missingTwoLargeUnordered))
	})
}

func BenchmarkFindOneMissingArrayValue(b *testing.B) {
	b.Run("Ordered", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			sequencecompare.FindOneMissingSliceValue(full, missingOneOrdered)
		}
	})

	b.Run("Unordered", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			sequencecompare.FindOneMissingSliceValue(full, missingOneUnordered)
		}
	})
}

func BenchmarkFindTwoMissingArrayValuesOrdered(b *testing.B) {
	b.Run("Ordered", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			sequencecompare.FindTwoMissingSliceValues(full, missingTwoOrdered)
		}
	})

	b.Run("Unordered", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			sequencecompare.FindTwoMissingSliceValues(full, missingTwoUnordered)
		}
	})
}
