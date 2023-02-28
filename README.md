# Matrix

[![Go Reference](https://pkg.go.dev/badge/github.com/prophittcorey/matrix.svg)](https://pkg.go.dev/github.com/prophittcorey/matrix)

A simple Matrix client package written in Go.

## Package Usage

```go
package main

import (
  "log"

  "github.com/prophittcorey/matrix"
)

func main() {
  message := `Hi, this is a <a href="https://github.com/prophittcorey">html friendly</a> message.`

  client := matrix.New("USERNAME", "PASSWORD")

  if err := client.Send("ROOM ID", message); err != nil {
    log.Fatal(err)
  }
}
```

## Tool Usage

```bash
# Install the latest tool.
go install github.com/prophittcorey/matrix/cmd/matrix@latest

# Fire a message away.
matrix --username "SOMEONE" --password "PASSWORD" --roomid "ROOM ID" --message "Your message goes here."
```

## License

The source code for this repository is licensed under the MIT license, which you can
find in the [LICENSE](LICENSE.md) file.
