package golist

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func verifyIntOrder(t *testing.T, ll *LinkedList[int], expectedOrder []int) {
	require.Equal(t, len(expectedOrder), ll.Length())

	{
		var foundOrder []int
		for item := ll.First(); item != nil; item = item.Next() {
			foundOrder = append(foundOrder, item.Value)
		}
		require.ElementsMatch(t, expectedOrder, foundOrder)
	}

	{
		var foundOrder []int
		for item := ll.Last(); item != nil; item = item.Prev() {
			foundOrder = append(foundOrder, item.Value)
		}
		sort.Reverse(sort.IntSlice(expectedOrder))
		require.ElementsMatch(t, expectedOrder, foundOrder)
		sort.Reverse(sort.IntSlice(expectedOrder))
	}
}

func TestLinkedList(t *testing.T) {
	ll := NewLinkedList[int]()
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
