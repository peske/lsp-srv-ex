# What?

This module extends the existing [`github.com/peske/lsp-srv`](https://github.com/peske/lsp-srv) module, by adding some
useful, commonly used features:

- Caching of the editor content in the server.

The usage is explained in [./doc/README.md](./doc/README.md).

# Why?

The current version implements only caching feature. Although it kinda goes beyond the base LSP wireframe scope, and the
calling code can implement it by using the existing LSP methods, as explained in the
[Text Document Synchronization](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_synchronization)
section of the LSP specification, we assume that it is general enough, and used enough so that it makes sense to
implement it in the base module. There's simply no need to make all the module users to reinvent the wheel by
implementing their own caching.

# License?

The same ["BSD-3-Clause license"](./LICENSE) used by the original repository.

# Version?

Current `main` branch is based on the original repository commit
[eec0939](https://github.com/peske/lsp-srv/commit/eec0939e3bba18f1628441f87e03db86547bfe0d) from December 25, 2022.
