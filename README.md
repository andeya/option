# option [![Docs](https://img.shields.io/badge/Docs-pkg.go.dev-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/henrylee2cn/option)

Go-generics option module inspired by rust.

Avoid `nil` value, handle value with `Option` type, will you choose her?

## Go Version

go≥1.18

## Example

- Option

```go
func ExampleOption() {
	type A struct {
		X int
	}
	var a = Some(A{X: 1})
	fmt.Println(a.IsSome(), a.IsNone())

	var b = None[A]()
	fmt.Println(b.IsSome(), b.IsNone())

	var x = b.UnwrapOr(A{X: 2})
	fmt.Println(x)

	type B struct {
		Y string
	}
	var c = Map(a, func(t A) B {
		return B{
			Y: strconv.Itoa(t.X),
		}
	})
	fmt.Println(c)

	// Output:
	// true false
	// false true
	// {2}
	// Some({1})
}
```

- Optnil

```go
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
```