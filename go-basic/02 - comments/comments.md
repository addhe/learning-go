# Lesson 2: Comments

**Goal:** Learn how to effectively use comments in Go to document and clarify your code.

## 1. What Are Comments?

Comments are lines in the source code that are not executed by the program. They are used to describe what the code does, making it easier to understand for yourself and others.

In Go, there are two types of comments:

- **Single-line comments:** Begin with `//`
- **Multi-line comments:** Enclosed between `/* */`

**Example:**

```go
// This is a single-line comment

/*
This is a
multi-line comment
*/
```

## 2. Best Practices for Writing Comments

- **Clarity:** Use comments to explain the why behind the code, not just the what. Avoid over-commenting trivial code.
- **Descriptive function and variable documentation:** It's helpful to describe the purpose of your functions and variables, especially in larger projects.

**Example:**

```go
// CalculateArea computes the area of a rectangle
func CalculateArea(length, width float64) float64 {
    return length * width
}
```

**Task:** Write a Go program that calculates the area of a circle. Add comments to describe the code and explain the logic behind each step.

## 3. Package-Level Comments

In Go, it's a common practice to add package-level comments to describe the purpose of the entire package. These comments are usually placed before the `package` keyword.

**Example:**

```go
// Package geometry provides basic geometric calculations.
package geometry
```

**Task:** Write a new Go file for a simple package (e.g., package `mathutils`) and add a package-level comment describing its purpose.

## 4. Documentation with `go doc`

Go supports documentation generation from comments using the `go doc` tool. By writing descriptive comments for your functions and types, you can generate useful documentation for your project.

**Example:** Run the following command in your terminal to see the documentation for the `fmt` package:

```bash
go doc fmt
```

**Task:** Create comments for each function in your program and use `go doc` to generate documentation for them.

## 5. Commenting for Public APIs

In Go, comments are especially important when writing code that others will use. If you are developing a library or exposing an API, your exported functions and types should always have comments describing their functionality.

**Example:**

```go
// Add takes two integers and returns their sum.
func Add(x, y int) int {
    return x + y
}
```

**Task:** Create a public function that performs a mathematical operation (e.g., multiplication), and write clear comments for it. Use `go doc` to verify the generated documentation.

## Next Steps:

Now that you understand how to use comments effectively, you're ready to explore Numbers and Arithmetic Operators in Go, where you'll work with various data types and perform calculations.

This prompt will help you develop the habit of writing clear and useful comments in your Go programs. As you proceed to more advanced topics, keeping your code well-documented will become crucial in maintaining readability and collaboration.