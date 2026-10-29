# faker

[![GoDoc](https://godoc.org/github.com/koaim/faker?status.svg)](https://godoc.org/github.com/koaim/faker) [![license](http://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/koaim/faker/main/LICENSE)
[![codecov](https://codecov.io/gh/koaim/faker/branch/main/graph/badge.svg)](https://codecov.io/gh/koaim/faker)

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
- Special generators for `time.Time`, `uuid.UUID`, `url.URL`, `net.IP`, `net.IPNet`, `net.IPMask`, `net.HardwareAddr` and `error`
- Custom generators matched by exact type

**Warning:** Do not use this library with recursive types (such as trees, graphs, or linked lists). Generation has no recursion limit or cycle detection and can recurse indefinitely, causing a stack overflow.

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

	// Apply overrides after generating the value.
	u = faker.MakeWithOverrides(func(v *User) {
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
		faker.WithBuilder(func() int { return 42 }),
	)
}
```
## Options

Default values returned by `DefaultOption()`:

| Field | Description | Default value |
| --- | --- | --- |
| `MinContainersLen` | Minimum requested length of slices, maps, and channels | `1` |
| `MaxContainersLen` | Maximum length of slices, maps, and channels | `3` |
| `StrLen` | Length of generated strings in runes | `6` |
| `AllowedRunes` | Characters used to generate random strings | `[]rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789123456789")` |
| `IgnoreFields` | Names of struct fields to skip | `nil` |
| `CloseChannels` | Whether to close generated channels after filling them | `false` |
| `CustomFakers` | Custom generators matched by exact type, overriding default generation | `nil` |
| `Url.QueryCount` | Number of query parameters to generate for URLs | `2` |
| `Url.WithUserInfo` | Whether generated URLs include a username and password | `true` |
| `Url.AllowedDomainZones` | Domain zones used for URL hosts | `[]string{"com", "net", "org", "en", "site"}` |
| `Url.AllowedSchemes` | Schemes used for generated URLs | `[]string{"http", "https"}` |
| `Url.PathLen` | Length of the generated URL path in runes | `3` |
| `Time.YearsBack` | Number of calendar years before now included in the time range | `200` |
| `Time.YearsForward` | Number of calendar years after now included in the time range | `200` |

## License

[MIT](LICENSE)
