# Go Generics: Constraints and Use Cases

## Overview

Generics were introduced in Go 1.18, allowing you to write functions, types, and methods that work with multiple types while maintaining type safety. The key to understanding Go generics is understanding **type constraints**.

## Type Constraints

### 1. `any` (Empty Interface)

**Definition:** Accepts any type.

```go
func PrintValue[T any](value T) {
    fmt.Println(value)
}
```

**Use Cases:**
- When you need complete flexibility with types
- Building utility functions that work with any type
- When only identity matters, not specific operations

**Differences from others:**
- No operations allowed on the value (can't compare, order, etc.)
- Most permissive constraint
- Functionally equivalent to `interface{}`

---

### 2. `comparable`

**Definition:** Only types that support `==` and `!=` operators.

```go
func FindIndex[T comparable](slice []T, value T) int {
    for i, v := range slice {
        if v == value {
            return i
        }
    }
    return -1
}
```

**Types that implement `comparable`:**
- All integers, floats, strings, bools
- Arrays (if element type is comparable)
- Structs (if all fields are comparable)
- Pointers
- **NOT**: slices, maps, functions

**Use Cases:**
- Searching for values in collections
- Using values as map keys in generic code
- Testing equality operations
- Deduplication logic

**Example:**
```go
func Contains[T comparable](slice []T, item T) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}

// Usage works with comparable types
Contains([]int{1, 2, 3}, 2)           // ✓ true
Contains([]string{"a", "b"}, "a")    // ✓ true
Contains([][]int{{1}, {2}}, []int{1}) // ✗ compile error: slices not comparable
```

---

### 3. `Ordered` (Custom Constraint)

**Definition:** A custom constraint from `golang.org/x/exp/constraints` for types that support ordering.

```go
package constraints

type Ordered interface {
    Integer | Float | ~string
}

type Integer interface {
    Signed | Unsigned
}

type Signed interface {
    int | int8 | int16 | int32 | int64
}

type Unsigned interface {
    uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

type Float interface {
    float32 | float64
}
```

**Types that implement `Ordered`:**
- All signed integers: `int`, `int8`, `int16`, `int32`, `int64`
- All unsigned integers: `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`
- Floating-point types: `float32`, `float64`
- Strings: `string`

**Use Cases:**
- Finding minimum/maximum values
- Sorting algorithms
- Range comparisons
- Implementing comparison-based data structures

**Example:**
```go
import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

func Max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

// Usage
Min(5, 3)              // ✓ 3
Max(3.14, 2.71)        // ✓ 3.14
Max("apple", "banana") // ✓ "banana"
```

---

## Comparison Table

| Constraint | Supports `==` | Supports `<`, `>` | Can be map key | Use Case |
|-----------|:-------------:|:-----------------:|:--------------:|----------|
| `any` | ✗ | ✗ | ✗ | Maximum flexibility |
| `comparable` | ✓ | ✗ | ✓ | Equality checks, dedup, maps |
| `Ordered` | ✓ (implied) | ✓ | ✓ | Sorting, min/max, ranges |

---

## Advanced: Custom Constraints

You can define your own constraints by combining existing ones:

```go
// Constraint for numeric types only
type Number interface {
    constraints.Integer | constraints.Float
}

func Sum[T Number](values []T) T {
    var sum T
    for _, v := range values {
        sum += v
    }
    return sum
}

// Constraint for types with a String() method (Stringer interface)
type Stringer interface {
    String() string
}

func PrintAll[T Stringer](items []T) {
    for _, item := range items {
        fmt.Println(item.String())
    }
}
```

---

## Practical Examples

### Example 1: Generic Cache (Comparable)
```go
type Cache[K comparable, V any] struct {
    data map[K]V
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
    val, ok := c.data[key]
    return val, ok
}

func (c *Cache[K, V]) Set(key K, value V) {
    c.data[key] = value
}
```

### Example 2: Generic Linked List (Any)
```go
type Node[T any] struct {
    Value T
    Next  *Node[T]
}

type LinkedList[T any] struct {
    head *Node[T]
}

func (ll *LinkedList[T]) Push(value T) {
    ll.head = &Node[T]{Value: value, Next: ll.head}
}
```

### Example 3: Generic Sorting Function (Ordered)
```go
func BubbleSort[T constraints.Ordered](slice []T) {
    n := len(slice)
    for i := 0; i < n; i++ {
        for j := 0; j < n-i-1; j++ {
            if slice[j] > slice[j+1] {
                slice[j], slice[j+1] = slice[j+1], slice[j]
            }
        }
    }
}
```

---

## Key Takeaways

1. **`any`** = Maximum flexibility, no operations allowed
2. **`comparable`** = Equality operations only, can be map keys
3. **`Ordered`** = Full comparison operators, includes comparable types
4. **Custom constraints** = Combine interfaces using `|` (union type)
5. Always choose the **least permissive** constraint that meets your needs

---

## References

- Go Generics: https://go.dev/blog/intro-generics
- Type Parameters: https://go.dev/doc/tutorial/generics
- `golang.org/x/exp/constraints`: https://pkg.go.dev/golang.org/x/exp/constraints
