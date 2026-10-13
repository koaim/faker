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

String lengths are measured in runes and must be at least 1. Container bounds
must satisfy `1 <= MinContainersLen <= MaxContainersLen`. Map keys may collide,
so a generated map can have fewer entries than requested.

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
	u := faker.Make[User](t)
	if u.ID == 0 {
		t.Fatal("expected a generated ID")
	}

	// Apply override after generating the value.
	u = faker.MakeAndOverride(func(v *User) {
		v.Email = "test@example.com"
	})

	// Apply functional options to the default options.
	u = faker.Make[User](t,
		faker.WithIgnoreFields("IsAdmin", "Email"),
		faker.WithStrLen(8),
		faker.WithAllowedRunes([]rune("abc123")),
	)

	// Override generation for an exact type, including nested fields.
	u = faker.Make[User](t,
		faker.WithCustomFaker(func() int { return 42 }),
	)
}
```

## Time generation

By default, `time.Time` values are generated within 200 calendar years before
and after now. Configure each side independently:

```go
createdAt := faker.Make[time.Time](t,
	faker.WithTimeYearsBack(5),
	faker.WithTimeYearsForward(0),
)
```

Alternatively, set `opt.Time = faker.TimeOption{YearsBack: 5, YearsForward: 0}`
and pass `opt` to `MakeWithOption`. Values must be non-negative. Bounds include
the entire Unix second; nanoseconds are random. Setting both values to zero
generates a value within the current Unix second. These settings also apply
to nested fields and pointers. The example requires importing `time`.

## Supported types

Built-in primitives and special types match by exact type. Defined primitives
such as `type ID int`, functions, and interfaces other than `error` remain zero
unless a custom generator is provided. Custom generators can return `nil` for
interfaces and pointers.

Only bidirectional channels are supported. Recursive type graphs, such as
`type Node struct { Next *Node }`, require ignoring the recursive field or
providing a custom generator to stop recursion.

## License

[MIT](LICENSE)
