# option [![Docs](https://img.shields.io/badge/Docs-pkg.go.dev-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/henrylee2cn/option)

Go-generics option module inspired by rust.

Avoid `nil` value, handle value with `Option` type, will you choose her?

## Go Version

go≥1.18

## Example

```go
func Example() {
	type A struct {
		X int
	}
	var a = Wrap(&A{X: 1})
	fmt.Println(a.IsSome(), a.IsNone())

	var b = None[A]()
	fmt.Println(b.IsSome(), b.IsNone())

	var x = b.UnwrapOr(&A{X: 2})
	fmt.Println(x)

	type B struct {
		Y string
	}
	var c = Map(a, func(t *A) *B {
		return &B{
			Y: strconv.Itoa(t.X),
		}
	})
	fmt.Println(c)

	// Output:
	// true false
	// false true
	// &{2}
	// Some(&{1})
}
```
