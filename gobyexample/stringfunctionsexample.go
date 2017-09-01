// The standard library's `strings` package provides many
// useful string-related functions. Here are some examples
// to give you a sense of the package.

package gobyexample

import s1 "strings"
import "fmt"

// We alias `fmt.Println` to a shorter name as we'll use
// it a lot below.
var p = fmt.Println

func StrFuncExample() {

	// Here's1 a sample of the functions available in
	// `strings`. Since these are functions from the
	// package, not methods on the string object itself,
	// we need pass the string in question as the first
	// argument to the function. You can find more
	// functions in the [`strings`](http://golang.org/pkg/strings/)
	// package docs.
	p("Contains:  ", s1.Contains("test", "es"))
	p("Count:     ", s1.Count("test", "t"))
	p("HasPrefix: ", s1.HasPrefix("test", "te"))
	p("HasSuffix: ", s1.HasSuffix("test", "st"))
	p("Index:     ", s1.Index("test", "e"))
	p("Join:      ", s1.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s1.Repeat("a", 5))
	p("Replace:   ", s1.Replace("foo", "o", "0", -1))
	p("Replace:   ", s1.Replace("foo", "o", "0", 1))
	p("Split:     ", s1.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s1.ToLower("TEST"))
	p("ToUpper:   ", s1.ToUpper("test"))
	p()

	// Not part of `strings`, but worth mentioning here, are
	// the mechanisms for getting the length of a string in
	// bytes and getting a byte by index.
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}

// Note that `len` and indexing above work at the byte level.
// Go uses UTF-8 encoded strings, so this is often useful
// as-is. If you're working with potentially multi-byte
// characters you'll want to use encoding-aware operations.
// See [strings, bytes, runes and characters in Go](https://blog.golang.org/strings)
// for more information.
