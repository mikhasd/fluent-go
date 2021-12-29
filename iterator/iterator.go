package iterator

import "github.com/mikhasd/fluent"

type Iterator[T any] interface {
	Next() fluent.Option[T]
}

type Iterable[T any] interface {
	Iterator() Iterator[T]
}

type FuncIterator[T any] func() fluent.Option[T]

func (fi FuncIterator[T]) Next() fluent.Option[T] {
	return fi()
}

func Func[T any](fn func() fluent.Option[T]) Iterator[T] {
	return FuncIterator[T](fn)
}

type Sized interface {
	Size() fluent.Option[int]
}

func Size[T any](it Iterator[T]) fluent.Option[int] {
	if withSize, ok := it.(Sized); ok {
		return withSize.Size()
	} else {
		return fluent.Empty[int]()
	}
}

func FromArray[T any](elements []T) Iterator[T] {
	if len(elements) == 0 {
		return &emptyIterator[T]{}
	} else if len(elements) == 1 {
		return &singleItemIterator[T]{
			item:     elements[0],
			consumed: false,
		}
	} else {
		return &arrayIterator[T]{
			data:  elements,
			index: 0,
		}
	}
}

func Of[T any](elements ...T) Iterator[T] {
	return FromArray(elements)
}
