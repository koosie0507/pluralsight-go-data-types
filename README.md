# pluralsight-go-data-types
Learn about data types in Go

## Type Alias vs Type Declaration

Type aliases copy data and methods. Type declarations copy data and allow enriching methods.

## Conflicts and Embedded Structs

```go
package a

type A struct {
    First string
}

type B struct {
    A
    First string
}
```

- both the field in the embedded and the field declared in the compound
- the `First` from `B` takes precendence when called directly
- to reach `First` from `A`, call it explicitly

## Conflicts and Embedded Interfaces

To solve interface conflicts, simply choose the interface(s) you want to use by explicitly
implementing the conflicting methods.
 
```go
package a

type A interface {
    Method()
}
type B interface {
    Method()
}

type S struct {
    A
    B
}

func (s *S) Method() {
    s.A.Method()
}
```