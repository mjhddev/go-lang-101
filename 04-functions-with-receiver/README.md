# Functions with Receiver

## Overview

A receiver allows a function to be attached to a specific type. When a function has a receiver, it becomes a **method**.

Methods can only be called by values of the associated type.

---

## Defining a Custom Type

```go
type deck []string
```

`deck` is a custom type whose underlying type is `[]string`.

---

## Creating a Method

```go
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
```

In this example:

- `d` is the receiver.
- `deck` is the receiver type.
- `print()` is a method of `deck`.

---

## Calling a Method

```go
cards := deck{"Ace", "King", "Queen"}

cards.print()
```

Output

```
0 Ace
1 King
2 Queen
```

---

## Function vs Method

Function

```go
func printDeck(d deck) {}
```

Method

```go
func (d deck) print() {}
```

A method is associated with a type, while a function is independent.

---

## Best Practices

- Use methods for behavior related to a specific type.
- Keep methods focused on a single responsibility.
- Choose short receiver names such as `d`, `u`, or `p`.

---

## Key Takeaways

- A receiver turns a function into a method.
- Methods belong to a specific type.
- Custom types can have their own methods.
- Methods make code more organized and readable.

---
