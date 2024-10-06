# Lesson 7: Strings

**Goal:** Learn how to work with strings in Go, including how to declare, manipulate, and access string elements.

## 1. Declaring and Initializing Strings

In Go, strings are a sequence of bytes representing UTF-8 encoded text.

You can declare a string using double quotes (`"`) for interpreted strings or backticks (`` ` ``) for raw string literals.

**Example:**

```go
// Interpreted string
var s1 string = "Hello, Go!"

// Raw string (preserves newlines and backslashes)
var s2 string = `This is a raw
string literal that spans multiple lines.`

fmt.Println(s1)
fmt.Println(s2)
```

**Task:** Declare both an interpreted string and a raw string. Print their values.

## 2. Accessing String Characters

- Strings in Go are **immutable**, meaning their contents cannot be changed after they are created.
- You can access individual bytes (not characters) in a string using indexing, but keep in mind that each index refers to a byte, not necessarily a character (due to UTF-8 encoding).

**Example:**

```go
s := "GoLang"
fmt.Println(s[0])  // 71 (ASCII value of 'G')
```

To get the actual character, you can convert the byte to a string type:

```go
fmt.Println(string(s[0]))  // G
```

**Task:** Access and print the first character of a string. Then, print its byte value.

## 3. Iterating Over Strings

You can iterate over a string using a `for` loop, either by byte or by character (rune).

When iterating over bytes, you might encounter partial characters for non-ASCII strings. To avoid this, you can use the `range` loop to iterate over runes (which represent Unicode code points).

**Example:**

```go
s := "GoLang"

// Iterate over bytes
for i := 0; i < len(s); i++ {
    fmt.Printf("%c ", s[i])  // G o L a n g
}

// Iterate over runes (characters)
fmt.Println()
for _, r := range s {
    fmt.Printf("%c ", r)  // G o L a n g
}
```

**Task:** Write a program that iterates over a string using both byte and rune iteration, printing each character.

## 4. String Concatenation

You can concatenate (combine) strings using the `+` operator.

For more complex string concatenations or if you need better performance, consider using the `strings.Builder` or the `fmt.Sprintf` function.

**Example:**

```go
s1 := "Go"
s2 := "Lang"
combined := s1 + s2
fmt.Println(combined)  // GoLang
```

**Task:** Concatenate two strings using the `+` operator and print the result.

## 5. Multiline Strings

Use raw string literals (backticks) for multiline strings. This is particularly useful when you need to include newlines, special characters, or large blocks of text without escape sequences.

**Example:**

```go
s := `This is
a multiline string!`
fmt.Println(s)
```

**Task:** Create a multiline string using raw string literals and print it.

## 6. String Length

- You can find the length of a string using the `len` function. This returns the number of bytes in the string, not the number of characters.

**Example:**

```go
s := "GoLang"
fmt.Println(len(s))  // 6
```

**Task:** Write a program that prints the length of a string.

## 7. Slicing Strings

You can create a substring (a part of the original string) by slicing it. Slicing works with byte indices.

**Example:**

```go
s := "GoLang"
sub := s[0:2]  // "Go"
fmt.Println(sub)
```

**Task:** Slice a string to extract the first three characters and print the result.

## 8. Immutable Nature of Strings

Strings in Go are immutable, meaning you cannot modify individual characters or bytes within the string.

To "modify" a string, you must create a new one based on the original.

**Example:**

```go
s := "GoLang"
s = "Modified"
fmt.Println(s)  // "Modified"
```

**Task:** Create a string, attempt to modify it (observe the result), and then create a new string with the modification.

## 9. Converting Between Strings and Byte Slices

Sometimes you may need to convert a string into a byte slice (for example, when working with binary data). You can convert a string to a byte slice and vice versa using `[]byte` and `string` type conversion.

**Example:**

```go
s := "GoLang"
bytes := []byte(s)
fmt.Println(bytes)  // [71 111 76 97 110 103]

s2 := string(bytes)
fmt.Println(s2)  // "GoLang"
```

**Task:** Convert a string to a byte slice and then back to a string. Print both the byte slice and the string.
