package gotl_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/treaster/gotl"
)

func TestSet(t *testing.T) {
	s := gotl.NewSet[int]()
	require.Equal(t, 0, s.Len())
	require.False(t, s.Has(5))

	s.Add(5)
	require.Equal(t, 1, s.Len())
	require.True(t, s.Has(5))

	s.Add(5)
	require.Equal(t, 1, s.Len())
	require.True(t, s.Has(5))

	s.Add(7)
	require.Equal(t, 2, s.Len())
	require.True(t, s.Has(5))
	require.True(t, s.Has(7))

	s.Remove(5)
	require.Equal(t, 1, s.Len())
	require.False(t, s.Has(5))
	require.True(t, s.Has(7))

	require.Equal(t, []int{7}, s.Items())

	s.AddMany([]int{11, 14, 16})
	require.True(t, s.Has(7))
	require.True(t, s.Has(11))
	require.True(t, s.Has(14))
	require.True(t, s.Has(16))

	s.RemoveMany([]int{11, 14})
	require.True(t, s.Has(7))
	require.False(t, s.Has(11))
	require.False(t, s.Has(14))
	require.True(t, s.Has(16))
}
