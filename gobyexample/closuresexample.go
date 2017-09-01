// Go supports [_anonymous functions_](http://en.wikipedia.org/wiki/Anonymous_function),
// which can form <a href="http://en.wikipedia.org/wiki/Closure_(computer_science)"><em>closures</em></a>.
// Anonymous functions are useful when you want to define
// a function inline without having to name it.

package gobyexample

import "fmt"

// This function `intSeq_clos` returns another function, which
// we define anonymously in the body of `intSeq_clos`. The
// returned function _closes over_ the variable `i` to
// form a closure.
func intSeq_clos() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func ClosureExample() {

	// We call `intSeq_clos`, assigning the result (a function)
	// to `nextInt`. This function value captures its
	// own `i` value, which will be updated each time
	// we call `nextInt`.
	nextInt := intSeq_clos()

	// See the effect of the closure by calling `nextInt`
	// a few times.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// To confirm that the state is unique to that
	// particular function, create and test a new one.
	newInts := intSeq_clos()
	fmt.Println(newInts())
}
