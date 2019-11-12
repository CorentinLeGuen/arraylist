package _interface

import (
	"testing"
)

func TestNew(t *testing.T) {
	t.Log("Simple new object")
	l := New()
	if l.elems == nil {
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
	if l.GetPosition("1") != 0 ||
		l.GetPosition(2) != 1 ||
		l.GetPosition('*') != 2 {
		t.Fail()
	}
}

func TestList_GetPosition2(t *testing.T) {
	t.Log("Get the position of a non existing string")
	l := New()
	if l.GetPosition("99") != -1 {
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