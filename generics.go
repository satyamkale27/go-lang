package main

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {

	/*

				    S ~[]E: S is a type parameter that represents a slice of elements of type E.
				    The ~ operator indicates that S can be any type that has the underlying type []E.

					E comparable: E is a type parameter that must satisfy the comparable constraint,
				    meaning that values of type E can be compared using the == operator.

					(s S, v E): The function takes two parameters: s of type S (a slice of E) and v of type E (an element to find in the slice).
					int: The function returns an integer, which is the index of the element in the slice.

			       ✅ This function works with both normal slices ([]E) and custom slices (MySlice) without explicitly defining type like this 	ints := []int{1, 2, 3, 4}.
		           S ~[]E allows custom slice types that have []E as their underlying type.


	*/

	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}
