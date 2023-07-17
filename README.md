# Num2Fa - Number to Farsi Words
This is a Golang library that helps you convert numbers to Farsi(Persian) words.

## Usage
To use the library, simply import it into your project and call the Num2fa function, passing in the number you want to convert as an argument. For example:
```go
import "github.com/siyanew/num2fa"

func main() {
    // Prints "سیصد هفتاد و شش"
    fmt.Println(num2fa.Num2fa(376))

    // Prints "منفی یک میلیارد و دویست و سی و چهار میلیون و پانصد و شصت و هفت هزار و هشتصد و نود"
    fmt.Println(num2fa.Num2fa(-1234567890))	
}
```

The function returns a string representing the Farsi words corresponding to the input number.

## Supported Range
This library can convert numbers in the range -2147483648 to 2147483647. If the input number is outside this range, the function will return an empty string.

## License
This library is released under the MIT License. Please see the LICENSE file for more details.