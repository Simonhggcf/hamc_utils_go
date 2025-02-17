# HMAC-Golang

HMAC-Golang is a Go package for generating and verifying HMAC (Hash-based Message Authentication Code) using SHA-256.

## Installation

To use this package, you can clone the repository or use it as a module in your Go project.

```bash
go get github.com/Simonhggcf/HMAC-Golang/hmac_util
```

## Usage

### Generating HMAC

```go
package main

import (
	"fmt"
	"github.com/Simonhggcf/HMAC-Golang/hmac_util"
)

func main() {
	message := []byte("Hello, World!")
	key := []byte("supersecretkey")
	hmacValue := hmacutil.GenerateHMAC(message, key)
	fmt.Println("Generated HMAC:", hmacValue)
}
```

### Verifying HMAC

```go
package main

import (
	"fmt"
	"github.com/Simonhggcf/HMAC-Golang/hmac_util"
)

func main() {
	message := []byte("Hello, World!")
	key := []byte("supersecretkey")
	hmacValue := hmacutil.GenerateHMAC(message, key)

	isValid, err := hmacutil.VerifyHMAC(message, key, hmacValue)
	if isValid {
		fmt.Println("HMAC verification succeeded")
	} else {
		fmt.Println("HMAC verification failed:", err)
	}
}
```

## Running Tests

To run the tests for this package, use the following command:

```bash
go test ./hmac_util
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License.
