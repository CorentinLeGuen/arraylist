package string

type List struct {
	elems []string
}

func New() *List {
	l := new(List)
	l.elems = []string{}
	return l
}

func (l *List) Len() int {
	return len(l.elems)
}

func (l *List) GetPosition(elem string) int {
	for k, v := range l.elems {
		if v == elem {
			return k
		}
	}
	return -1
}

func (l *List) GetValue(pos int) string {
	return l.elems[pos]
}

func (l *List) Contains(elem string) bool {
	for _, v := range l.elems {
		if v == elem {
			return true
		}
	}
	return false
}

func (l *List) Add(elem string) {
	l.elems = append(l.elems, elem)
}

func (l *List) AddAll(elems []string) {
	for _, elem := range elems {
		l.Add(elem)
	}
}

func (l *List) RemoveFirst(elem string) {
	for i, val := range l.elems {
		if val == elem {
			newArray := l.elems[:i]
			newArray = append(newArray, l.elems[i+1:]...)
			l.elems = newArray
			return
		}
	}
}

func (l *List) RemoveAll(elem string) {
	for l.Contains(elem) {
		l.RemoveFirst(elem)
	}
}

func (l *List) ToArray() []string {
	return l.elems
}
