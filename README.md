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

## Comparison Gotchas

### Struct Comparisons

One can't compare instances of structs with different names.

```go
package main

type Cmp1 struct {
	field string
}

type Cmp2 struct {
	field string
}

func main() {
	cmp1 := Cmp1{field: "a"}
	cmp2 := Cmp2{field: "a"}
    // compiler error
	println("Does Go think types with same memory layout, but different names are equal?", cmp1 == cmp2) 
```

One also can't implicitly compare structs with a variable memory print.

```go
package main

type Cmp struct {
	field1 string
    field2 []string
}

func main() {
	cmp1 := Cmp{field1: "a", field2: []string{"1"}}
	cmp2 := Cmp{field1: "a", field2: []string{"1"}}
    // compiler error
	println("Does Go think types with same memory layout, but different names are equal?", cmp1 == cmp2) 
```

### Interface Comparison

```go
package main
type Intf interface {
    Method() string
}

type cmp struct {
	field1 string
    field2 []string
}

func (i cmp) Method() string {
    return i.field1
}
func NewCmp(field1 string, field2 string) Intf {
    return cmp{field1: field1, field2: []string {field2}}
}

func main() {
	cmp1 := NewCmp("a", "b")
	cmp2 := NewCmp("a", "b")
    // panic: runtime error: comparing uncomparable type main.cmp
	println("Does Go think types with same memory layout, but different names are equal?", cmp1 == cmp2)
}
```

### Zero Value Comparison

It's possible to check if something is equal to the empty value for a given type.

```go
package main

type cmp struct {
	field1 string
}

func main() {
    a := cmp {field1: "a"}
    println(a == (cmp{}))
}
```