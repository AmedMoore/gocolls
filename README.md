# gocolls

`gocolls` provides a set of utility functions to work with different types of collections in Go.

## Prerequisites

The following are needed for this project:

```text
Go version 1.21 or higher.
```

## Installing

To install the package:

```terminal
go get github.com/amedmoore/gocolls
```

## Usage

Here's an example of how to use `gocolls`:

```go
package main

import (
	"fmt"
	"github.com/amedmoore/gocolls/slice"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(slice.BoundedSub(s, 2, 12)) // Output: [3,4,5]
	fmt.Println(slice.Rotate(s, 2))         // Output: [3,4,5,1,2]
	fmt.Println(slice.Chunk(s, 3))          // Output: [[1 2 3] [4 5]]
}

```

## Contributing

Contributions to the `gocolls` package are always welcome.

## License

This `gocolls` project is distributed under the MIT License.
See the accompanying [LICENSE](./LICENSE) file for more details.

## Project Status

Please be aware that this project is currently under development, and while the API is stable,
additional features and improvements are forthcoming.
You are welcome to utilize it as needed; however, note that certain functionalities may still be in progress at this
stage of development.
Your understanding and patience are appreciated.
