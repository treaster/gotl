package gotl

import (
	"fmt"
	"strings"
)

// TODO(treaster): Write docs

type ListElement[T any] struct {
	Value T

	prev *ListElement[T]
	next *ListElement[T]

	removed bool
}

func (le ListElement[T]) Prev() *ListElement[T] {
	if le.removed {
		panic("Calling Prev() on a removed LinkedList element is undefined and unsupported.")
	}
	return le.prev
}

func (le ListElement[T]) Next() *ListElement[T] {
	if le.removed {
		panic("Calling Next() on a removed LinkedList element is undefined and unsupported.")
	}
	return le.next
}

type LinkedList[T any] interface {
	First() *ListElement[T]
	Last() *ListElement[T]
	Length() int
	Append(value T)
	Prepend(value T)
	Remove(element *ListElement[T])
	DebugString(bool) string
}

func NewLinkedList[T any]() LinkedList[T] {
	return &linkedList[T]{}
}

type linkedList[T any] struct {
	first  *ListElement[T]
	last   *ListElement[T]
	length int
}

func (ll *linkedList[T]) First() *ListElement[T] {
	return ll.first
}

func (ll *linkedList[T]) Last() *ListElement[T] {
	return ll.last
}

func (ll *linkedList[T]) Length() int {
	return ll.length
}

func (ll *linkedList[T]) Append(value T) {
	newItem := &ListElement[T]{
		Value:   value,
		prev:    ll.last,
		next:    nil,
		removed: false,
	}

	if ll.last != nil {
		ll.last.next = newItem
	} else {
		ll.first = newItem
		ll.last = newItem
	}
	ll.last = newItem
	ll.length++
}

func (ll *linkedList[T]) Prepend(value T) {
	newItem := &ListElement[T]{
		Value:   value,
		prev:    nil,
		next:    ll.first,
		removed: false,
	}

	if ll.first != nil {
		ll.first.prev = newItem
	} else {
		ll.first = newItem
		ll.last = newItem
	}

	ll.first = newItem
	ll.length++
}

func (ll *linkedList[T]) Remove(element *ListElement[T]) {
	if element.prev != nil {
		element.prev.next = element.next
	}

	if element.next != nil {
		element.next.prev = element.prev
	}

	if element == ll.first {
		ll.first = element.next
	}

	if element == ll.last {
		ll.last = element.prev
	}

	element.removed = true
	ll.length--
}

func (ll *linkedList[T]) DebugString(withValues bool) string {
	parts := make([]string, 0, ll.Length())
	for item := ll.First(); item != nil; item = item.Next() {
		var part string
		if withValues {
			part = fmt.Sprintf("(%p, %p (%v), %p)", item.prev, item, item.Value, item.next)
		} else {
			part = fmt.Sprintf("(%p, %p, %p)", item.prev, item, item.next)
		}
		parts = append(parts, part)
	}

	return strings.Join(parts, ", ")
}
