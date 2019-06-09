# envar

[![GoDoc](https://godoc.org/github.com/fabritsius/envar?status.svg)](https://godoc.org/github.com/fabritsius/envar)

This module simplifies use of environment variables in Go programs. The package was inspired by [caarlos0/env](https://github.com/caarlos0/env) and serves a similar purpose.

When working on a small project I decided not to hardcode my sensitive information and to use environment variables instead. The `env` module does exactly what I wanted but it is actually pretty big for what it's doing. Adding additonal 358 lines of code just to read couple variables in a more convenient way seemed wrong.

Module `envar` works only with `strings` but it also only has 36 lines of code. I've never actually needed to export anything other than `strings` so this trade seems fair for my usecases (and maybe for yours as well).

## Example

A very basic example:

1. Export environment variables: `export THING1=living THING2=dying`

2. Run example and see a quote

```go
package main

import (
	"fmt"

	"github.com/fabritsius/envar"
)

type config struct {
	Thing1 string `env:"THING1" default:"coding"`
	Thing2 string `env:"THING2" default:"eating"`
}

func main() {
	cfg := config{}
	// Populate config struct (pass a pointer)
	if err := envar.Fill(&cfg); err != nil {
		panic(err)
	}
	// Print a formatted quote
	fmt.Printf("\"Get busy %s or get busy %s.\" – Stephen King\n", cfg.Thing1, cfg.Thing2)
}
```

Result:

```
"Get busy living or get busy dying." – Stephen King
```

You can visit [this gist](https://gist.github.com/fabritsius/8d7e53a90c01f8c3dddf86a5c5232fa3) to see another example.

## Usage

- pass a pointer of an empty struct to a function `Fill()`
- each field in a struct can have two Tags `env` and `default`
- use `env` to set a name of environment variable
- use `default` to set a default value
- each field must be of type `string`
- each field is considered required (you can set empty default value)
- error is returned when `Fill()` fails to fill a field

Examples of use can be found in [envar_test.go](./envar_test.go) file or on the [GoDoc page](https://godoc.org/github.com/fabritsius/envar).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
