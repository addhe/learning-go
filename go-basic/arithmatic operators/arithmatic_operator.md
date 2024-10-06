# Lesson 4: Arithmetic Operators

**Goal:** Learn how to use Go's arithmetic operators to perform basic mathematical operations.

## 1. Introduction to Arithmetic Operators

Arithmetic operators in Go allow you to perform basic mathematical calculations on numbers, including integers and floating-point numbers.

The primary arithmetic operators are:

- `+` for addition
- `-` for subtraction
- `*` for multiplication
- `/` for division
- `%` for modulus (remainder after division)

**Example:**

```go
var a, b int = 10, 3
fmt.Println(a + b) // 13 (addition)
fmt.Println(a - b) // 7 (subtraction)
fmt.Println(a * b) // 30 (multiplication)
fmt.Println(a / b) // 3 (integer division)
fmt.Println(a % b) // 1 (modulus)
```

**Task:** Write a program that performs addition, subtraction, multiplication, division, and modulus on two integers and prints the results.

## 2. Division in Go

- **Integer division:** When dividing two integers, the result will also be an integer, meaning the remainder is discarded.

  **Example:**

  ```go
  var x, y int = 7, 2
  fmt.Println(x / y) // 3 (integer division)
  ```

- **Floating-point division:** If either operand is a floating-point number, Go performs floating-point division, preserving the remainder.

  **Example:**

  ```go
  var m, n float64 = 7.0, 2.0
  fmt.Println(m / n) // 3.5 (floating-point division)
  ```

**Task:** Write a program that demonstrates both integer division and floating-point division.

## 3. Modulus Operator (`%`)

The modulus operator returns the remainder of an integer division. It only works with integers.

**Example:**

```go
var p, q int = 10, 4
fmt.Println(p % q) // 2 (remainder)
```

**Task:** Write a program that uses the modulus operator to find the remainder when one number is divided by another.

## 4. Compound Assignment Operators

Go also supports compound assignment operators, which combine an arithmetic operation with assignment in one step. These include:

- `+=` for addition
- `-=` for subtraction
- `*=` for multiplication
- `/=` for division
- `%=` for modulus

**Example:**

```go
var z int = 10
z += 5  // Equivalent to z = z + 5
fmt.Println(z) // 15
```

**Task:** Write a program that uses all the compound assignment operators to modify the value of a variable.

## 5. Operator Precedence

Just like in math, operator precedence dictates the order in which operations are performed in Go.

The order of precedence is:

- `*`, `/`, `%` (multiplication, division, modulus) are evaluated first.
- `+`, `-` (addition, subtraction) are evaluated after.

**Example:**

```go
var a, b, c int = 2, 3, 5
result := a + b * c  // Multiplication is performed first
fmt.Println(result)  // 17 (not 25)
```

**Task:** Write a program that uses a combination of arithmetic operators and demonstrates operator precedence. Use parentheses to override precedence when necessary.

## 6. Overflow and Underflow in Integer Operations

Be cautious with large integer values. In Go, integers have fixed sizes (`int8`, `int16`, `int32`, `int64`), and exceeding the maximum value causes overflow.

**Example:**

```go
var smallInt int8 = 127
smallInt += 1  // Overflow
fmt.Println(smallInt)  // -128 (wrap-around behavior)
```

**Task:** Experiment with integer overflow by using large numbers and small integer types (e.g., `int8`). Observe the results and how Go handles overflow.
