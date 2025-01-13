# Lesson 6: Strings Package

**Goal:** Learn how to manipulate and process strings in Go using the built-in `strings` package.

## 1. Introduction to the Strings Package

The Go standard library provides a package called `strings` that contains a variety of functions to work with string values.

To use the `strings` package, you need to import it:

```go
import "strings"
```

**Example:**

```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    s := "Hello, World!"
    fmt.Println(strings.ToUpper(s))  // Convert to uppercase
}
```

**Task:** Import the `strings` package and use the `ToUpper` function to convert a string to uppercase. Print the result.

## 2. Commonly Used Functions in the Strings Package

The `strings` package provides a variety of useful functions for string manipulation. Some of the most commonly used functions include:

- `ToUpper(s string) string`: Converts all characters in the string to uppercase.
- `ToLower(s string) string`: Converts all characters in the string to lowercase.
- `TrimSpace(s string) string`: Removes all leading and trailing whitespace from the string.
- `HasPrefix(s, prefix string) bool`: Checks if the string starts with the given prefix.
- `HasSuffix(s, suffix string) bool`: Checks if the string ends with the given suffix.
- `Contains(s, substr string) bool`: Checks if the string contains the given substring.

**Example:**

```go
s := "   Go Language   "
fmt.Println(strings.TrimSpace(s))  // "Go Language"
fmt.Println(strings.Contains(s, "Go"))  // true
fmt.Println(strings.HasPrefix(s, "Go")) // false (due to leading spaces)
```

**Task:** Write a program that takes a string and performs the following:

- Converts it to lowercase.
- Removes leading and trailing whitespace.
- Checks if it contains a specific word (e.g., "go").
- Checks if it starts with a certain prefix (e.g., "Go").

## 3. Splitting and Joining Strings

The `strings.Split` function splits a string into a slice of substrings based on a delimiter.

The `strings.Join` function joins elements of a slice into a single string with a specified separator.

**Example:**

```go
str := "apple,banana,cherry"
fruits := strings.Split(str, ",")
fmt.Println(fruits)  // ["apple", "banana", "cherry"]

result := strings.Join(fruits, " - ")
fmt.Println(result)  // "apple - banana - cherry"
```

**Task:** Write a program that splits a sentence into words and then joins them back together with a different delimiter (e.g., a hyphen).

## 4. Replacing Parts of a String

The `strings.Replace` function replaces occurrences of a substring with a new value. You can specify how many replacements to make, or replace all occurrences using `-1`.

**Example:**

```go
s := "foo bar foo"
result := strings.Replace(s, "foo", "baz", -1)
fmt.Println(result)  // "baz bar baz"
```

**Task:** Write a program that replaces all occurrences of a word in a string with another word.

## 5. Counting Substrings

The `strings.Count` function counts how many times a substring occurs within a string.

**Example:**

```go
s := "hello hello hello"
fmt.Println(strings.Count(s, "hello")) // 3
```

**Task:** Write a program that counts the number of times a certain word appears in a string.

## 6. String Length and Indexing

You can find the length of a string using the built-in `len` function.

The `strings.Index` function returns the position of the first occurrence of a substring, or `-1` if it is not found.

**Example:**

```go
s := "golang"
fmt.Println(len(s))  // 6
fmt.Println(strings.Index(s, "go"))  // 0
```

**Task:** Write a program that finds the length of a string and the index of a specific substring within it.

## 7. String Comparison

Go provides functions for comparing strings:

- `strings.Compare(s1, s2 string) int`: Compares two strings lexicographically. Returns `0` if they are equal, a negative number if `s1` is less than `s2`, and a positive number if `s1` is greater than `s2`.
- `strings.EqualFold(s1, s2 string) bool`: Compares two strings case-insensitively. Returns `true` if they are equal, ignoring case.

**Example:**

```go
fmt.Println(strings.Compare("apple", "banana"))  // -1
fmt.Println(strings.EqualFold("GoLang", "golang"))  // true
```

**Task:** Write a program that compares two strings lexicographically and checks if they are equal, ignoring case.
