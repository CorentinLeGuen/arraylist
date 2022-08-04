package arraylist

// List : The list structure where you can put anything
type List struct {
	objects []interface{}
}

// New : Return a new empty list
func New() *List {
	l := new(List)
	l.objects = []interface{}{}
	return l
}

// Len : Length of the list
func (l *List) Len() int {
	return len(l.objects)
}

// Clear : Empty the list
func (l *List) Clear() {
	l.objects = []interface{}{}
}

// Clone : Duplicate the list
func (l *List) Clone() *List {
	n := New()
	for _, v := range l.objects {
		n.objects = append(n.objects, v)
	}
	return n
}

// Sublist : Return the sub list between two ranges
// With x the fromRange and y the toRange, we are sending [x,y] (full inclusive) sublist
func (l *List) Sublist(fromRange int, toRange int) *List {
	if fromRange > toRange || fromRange < 0 || toRange > l.Len()-1 {
		return nil
	}
	nl := New()
	nl.AddAll(l.objects[fromRange : toRange+1])
	return nl
}

// IndexOf : Return the index of the first elem find in the arraylist or return -1
func (l *List) IndexOf(elem interface{}) int {
	for k, v := range l.objects {
		if v == elem {
			return k
		}
	}
	return -1
}

// GetValue : Return a value at a specific position
func (l *List) GetValue(pos int) interface{} {
	if pos < 0 || pos > l.Len()-1 {
		return nil
	}
	return l.objects[pos]
}

// Contains : Indicates if the arraylist contains an element
func (l *List) Contains(elem interface{}) bool {
	for _, v := range l.objects {
		if v == elem {
			return true
		}
	}
	return false
}

// Add : Add an element at the end of the list
func (l *List) Add(elem interface{}) {
	l.objects = append(l.objects, elem)
}

// AddAll : Add elements at the end of the list
func (l *List) AddAll(elems []interface{}) {
	for _, elem := range elems {
		l.Add(elem)
	}
}

// ReplaceAll : Replace fromElem by toElem
func (l *List) ReplaceAll(fromElem interface{}, toElem interface{}) {
	for k, v := range l.objects {
		if v == fromElem {
			l.objects[k] = toElem
		}
	}
}

// RemoveFirst : Remove the first occurrence of an element from the arraylist
func (l *List) RemoveFirst(elem interface{}) {
	for i, val := range l.objects {
		if val == elem {
			// We are splitting around the element
			newArray := l.objects[:i]                       // before the element
			newArray = append(newArray, l.objects[i+1:]...) // after the element
			l.objects = newArray
			return
		}
	}
}

// RemoveAtIndex : Remove the element at the index position
func (l *List) RemoveAtIndex(index int) {
	// If the index is out of bound
	if index < 0 || index > l.Len()-1 {
		return
	}
	// We are splitting around the element
	newArray := l.objects[:index]                       // before the element
	newArray = append(newArray, l.objects[index+1:]...) // after the element
	l.objects = newArray

}

// RemoveAll : Remove all the occurrence of element from the arraylist
func (l *List) RemoveAll(elem interface{}) {
	// while the arraylist contains elem, we are removing the first occur
	l.SubRemoveStarting(elem, 0)
}

// SubRemoveStarting : This sub method is used to browse only one time the arraylist
func (l *List) SubRemoveStarting(elem interface{}, start int) {
	if start == l.Len() {
		return
	}
	for i := start; i < l.Len(); i++ {
		if l.objects[i] == elem {
			l.RemoveAtIndex(i)
			l.SubRemoveStarting(elem, i)
		}
	}
}

// ToArray : Return the arraylist as an array
func (l *List) ToArray() []interface{} {
	return l.objects
}

// Equals : Compare two arraylist and return true if they are the same object or if they have the same size and every
// object compare two by two in the same order are the same
// I have used the same logic as ArrayList.Equals of the Javadoc 8
func (l *List) Equals(o *List) bool {
	if l == o {
		return true
	}
	if l.Len() != o.Len() {
		return false
	}
	for k := range l.objects {
		if l.objects[k] != o.objects[k] {
			return false
		}
	}
	return true
}
