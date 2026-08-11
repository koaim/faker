# faker
[![GoDoc](https://godoc.org/github.com/koaim/faker?status.svg)](https://godoc.org/github.com/koaim/faker) [![license](http://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/koaim/faker/main/LICENSE)

A lightweight random value generator **for tests**.

## Installation

```bash
go get github.com/koaim/faker
```

## Why?

I wanted a simple way to populate Go structs with random data for tests. Libraries like [gofakeit](https://github.com/brianvoe/gofakeit) offer much more functionality than I needed, while I only used `gofakeit.Struct()`.

This library focuses on that single use case: generating random values and populating structs, with a small API and no unnecessary dependencies.

## Features

- Convenient API with generics support
- Supports `int`, `uint`, `float`, `complex`, `string`, `bool`, `struct`, `slice`, `array`, `map`, `chan` and pointer types
- Zero dependencies

## Usage

```go
package main

import (
	"fmt"

	"github.com/koaim/faker"
)

type User struct {
	ID      int
	Name    string
	Email   string
	IsAdmin bool
}

func main() {
	// fill struct with random data
	u := faker.Make[User]()

	// fill struct with random data and override specific fields
	u = faker.MakeAndOverride(func(v *User) {
		v.Email = "test@example.com"
	})

	// fill struct and ignore some fields
	u = faker.Make[User](faker.WithIgnoreFields("IsAdmin", "Email"))

	// fill struct with custom options
	u = faker.MakeWithOption[User](faker.Option{
		MinContainersLen: 1,
		MaxContainersLen: 5,
		AllowedRunes:     []rune("123abc"),
		StrLen:           6,
		IgnoreFields: 	  []string{"Email"},
		CloseChannels:    true,
	})
}
```

## License

[MIT](LICENSE)
