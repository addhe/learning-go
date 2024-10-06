# Lesson 3: Numbers

**Goal:** Learn how to work with different number types in Go, including integers and floating-point numbers.

## 1. Introduction to Number Types in Go

Go supports various types of numbers:

- **Integers:** Whole numbers, both positive and negative, without fractional parts.
  - `int`, `int8`, `int16`, `int32`, `int64`: Signed integers (allow negative numbers).
  - `uint`, `uint8`, `uint16`, `uint32`, `uint64`: Unsigned integers (only positive numbers).
- **Floating-point numbers:** Numbers with decimal points (fractional parts).
  - `float32`, `float64`: Represents floating-point numbers.

**Example:**

```go
var x int = 10       // An integer
var y float64 = 3.14 // A floating-point number
```

**Task:** Declare variables of type `int` and `float64`, assign them values, and print their values.

## 2. Integer Operations

Go allows basic arithmetic operations on integers, such as:

- Addition (`+`), subtraction (`-`), multiplication (`*`), and division (`/`).
- Integer division discards the remainder (if any).

**Example:**

```go
var a, b int = 15, 4
fmt.Println(a + b) // 19
fmt.Println(a - b) // 11
fmt.Println(a * b) // 60
fmt.Println(a / b) // 3 (integer division)
```

**Task:** Write a Go program that performs all the basic arithmetic operations on two integer variables and prints the results.

## 3. Floating-Point Operations

Floating-point numbers allow you to perform calculations with decimals. Go uses `float32` and `float64` to store these values.

Division between floating-point numbers preserves the fractional part.

**Example:**

```go
var c, d float64 = 7.5, 2.3
fmt.Println(c + d) // 9.8
fmt.Println(c - d) // 5.2
fmt.Println(c * d) // 17.25
fmt.Println(c / d) // 3.26
```

**Task:** Write a Go program that performs floating-point arithmetic on two `float64` variables, including division.

## 4. Type Conversion Between Numbers

Sometimes, you need to convert between different number types (e.g., from `int` to `float64`). Go requires explicit type conversion because it doesn’t do it automatically.

**Example:**

```go
var e int = 10
var f float64 = float64(e) // Convert int to float64
fmt.Println(f)             // 10.0
```

**Task:** Write a Go program that converts an integer to a floating-point number and performs a division operation.

## 5. Working with Constants

In Go, you can define constants using the `const` keyword. Constants are variables whose values cannot be changed after they are set.

**Example:**

```go
const pi float64 = 3.14159
fmt.Println(pi)
```

**Task:** Define a constant for the value of pi (3.14159) and use it to calculate the circumference of a circle with a given radius.

## 6. The math Package

Go provides a `math` package for performing more complex mathematical operations, such as square roots, exponentiation, and trigonometric functions.

**Example:**

```go
import "math"

var g float64 = 16
fmt.Println(math.Sqrt(g)) // 4 (square root)
```

**Task:** Write a program that calculates the square root of a number using the `math.Sqrt` function.
