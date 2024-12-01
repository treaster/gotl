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
}

func (le ListElement[T]) Prev() *ListElement[T] {
	return le.prev
}

func (le ListElement[T]) Next() *ListElement[T] {
	return le.next
}

type LinkedList[T any] interface {
	First() *ListElement[T]
	Last() *ListElement[T]
	Length() int
	Append(value T)
	Prepend(value T)
	Remove(element *ListElement[T])
	DebugString() string
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
		Value: value,
		prev:  ll.last,
		next:  nil,
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
		Value: value,
		prev:  nil,
		next:  ll.first,
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
	if element.prev != nil && element.next != nil {
		element.prev.next = element.next
		element.next.prev = element.prev
	}

	if element.prev == nil {
		ll.first = element.next

		if ll.first != nil {
			ll.first.prev = nil
		}
	}

	if element.next == nil {
		ll.last = element.prev

		if ll.last != nil {
			ll.last.next = nil
		}
	}
	ll.length--
}

func (ll *linkedList[T]) DebugString() string {
	parts := make([]string, 0, ll.Length())
	for item := ll.First(); item != nil; item = item.Next() {
		parts = append(parts, fmt.Sprintf("%p(%p, %p)", item, item.prev, item.next))
	}

	return strings.Join(parts, ", ")
}
