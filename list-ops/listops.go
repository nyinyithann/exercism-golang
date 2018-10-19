// Package listops includes some list operation functions
package listops

// IntList is of int slice type
type IntList []int
type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(x int) int

// Foldl picks up elements of the list starting from the beginning and passes it into fn as input and
// return the accumulated result
func (l IntList) Foldl(fn binFunc, init int) int {
	for _, v := range l {
		init = fn(init, v)
	}
	return init
}

// Foldr picks up elements of the list starting from the end and passes it into fn as input and
// return the accumulated result
func (l IntList) Foldr(fn binFunc, init int) int {
	if l.Length() == 0 {
		return init
	}
	for i := l.Length() - 1; i >= 0; i-- {
		init = fn(l[i], init)
	}
	return init
}

// Filter filters a sequence of values based on a predicate
func (l IntList) Filter(fn predFunc) IntList {
	c := l.Length()
	if c == 0 {
		return IntList{}
	}
	r := make(IntList, c)
	i := 0
	for _, v := range l {
		if fn(v) {
			r[i] = v
			i++
		}
	}
	return r[0:i]
}

// Length calculates the number of elements in a list
func (l IntList) Length() int {
	r := 0
	for range l {
		r++
	}
	return r
}

// Map projects each element of a list into a new form
func (l IntList) Map(fn unaryFunc) IntList {
	r := make(IntList, l.Length())
	for i, v := range l {
		r[i] = fn(v)
	}
	return r
}

// Reverse inverts the order of the element of a list
func (l IntList) Reverse() IntList {
	c := l.Length()
	r := make(IntList, c)
	for i, j := 0, c-1; i < c; i, j = i+1, j-1 {
		r[i] = l[j]
	}
	return r
}

// Append inserts another list to the end of source list
func (l IntList) Append(other IntList) IntList {
	sl := l.Length()
	r := make(IntList, sl+other.Length())
	for i, v := range l {
		r[i] = v
	}
	for i, v := range other {
		r[sl+i] = v
	}
	return r
}

// Concat appends other lists to the end of source list
func (l IntList) Concat(lists []IntList) IntList {
	for _, v := range lists {
		l = l.Append(v)
	}
	return l
}
