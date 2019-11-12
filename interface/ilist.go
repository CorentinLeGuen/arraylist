package _interface


type List struct {
	elems []interface{}
}

func New() *List {
	l := new(List)
	l.elems = []interface{}{}
	return l
}

func (l *List) Len() int {
	return len(l.elems)
}

func (l *List) GetPosition(elem interface{}) int {
	for k, v := range l.elems {
		if v == elem {
			return k
		}
	}
	return -1
}

func (l *List) GetValue(pos int) interface{} {
	return l.elems[pos]
}

func (l *List) Contains(elem interface{}) bool {
	for _, v := range l.elems {
		if v == elem {
			return true
		}
	}
	return false
}

func (l *List) Add(elem interface{}) {
	l.elems = append(l.elems, elem)
}

func (l *List) AddAll(elems []interface{}) {
	for _, elem := range elems {
		l.Add(elem)
	}
}

func (l *List) RemoveFirst(elem interface{}) {
	for i, val := range l.elems {
		if val == elem {
			newArray := l.elems[:i]
			newArray = append(newArray, l.elems[i+1:]...)
			l.elems = newArray
			return
		}
	}
}

func (l *List) RemoveAll(elem interface{}) {
	for l.Contains(elem) {
		l.RemoveFirst(elem)
	}
}

func (l *List) ToArray() []interface{} {
	return l.elems
}