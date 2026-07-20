# Functions and Return Types

## Overview

Functions are reusable blocks of code that perform a specific task. They can accept input through parameters and optionally return one or more values.

---

## Function Without Return Value

```go
func sayHello() {
	fmt.Println("Hello, Go!")
}
```

---

## Function With Parameters

```go
func greet(name string) {
	fmt.Println("Hello,", name)
}
```

---

## Function With Return Value

```go
func add(a, b int) int {
	return a + b
}
```

---

## Multiple Return Values

```go
func getUser() (string, int) {
	return "Mujahid", 25
}
```

---

## Named Return Values

```go
func rectangle(length, width int) (area int) {
	area = length * width
	return
}
```

---

## Best Practices

- Keep functions focused on a single responsibility.
- Use meaningful function names.
- Keep functions short and readable.
- Return only the values that are needed.
- Prefer simple and reusable functions.

---

## Key Takeaways

- Functions are declared using the `func` keyword.
- Functions can receive parameters.
- Functions can return one or multiple values.
- Go supports named return values.
- Functions help make code reusable and maintainable.

---