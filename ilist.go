package arraylist

// The list structure where you can put anything
type List struct {
	elems []interface{}
}

// Return a new empty list
func New() *List {
	l := new(List)
	l.elems = []interface{}{}
	return l
}

// Length of the list
func (l *List) Len() int {
	return len(l.elems)
}

// Empty the list
func (l *List) Clear() {
	l.elems = []interface{}{}
}

// Duplicate the list
func (l *List) Clone() *List {
	n := New()
	for _, v := range l.elems {
		n.elems = append(n.elems, v)
	}
	return n
}

// Return the index of the first elem find in the arraylist or return -1
func (l *List) IndexOf(elem interface{}) int {
	for k, v := range l.elems {
		if v == elem {
			return k
		}
	}
	return -1
}

// Return a value at a specific position
func (l *List) GetValue(pos int) interface{} {
	return l.elems[pos]
}

// Indicates if the arraylist contains an element
func (l *List) Contains(elem interface{}) bool {
	for _, v := range l.elems {
		if v == elem {
			return true
		}
	}
	return false
}

// Add an element at the end of the list
func (l *List) Add(elem interface{}) {
	l.elems = append(l.elems, elem)
}

// Add elements at the end of the list
func (l *List) AddAll(elems []interface{}) {
	for _, elem := range elems {
		l.Add(elem)
	}
}

// Remove the first occurrence of an element from the arraylist
func (l *List) RemoveFirst(elem interface{}) {
	for i, val := range l.elems {
		if val == elem {
			// We are splitting around the element
			newArray := l.elems[:i] // before the element
			newArray = append(newArray, l.elems[i+1:]...) // after the element
			l.elems = newArray
			return
		}
	}
}

// Remove the element at the index position
func (l *List) RemoveAtIndex(index int) {
	for i, _ := range l.elems {
		if i == index {
			// We are splitting around the element
			newArray := l.elems[:i] // before the element
			newArray = append(newArray, l.elems[i+1:]...) // after the element
			l.elems = newArray
			return
		}
	}
}

// Remove all the occurrence of element from the arraylist
// TODO : optimize because we are using RemoveFirst, we are reading multiple times the beginning of the arraylist
func (l *List) RemoveAll(elem interface{}) {
	// while the arraylist contains elem, we are removing the first occur
	for l.Contains(elem) {
		l.RemoveFirst(elem)
	}
}

// Return the arraylist as an array
func (l *List) ToArray() []interface{} {
	return l.elems
}

// Compare two arraylist and return true if they are the same object or if they have the same size and every
// object compare two by two in the same order are the same
// I have used the same logic as ArrayList.Equals of the Javadoc 8
func (l *List) Equals(o *List) bool {
	if l == o {
		return true
	}
	if l.Len() != o.Len() {
		return false
	}
	for k, _ := range l.elems {
		if l.elems[k] != o.elems[k] {
			return false
		}
	}
	return true
}