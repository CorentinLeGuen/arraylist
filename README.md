# ArrayList

This library is using Go slices and arrays to provide a Java arraylist implementation like.

Don't hesitate to send me an email (leguen.corentin@protonmail.com) or contribute to the project !

## Go Get

## How to use

### Constructor

To create a new arraylist : `New() *List`

### Methods

| Signature | Description |
|-----------|-------------|
| Len() int | Return the length of the arraylist |
| Clear() | Empty the arraylist |
| Clone() *List | Return a copy of the arraylist |
| IndexOf(elem interface{}) int | Return the position of _elem_ in the arraylist |
| GetValue(position int) interface{} | Return the value at the position _position_ |
| Contains(elem interface{}) bool | Show if the value _elem_ is in the arraylist |
| Add(elem interface{}) | Add _elem_ to the tail of the arraylist |
| AddAll(elems []interface{}) | Add all _elems_ to the tail of the arraylist |
| ReplaceAll(fromElem interface{}, toElem interface{}) | Replace all occurrences of an element by another one | 
| RemoveFirst(elem interface{}) | Remove the first occurrence of _elem_ from the arraylist |
| RemoveAtIndex(index int) | Remove the item at the position _index_ |
| RemoveAll(elem interface{}) | Remove all occurrences of _elem_ from the arraylist |
| ToArray() []interface{} | Return the arraylist as an array |
| Equals(o *List) bool | Compare two arraylist |

## Example

```go
package main

import (
	"fmt"
	arrayList "github.com/colegno/arraylist/string"
)

func main() {
	s := arrayList.New()
	s.Add("toto")
	s.Add("titi")
	s.Add("toto")
	s.Add("tutu")
	s.Add("toto")                               // add five strings to the arraylist
	fmt.Printf("%t\n", s.Contains("tata"))      // false
	fmt.Printf("%t\n", s.Contains("toto"))      // true
	s.RemoveFirst("toto")                       // we are removing the first "toto"
	fmt.Printf("%t\n", s.Contains("toto"))      // true
	s.RemoveAll("toto")                         // we are removing every "toto"
	fmt.Printf("%t\n", s.Contains("toto"))      // false
	fmt.Printf("%d\n", s.Len())                 // 2 there is only ["titi", "tutu"] in the arraylist
	s.AddAll([]string{"toto", "tata"})          // adding at the end of the arraylist
	fmt.Printf("%t\n", s.Contains("tata"))      // true
	fmt.Printf("%d\n", s.GetPosition("tutu"))   // 1
	fmt.Printf("%s\n", s.GetValue(2))           // "toto"
}
```

## TODO

- Update folder _/examples_
- Add other methods
