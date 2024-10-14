# Lesson 8: Conditionals If
Goal: Learn how to control the flow of a program using conditional statements in Go, starting with `if`, `else if`, and `else` clauses.

## 1. The Basic `if` Statement

In Go, the if statement evaluates a condition and executes the code block if the condition is true.

There is no need to wrap the condition in parentheses, but the code block must be enclosed in curly braces {}.

**Example:**
```go
Copy code
x := 10
if x > 5 {
    fmt.Println("x is greater than 5")
}
```

**Task**: Write an if statement that checks if a number is positive, and prints a message if it is.

## 2. `if-else` Statement

Use else to define a code block that runs if the if condition is false.

**Example:**
```go
Copy code
x := 3
if x > 5 {
    fmt.Println("x is greater than 5")
} else {
    fmt.Println("x is not greater than 5")
}
```
**Task:** Write an if-else statement that checks if a number is even or odd, and prints the appropriate message.

## 3. `else if` Statement

You can chain multiple conditions together using else if. This allows you to check multiple conditions in sequence.

**Example:**

```go
Copy code
x := 5
if x > 10 {
    fmt.Println("x is greater than 10")
} else if x == 5 {
    fmt.Println("x is equal to 5")
} else {
    fmt.Println("x is less than 5")
}
```

**Task:** Write an if-else if-else statement that checks if a number is negative, zero, or positive.

## 4. Short Variable Declaration in `if`

In Go, you can declare and initialize a variable directly in the if statement. The scope of the variable is limited to the if-else block.

**Example:**

```go
Copy code
if x := 10; x > 5 {
    fmt.Println("x is greater than 5")
}
```

**Task:** Use short variable declaration in an if statement to check if a string has a length greater than 5. Print the string if it does.

## 5. Nested `if` Statements

You can nest if statements inside other if statements to handle more complex conditions.

**Example:**

```go
Copy code
x := 15
if x > 10 {
    if x < 20 {
        fmt.Println("x is between 10 and 20")
    }
}
```

**Task:** Write a nested if statement that checks if a number is within a specific range (e.g., between 50 and 100).

## 6. Using Logical Operators with if

Go supports logical operators like && (AND) and || (OR) to combine multiple conditions in an if statement.

**Example:**

```go
Copy code
age := 20
if age >= 18 && age <= 30 {
    fmt.Println("You are in your twenties")
}
```
**Task:** Write an if statement using logical operators to check if a number is divisible by both 3 and 5.

## 7. The ! (NOT) Operator

The ! operator negates a boolean expression, allowing you to check if a condition is false.

**Example:**

```go
Copy code
x := false
if !x {
    fmt.Println("x is false")
}
```
**Task:** Write an if statement using the ! operator to check if a string is empty.

## 8. Combining if with Booleans

You can use boolean values directly in an if statement, as the condition is already a true or false expression.

**Example:**

```go
Copy code
isValid := true
if isValid {
    fmt.Println("The value is valid")
}
```
**Task:** Write an if statement that checks a boolean flag and prints a message based on its value.

## 9. if Statements Without else

Sometimes, an if statement doesn't need an else or else if. You might just want to check one condition and move on if it's false.

**Example:**

```go
Copy code
age := 25
if age >= 18 {
    fmt.Println("You are an adult")
}
```
**Task:** Write an if statement that checks if a user is logged in (using a boolean flag) and prints a message only if they are logged in.

## 10. Checking Multiple Conditions

You can chain multiple if statements for different conditions, or use else if to check alternatives. This can be useful for validating multiple inputs.

**Example:**

```go
Copy code
age := 16
if age >= 18 {
    fmt.Println("Adult")
} else if age >= 13 {
    fmt.Println("Teenager")
} else {
    fmt.Println("Child")
}
```
**Task:** Write an if-else if statement to categorize an age as "Child", "Teenager", or "Adult" based on its value.

