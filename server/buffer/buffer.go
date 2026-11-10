package buffer

import (
	"dash/consts"
	"errors"
	"iter"
)

type Item[T any] struct {
	next   *Item[T]
	buffer *Buffer[T]
	value  T
}

func NewItem[T any](value T) *Item[T] {
	return &Item[T]{
		value: value,
	}
}

func (i *Item[T]) Value() T {
	return i.value
}

type Buffer[T any] struct {
	head *Item[T]
	tail *Item[T]
	len  int
}

func (b *Buffer[T]) Queue(item *Item[T]) error {
	if item.buffer != nil || item.next != nil {
		return errors.New("impossível enfileirar mensagem já enfileirada")
	}

	item.buffer = b

	if b.len == 0 {
		b.head = item
		b.tail = item
		b.len++

		return nil
	}

	if b.len >= consts.BUFFER_SIZE {
		removed := b.head
		b.head = b.head.next
		b.tail.next = item
		b.tail = item

		removed.buffer = nil
		removed.next = nil

		return nil
	}

	b.tail.next = item
	b.tail = item
	b.len++
	return nil
}

func (b *Buffer[T]) Dequeue() *Item[T] {
	if b.len == 0 {
		return nil
	}

	dequeued := b.head
	b.head = b.head.next
	b.len--

	if b.head == nil {
		b.tail = nil
	}

	dequeued.next = nil
	dequeued.buffer = nil
	return dequeued
}

func (b *Buffer[T]) All() iter.Seq[T] {
	// yield é o codigo do for
	return func(yield func(T) bool) {
		for item := b.head; item != nil; item = item.next {
			if !yield(item.value) {
				return
			}
		}
	}
}
