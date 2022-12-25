# Usage

The usage is the same as with the original [`github.com/peske/lsp-srv`](https://github.com/peske/lsp-srv) module. The
only differences are:

- When crating the config, instead of using `lsp_srv.Config` type, you should use `Config` type from this module. Config
  type from this module contains all the fields as the original, with few more fields added.
- When starting the server, instead of using `lsp_srv.Run` function, you should use `Run` function from this module.

In short, you can just replace the `import` so that instead:

```go
import (
	lsp_srv "github.com/peske/lsp-srv"
)
```

you have:

```go
import (
	lsp_srv "github.com/peske/lsp-srv-ex"
)
```

Another difference is that with this module `Config` often needs to be created because it defines which features you
want to use.

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
	if err := lsp_srv.Run(lsp.NewServer, cfg); err != nil {
		log.Fatal(err)
	}
}
```

## Examples

You can find a few usage examples in https://github.com/peske/lsp-example repository.
