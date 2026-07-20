# Variables

## 📌 Overview

Variables are used to store data in memory. In Go, variables can be declared using the `var` keyword or the short declaration operator `:=`.

---

## Variable Declaration

### Using `var`

```go
package main

import "fmt"

func main() {
	var name string = "Mujahid"

	fmt.Println(name)
}
```

### Type Inference

Go can automatically infer the variable type.

```go
var age = 25
```

### Short Variable Declaration

```go
name := "Gopher"
age := 25
```

> **Note:** `:=` can only be used inside a function.

---

## Multiple Variables

```go
var firstName, lastName string = "John", "Doe"
```

Or

```go
name, age := "John", 25
```

---

## Zero Values

Variables that are declared without an initial value receive a default value.

| Type | Zero Value |
|------|------------|
| int | 0 |
| float64 | 0 |
| bool | false |
| string | "" |
| pointer | nil |

Example:

```go
var number int
var active bool
var text string

fmt.Println(number)
fmt.Println(active)
fmt.Println(text)
```

Output

```
0
false

```

---

## Best Practices

- Use `:=` whenever possible inside functions.
- Use meaningful variable names.
- Follow Go naming conventions (`camelCase`).
- Avoid unnecessary type declarations when Go can infer the type.

---

## What I Learned

- Declaring variables with `var`
- Using short variable declaration (`:=`)
- Type inference
- Declaring multiple variables
- Understanding zero values

---