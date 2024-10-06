# Lesson 5: Booleans

**Goal:** Learn how to work with Boolean values in Go, which are used for logical operations and control flow.

## 1. Introduction to Booleans

In Go, the boolean data type represents values that can be either `true` or `false`.

Boolean expressions are often used in conditions for control flow (like `if` statements) and logical comparisons.

**Example:**

```go
var isGoFun bool = true
fmt.Println(isGoFun)  // true
```

**Task:** Declare a boolean variable and set its value to either `true` or `false`. Print the value of the boolean variable.

## 2. Logical Operators

Go provides three logical operators to work with booleans:

- `&&` (AND): Returns `true` if both operands are `true`.
- `||` (OR): Returns `true` if at least one of the operands is `true`.
- `!` (NOT): Negates the boolean value.

**Example:**

```go
var a, b bool = true, false
fmt.Println(a && b)  // false (AND)
fmt.Println(a || b)  // true (OR)
fmt.Println(!a)      // false (NOT)
```

**Task:** Write a program that demonstrates the use of `&&`, `||`, and `!` with boolean variables. Print the results.

## 3. Comparison Operators

Comparison operators return a boolean result when comparing values:

- `==`: Equal to
- `!=`: Not equal to
- `<`: Less than
- `<=`: Less than or equal to
- `>`: Greater than
- `>=`: Greater than or equal to

**Example:**

```go
var x, y int = 5, 10
fmt.Println(x == y)  // false
fmt.Println(x != y)  // true
fmt.Println(x < y)   // true
fmt.Println(x > y)   // false
```

**Task:** Write a program that compares two integer variables using all the comparison operators and prints the results.

## 4. Booleans in Control Flow (`if`/`else`)

Booleans are commonly used in control flow with `if` statements to execute code based on certain conditions.

**Example:**

```go
var temperature int = 30
if temperature > 25 {
    fmt.Println("It's hot!")
} else {
    fmt.Println("It's cool!")
}
```

**Task:** Write a program that checks if a number is positive, negative, or zero using an `if` statement. Print an appropriate message for each case.

## 5. Combining Logical and Comparison Operators

You can combine logical operators with comparison operators to form complex conditions.

**Example:**

```go
var age int = 20
if age >= 18 && age < 30 {
    fmt.Println("You are a young adult.")
}
```

**Task:** Write a program that checks if a person's age is within a specific range (e.g., between 18 and 60) using both logical and comparison operators.

## 6. Short-Circuiting Behavior

In Go, logical operators `&&` and `||` have short-circuiting behavior:

- In `&&`, if the first condition is `false`, the second condition is not evaluated because the whole expression can never be `true`.
- In `||`, if the first condition is `true`, the second condition is not evaluated because the whole expression is already `true`.

**Example:**

```go
var a, b bool = false, true
fmt.Println(a && b)  // false (b is not evaluated)
fmt.Println(a || b)  // true (b is evaluated)
```

**Task:** Write a program that demonstrates short-circuiting with `&&` and `||` operators. Add `fmt.Println` statements inside the conditions to show which parts are being evaluated.

## 7. Default Value of Booleans

In Go, when a boolean variable is declared without initialization, its default value is `false`.

**Example:**

```go
var isAvailable bool
fmt.Println(isAvailable)  // false
```

**Task:** Declare a boolean variable without initializing it. Print its value to confirm the default.
