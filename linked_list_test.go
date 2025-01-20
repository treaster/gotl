package gotl_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/treaster/gotl"
)

func verifyIntOrder(t *testing.T, ll gotl.LinkedList[int], expectedOrder []int) {
	require.Equal(t, len(expectedOrder), ll.Length())

	{
		foundOrder := []int{}
		for item := ll.First(); item != nil; item = item.Next() {
			foundOrder = append(foundOrder, item.Value)
		}
		require.Equal(t, expectedOrder, foundOrder)
	}

	{
		foundOrder := []int{}
		for item := ll.Last(); item != nil; item = item.Prev() {
			foundOrder = append(foundOrder, item.Value)
		}
		slices.Reverse(expectedOrder)
		require.Equal(t, expectedOrder, foundOrder)
		slices.Reverse(expectedOrder)
	}
}

func TestLinkedList(t *testing.T) {
	ll := gotl.NewLinkedList[int]()
	verifyIntOrder(t, ll, []int{})

	ll.Append(1)
	verifyIntOrder(t, ll, []int{1})

	ll.Append(2)
	verifyIntOrder(t, ll, []int{1, 2})

	ll.Prepend(3)

	verifyIntOrder(t, ll, []int{3, 1, 2})

	ll.Prepend(4)
	verifyIntOrder(t, ll, []int{4, 3, 1, 2})

	item := ll.First().Next()
	ll.Remove(item)
	verifyIntOrder(t, ll, []int{4, 1, 2})

	item = ll.First()
	ll.Remove(item)
	verifyIntOrder(t, ll, []int{1, 2})

	item = ll.Last()
	ll.Remove(item)
	verifyIntOrder(t, ll, []int{1})

	item = ll.Last()
	ll.Remove(item)
	verifyIntOrder(t, ll, []int{})
}

func TestLinkedList_StaleItems(t *testing.T) {
	ll := gotl.NewLinkedList[int]()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Append(4)

	verifyIntOrder(t, ll, []int{1, 2, 3, 4})
	oldFirst := ll.First()
	ll.Remove(ll.First())
	ll.Remove(ll.First())
	require.Equal(t, 3, ll.First().Value)
	require.Panics(t, func() { _ = oldFirst.Prev() })
	require.Panics(t, func() { _ = oldFirst.Next() })
}
