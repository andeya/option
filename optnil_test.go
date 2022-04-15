package option

import (
	"fmt"
	"strconv"
)

func ExampleOptnil() {
	type A struct {
		X int
	}
	var a = Ptr(&A{X: 1})
	fmt.Println(a.NotNil(), a.IsNil())

	var b = Nil[A]()
	fmt.Println(b.NotNil(), b.IsNil())

	var x = b.UnwrapOr(&A{X: 2})
	fmt.Println(x)

	type B struct {
		Y string
	}
	var c = OptnilMap(a, func(t *A) *B {
		return &B{
			Y: strconv.Itoa(t.X),
		}
	})
	fmt.Println(c)

	// Output:
	// true false
	// false true
	// &{2}
	// NonNil(&{1})
}
