![Test](https://github.com/siyanew/num2fa/actions/workflows/test.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/siyanew/num2fa.svg)](https://pkg.go.dev/github.com/siyanew/num2fa)

# Num2Fa - Number to Farsi Words

This is a Golang library that helps you convert numbers to Farsi(Persian) words.

## Installation
You can install this library using the `go get` command:

```shell
go get github.com/siyanew/num2fa
```

## Usage
To use the library, simply import it into your project and call the Num2fa function, passing in the number you want to convert as an argument. For example:
```go
import "github.com/siyanew/num2fa"

func main() {
// Prints "سیصد هفتاد و شش"
fmt.Println(num2fa.Convert(376))

// Prints "منفی یک میلیارد و دویست و سی و چهار میلیون و پانصد و شصت و هفت هزار و هشتصد و نود"
fmt.Println(num2fa.Convert(-1234567890))
}
```

The function returns a string representing the Farsi words corresponding to the input number.

## License
This library is released under the MIT License. Please see the LICENSE file for more details.