# faker
[![GoDoc](https://godoc.org/github.com/koaim/faker?status.svg)](https://godoc.org/github.com/koaim/faker) [![license](http://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/koaim/faker/main/LICENSE)

A lightweight random value generator for tests, written in Go.

## Installation

```bash
go get github.com/koaim/faker
```

## Why?

I was using [gofakeit](https://github.com/brianvoe/gofakeit) but only ever needed `gofakeit.Struct()` — filling a struct with random data for tests. Everything else was unused. So I wrote a small, focused library that does just that, without pulling in the rest of a big dependency.

## Features

- Convenient API with generics support
- Supports int/uint, float/complex, string, bool, struct, slice, array, map, chan and pointer types
- Zero dependencies

## Usage

```go
package main

import (
	"fmt"

	"github.com/koaim/faker"
)

type User struct {
	ID    int
	Name  string
	Email string
	Admin bool
}

func main() {
	// fill struct with random data
	u := faker.Make[User]()

	// fill struct with random data and override specific fields
	u = faker.MakeAndOverride(func(v *User) {
		v.Email = "test@example.com"
	})

	// fill struct and ignore some fields
	u = faker.Make[User](faker.WithIgnoreFields("Admin"))
}
```

## License

[MIT](LICENSE)
