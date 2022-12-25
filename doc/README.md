# Usage

The usage is very similar to the original [`github.com/peske/lsp-srv`](https://github.com/peske/lsp-srv) module, with a
few differences.

## Factory function

In this module the factory function has one additional argument: `helper` of type [`*Helper`](../helper.go#L18), that
can be used for accessing the extended features. So the `Server` type and the factory function should look something
like:

```go
package lsp

import (
  "context"

  lsp_srv_ex "github.com/peske/lsp-srv-ex"
  "github.com/peske/lsp-srv/lsp/protocol"
)

// Server is the type mentioned in the instructions above.
type Server struct {
  client protocol.ClientCloser
  ctx    context.Context
  cancel func()
  helper *lsp_srv_ex.Helper
}

// NewServer is the factory function mentioned in the instructions above.
func NewServer(client protocol.ClientCloser, ctx context.Context, cancel func(), helper *lsp_srv_ex.Helper) *Server {
  return &Server{
    client: client,
    ctx:    ctx,
    cancel: cancel,
	helper: helper,
  }
}
```

## Config

Another difference is in the `Config` type. With this module you should use `lsp_srv_ex.Config`, which has all the
fields as the original `lsp_srv.Config`, with a few additional ones:

- `Caching`, of type `bool`, which determines if the caching feature will be used or not;
- `ZapConfig`, of type `*zap.Config`, which specifies the configuration for `zap.Logger` that will be created and used
  by the server. Content of this field will be ignored if you specify `zapLogger` argument when calling `lsp_srv_ex.Run`
  function.

## Starting the server

When starting the server, instead of using `lsp_srv.Run` function, you should use `Run` function from this module. Note
that `Run` function here accepts one additional argument: `zapLogger`, of type `*zap.Logger`. It is the logger that will
be used by the server. If not specified, this module will create a new logger. The logger will be created by using
`lsp_srv_ex.Config.ZapLoger` configuration if specified, or if it is `nil` by using the default `zap.NewProduction()`.

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
