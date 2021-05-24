## Strings in Go

**Introduction**

Strings in **Go** are read-only. When updated a new memory is allocated. Strings are defined between double quotes. Go doesn't support single quotes for strings.

Here are some examples of strings in __Go__

```go
const name = "Amir Saleem"

const address string = "Delhi, India"

country := "India"
```

In Go, strings are immutable, we can safely update them inside functions

```go
const name string = " Amir "

func TrimStr(value string) string {
    // the original value is not updated instead a new 
    // memory location is allocated to the updated value
    value := strings.TrimSpace(value) 
    return value
}
```

**How strings are stored in memory?**

In Go, a string is readonly slice of bytes.

