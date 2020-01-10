package arraylist

import (
	"testing"
)

func TestNew(t *testing.T) {
	t.Log("Simple new object")
	l := New()
	if l.elems == nil || l.Len() != 0 {
		t.Fail()
	}
}

func TestList_Clear(t *testing.T) {
	t.Log("Clear the List")
	l := New()
	l.Add(1)
	l.Add(10)
	l.Add(100)
	l.Add(1000)
	if l.Len() != 4 {
		t.Fail()
	}
	l.Clear()
	if l.Len() != 0 {
		t.Fail()
	}
}

func TestList_Clone(t *testing.T) {
	t.Log("Clone an arraylist")
	l := New()
	l.Add(1)
	l.Add(10)
	l.Add(100)
	l.Add(1000)
	l2 := l.Clone()
	l.RemoveFirst(10)
	l2.Add(1)
	// We are checking if any modified values for any list is modified for the other list
	if l.GetValue(1) != 100 || l2.GetValue(1) != 10 || l.Len() != 3 || l2.Len() != 5 {
		t.Fail()
	}
}

func TestList_Len(t *testing.T) {
	t.Log("Length of the List")
	l := New()
	l.Add("1")
	l.Add(2)
	l.Add(new(List))
	if l.Len() != 3 {
		t.Fail()
	}
}

func TestList_GetPosition(t *testing.T) {
	t.Log("Get the position of a string")
	l := New()
	l.Add("1")
	l.Add(2)
	l.Add('*')
	if l.IndexOf("1") != 0 ||
		l.IndexOf(2) != 1 ||
		l.IndexOf('*') != 2 {
		t.Fail()
	}
}

func TestList_GetPosition2(t *testing.T) {
	t.Log("Get the position of a non existing string")
	l := New()
	if l.IndexOf("99") != -1 {
		t.Fail()
	}
}

func TestList_GetValue(t *testing.T) {
	t.Log("Get the value on a position")
	l := New()
	l.Add("1")
	l.Add(2)
	l.Add('*')
	if l.GetValue(0) != "1" ||
		l.GetValue(1) != 2 ||
		l.GetValue(2) != '*' {
		t.Fail()
	}
}

func TestList_Contains(t *testing.T) {
	t.Log("List contains string")
	l := New()
	l.Add("1")
	l.Add(2)
	l.Add('*')
	if !l.Contains("1") ||
		!l.Contains(2) ||
		!l.Contains('*') ||
		l.Contains("8") {
		t.Fail()
	}
}

func TestList_Add(t *testing.T) {
	t.Log("Adding few elements")
	l := New()
	l.Add("1")
	l.Add(2)
	l.Add('*')
	if len(l.elems) != 3 {
		t.Fail()
	}
}

func TestList_AddAll(t *testing.T) {
	t.Log("Adding few elements from array")
	var s []interface{}
	s = append(s, "1")
	s = append(s, 2)
	s = append(s, '*')
	l := New()
	l.AddAll(s)
	if len(l.elems) != 3 {
		t.Fail()
	}
}

func TestList_AddAll2(t *testing.T) {
	t.Log("Adding no elements from array")
	var s []interface{}
	l := New()
	l.AddAll(s)
	if len(l.elems) != 0 {
		t.Fail()
	}
}

func TestList_RemoveFirst(t *testing.T) {
	t.Log("Removing the first occurrence")
	l := New()
	l.Add("1")
	l.Add(2)
	l.Add('*')
	l.RemoveFirst("1")
	if l.Contains("1") || l.Len() != 2 {
		t.Fail()
	}
}

func TestList_RemoveAtIndex(t *testing.T) {
	t.Log("Removing an occurrence by the position")
	l := New()
	l.Add("1")
	l.Add(2)
	l.Add('*')
	l.RemoveAtIndex(1)
	if l.Len() != 2 || l.Contains(2) {
		t.Fail()
	}
}

func TestList_RemoveAll(t *testing.T) {
	t.Log("Removing all the occurrences")
	l := New()
	l.Add("1")
	l.Add("2")
	l.Add('*')
	l.Add('*')
	l.Add('*')
	l.Add('*')
	l.Add('*')
	l.Add('*')
	l.RemoveAll('*')
	if l.Contains('*') || l.Len() != 2 {
		t.Fail()
	}
}

func TestList_ToArray(t *testing.T) {
	t.Log("Getting the array")
	var s []interface{}
	s = append(s, "1")
	s = append(s, 2)
	s = append(s, '*')
	l := New()
	l.AddAll(s)
	s = l.ToArray()
	for k, v := range s {
		if l.GetValue(k) != v {
			t.Fail()
		}
	}
}

func TestList_Equals(t *testing.T) {
	l1 := New()
	l1.Add(1)
	if !l1.Equals(l1) {
		t.Fail()
	}
	l2 := New()
	l2.Add(1)
	if !l1.Equals(l2) {
		t.Fail()
	}
	l1.Add(2)
	if l1.Equals(l2) {
		t.Fail()
	}
	l2.Add(2)
	if !l1.Equals(l2) {
		t.Fail()
	}
	l1.Add(3)
	l2.Add(4)
	if l1.Equals(l2) {
		t.Fail()
	}
}

func TestList_RemoveAll_With_OutOfBound(t *testing.T) {
	l := New()
	l.Add("A")
	l.Add("B")
	l.Add("C")
	l.Add("D")
	l.Add("B")
	l.Add("C")
	l.Add("A")
	l.Add("E")
	l.Add("A")
	l.Add("A")
	l.RemoveAll("A")
	if l.Contains("A") {
		t.Fail()
	}
	l.RemoveAll("D")
	if l.Contains("D") {
		t.Fail()
	}
	if l.Len() != 5 {
		t.Fail()
	}
	l.RemoveAtIndex(-1)
	l.RemoveAtIndex(l.Len())
	if l.Len() != 5 {
		t.Fail()
	}
}

func TestList_RemoveAll_OnlyOneObject(t *testing.T) {
	l := New()
	l.Add("X")
	l.Add("X")
	l.Add("X")
	l.Add("X")
	l.Add("X")
	l.Add("X")
	l.RemoveAll("X")
	if l.Len() != 0 {
		t.Fail()
	}
	l.Add("Y")
	l.RemoveAll("Y")
	if l.Contains("Y") {
		t.Fail()
	}
}

func TestList_OutOfBound(t *testing.T) {
	l := New()
	l.Add(1)
	l.Add("2")
	l.Add(3)
	if l.IndexOf("7") != -1 {
		t.Fail()
	}
	if l.IndexOf(nil) != -1 {
		t.Fail()
	}
	if l.GetValue(-1) != nil || l.GetValue(l.Len()) != nil {
		t.Fail()
	}
	l.AddAll(nil)
	if l.Len() != 3 {
		t.Fail()
	}
}