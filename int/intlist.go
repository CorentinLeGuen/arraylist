package int

type List struct {
	elems []int
}

func New() *List {
	l := new(List)
	l.elems = []int{}
	return l
}

func (l *List) Len() int {
	return len(l.elems)
}

func (l *List) GetPosition(elem int) int {
	for k, v := range l.elems {
		if v == elem {
			return k
		}
	}
	return -1
}

func (l *List) GetValue(pos int) int {
	return l.elems[pos]
}

func (l *List) Contains(elem int) bool {
	for _, v := range l.elems {
		if v == elem {
			return true
		}
	}
	return false
}

func (l *List) Add(elem int) {
	l.elems = append(l.elems, elem)
}

func (l *List) AddAll(elems []int) {
	for _, elem := range elems {
		l.Add(elem)
	}
}

func (l *List) RemoveFirst(elem int) {
	for i, val := range l.elems {
		if val == elem {
			newArray := l.elems[:i]
			newArray = append(newArray, l.elems[i+1:]...)
			l.elems = newArray
			return
		}
	}
}

func (l *List) RemoveAll(elem int) {
	for l.Contains(elem) {
		l.RemoveFirst(elem)
	}
}

func (l *List) ToArray() []int {
	return l.elems
}
