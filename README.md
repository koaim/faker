# faker

[![GoDoc](https://godoc.org/github.com/koaim/faker?status.svg)](https://godoc.org/github.com/koaim/faker) [![license](http://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/koaim/faker/main/LICENSE)

A lightweight random value generator **for tests**.

## Installation

```bash
go get github.com/koaim/faker
```

## Why?

I wanted a simple way to populate Go structs with random data for tests. Libraries like [gofakeit](https://github.com/brianvoe/gofakeit) offer much more functionality than I needed, while I only used `gofakeit.Struct()`.

This library focuses on that single use case: generating random values and populating structs, with a small API and no external dependencies.

## Features

- Generic helpers for generating values and applying overrides
- Supports built-in signed and unsigned integers (including `byte`, `rune`, and `uintptr`), floats, complex numbers, strings, and booleans
- Recursively fills structs, slices, arrays, maps, buffered bidirectional channels, and pointers
- Special generators for `time.Time`, `uuid.UUID`, `url.URL`, `net.IP`, `net.IPNet`, `net.IPMask`, `net.HardwareAddr`, and `error`
- Custom generators matched by exact type
- Option validation reports all detected errors together

## Usage

```go
package example_test

import (
	"testing"

	"github.com/koaim/faker"
)

type User struct {
	ID      int
	Name    string
	Email   string
	IsAdmin bool
}

func TestUser(t *testing.T) {
	// Generate a value using the default options.
	u := faker.Make[User]()
	if u.ID == 0 {
		t.Fatal("expected a generated ID")
	}

	// Apply override after generating the value.
	u = faker.MakeAndOverride(func(v *User) {
		v.Email = "test@example.com"
	})

	// Apply functional options to the default options.
	u = faker.Make[User](
		faker.WithIgnoreFields("IsAdmin", "Email"),
		faker.WithStrLen(8),
		faker.WithAllowedRunes([]rune("abc123")),
	)

	// Override generation for an exact type, including nested fields.
	u = faker.Make[User](t, faker.WithCustomFaker(func() int { return 42 }))
}
```

## License

[MIT](LICENSE)
