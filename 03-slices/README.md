# Slices, Append, and For Loops

## Overview

A slice is a dynamic data structure used to store a collection of values. Unlike arrays, slices can grow or shrink as needed.

Go provides the `append()` function to add new elements to a slice. To iterate over slices, Go uses the `for` loop, often combined with the `range` keyword.

---

## Creating a Slice

```go
numbers := []int{10, 20, 30}
```

---

## Appending Elements

```go
numbers = append(numbers, 40)
numbers = append(numbers, 50)
```

Output

```
[10 20 30 40 50]
```

---

## Basic For Loop

```go
for i := 0; i < 5; i++ {
	fmt.Println(i)
}
```

---

## Using `range`

```go
numbers := []int{10, 20, 30}

for index, value := range numbers {
	fmt.Println(index, value)
}
```

Output

```
0 10
1 20
2 30
```

---

## Ignoring the Index

```go
for _, value := range numbers {
	fmt.Println(value)
}
```

Output

```
10
20
30
```

---

## Best Practices

- Use slices instead of arrays in most cases.
- Remember that `append()` returns a new slice.
- Use `range` when iterating through slices.
- Use `_` to ignore unused variables.

---

## Key Takeaways

- A slice is a dynamic collection of elements.
- `append()` adds new elements to a slice.
- Go uses the `for` keyword for all loops.
- `range` makes iterating over slices easier.

---
