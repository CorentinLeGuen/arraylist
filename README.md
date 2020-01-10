# ArrayList

This library is using Go slices and arrays to provide a Java arraylist implementation like.

Don't hesitate to send me an email (leguen.corentin@protonmail.com) or contribute to the project !

## Go Get

`go get github.com/colegno/arraylist/int` for the int arraylist.

`go get github.com/colegno/arraylist/string` for the string arraylist.

You can also use _github.com/colegno/arraylist_ for your custom structs.

## How to use

### Constructor

To create a new arraylist : `New() *List`

### Methods

| Signature | Description |
|-----------|-------------|
| Len() int | Return the length of the arraylist |
| Clear() | Empty the arraylist |
| Clone() *List | Return a copy of the arraylist |
| IndexOf(elem <string/int/interface{}>) int | Return the position of _elem_ in the arraylist |
| GetValue(position int) <string/int/interface{}> | Return the value at the position _position_ |
| Contains(elem <string/int/interface{}>) bool | Show if the value _elem_ is in the arraylist |
| Add(elem <string/int/interface{}>) | Add _elem_ to the tail of the arraylist |
| AddAll(elems []<string/int/interface{}>) | Add all _elems_ to the tail of the arraylist |
| RemoveFirst(elem <string/int/interface{}>) | Remove the first occurence of _elem_ from the arraylist |
| RemoveAtIndex(index int) | Remove the item at the position _index_ |
| RemoveAll(elem <string/int/interface{}>) | Remove all occurences of _elem_ from the arraylist |
| ToArray() []<string/int/interface{}> | Return the arraylist as an array |

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

- Update the method `(l *List) RemoveAll(elem interface{})`
- Update folder _/examples_