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
	if pos < 0 || pos > l.Len() -1 {
		return nil
	}
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

// Replace fromElem by toElem
func (l *List) ReplaceAll(fromElem interface{}, toElem interface{}) {
	for k, v := range l.elems {
		if v == fromElem {
			l.elems[k] = toElem
		}
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
	// If the index is out of bound
	if index < 0 || index > l.Len() -1 {
		return
	}
	// We are splitting around the element
	newArray := l.elems[:index] // before the element
	newArray = append(newArray, l.elems[index+1:]...) // after the element
	l.elems = newArray

}

// Remove all the occurrence of element from the arraylist
func (l *List) RemoveAll(elem interface{}) {
	// while the arraylist contains elem, we are removing the first occur
	l.subRemoveStarting(elem, 0)
}

// This sub method is used to browse only one time the arraylist
func (l *List) subRemoveStarting(elem interface{}, start int) {
	if start == l.Len() {
		return
	}
	for i := start; i < l.Len(); i++ {
		if l.elems[i] == elem {
			l.RemoveAtIndex(i)
			l.subRemoveStarting(elem, i)
		}
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