# Usage

The usage is the same as with the original [`github.com/peske/lsp-srv`](https://github.com/peske/lsp-srv) module. The
only differences are:

- When crating the config, instead of using `lsp_srv.Config` type, you should use `Config` type from this module. Config
  type from this module contains all the fields as the original, with few more fields added.
- When starting the server, instead of using `lsp_srv.Run` function, you should use `Run` function from this module.
  Note that `Run` function here accepts one additional optional argument: `zapLogger`, of type `*zap.Logger`. It is the
  logger that will be used by the server. If not specified, `lsp-srv-ex` module will create a logger.

Here's a simple example:

```go
package main

import (
	"log"
	
	lsp_srv "github.com/peske/lsp-srv-ex"

	"github.com/yourgh/yourmodule/lsp"
)

func main() {
	cfg := &lsp_srv.Config{
		Caching: true, // we want to use caching.
    }

	// Here we assume that your factory function resides in `lsp` package, thus `lsp.NewServer`.
	// Also, we are not specifying logging, so the last argument is `nil`.
	if err := lsp_srv.Run(lsp.NewServer, cfg, nil); err != nil {
		log.Fatal(err)
	}
}
```

## Examples

You can find a few usage examples in https://github.com/peske/lsp-example repository.
