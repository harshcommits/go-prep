# Not-so-used concepts

This is just to detail out the concepts that came up during prep, and are not used very frequently. They have been written down here for reference at a later point.

## Comparable Types in Go

Types of generics have been detailed out in Go documentation [here](https://go.dev/blog/comparable). The ordered data type is provided in [this library](https://pkg.go.dev/cmp). A short summary has been provided below:

1. **any** - This is denoted as **[T any]** for generic types, . The **any** type is basically an empty interface that can be instantiated for any type. This type does not allow any kind of comparison operators (=, !=, <, >)

2. **Comparable** - This is to declare generic types which can be compared. == and != operators can be used with these; <, > and similar operators can't be embedded however. More details can be found [here](https://go.dev/ref/spec#Comparison_operators)

3. **Ordered** - This is another generic type impelmentation done as [T cmp.Ordered]. This allows for all kinds of comparison. Check the BinarySearch() function for implementation use-case.

## Anonymous functions in Go

This is an example I have encountered a lot when coming to defining goroutines and channels. A lot of those have something similar to this:

```golang

go func() {
    // something here
}()

```

Those are called **anonymous functions** in Golang. [Here](https://jeremybytes.blogspot.com/2021/02/go-golang-anonymous-functions-inlining.html) is an article explaining that concept.

The parentheses at the end are there for passing values for the parameters defined in the function. An actual function using that would look something like this:

```golang

 go func(message string) {
      fmt.Println(message)
    }("hello")

```

## Generic observations

Pointer receiver -> in case you want to change the value of the function after the method scope
Value receiver -> in case you want to change the value only during the functions scope