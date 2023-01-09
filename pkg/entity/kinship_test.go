package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPop(t *testing.T) {
	t.Run("Pop middle", func(t *testing.T) {
		collection := &KinshipCollection{
			{
				Parent: 1,
				Child:  2,
			},
			{
				Parent: 2,
				Child:  3,
			},
			{
				Parent: 4,
				Child:  2,
			},
		}

		got := collection.Pop(&Kinship{
			Parent: 2,
			Child:  3,
		})

		require.Len(t, *collection, 2)
		require.NotNil(t, got)

	})
}
