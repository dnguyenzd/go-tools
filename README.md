# go tools setup

## Common way
A common and recommended way to track and lock tool dependencies in Go projects is to add a file `tools.go` with build constraint `// +build tools` as mentioned here: [https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module](https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module)

But doing it this way in main module also means tool dependencies are mixed together with main module dependencies. If main module is vendored, tool dependencies are also vendored together.

## This setup
A tweak to above setup is to create a separate Go module for `tools.go` (.i.e. a separate `go.mod` in `tools` sub-directory)

### Advantages
- Main module and tools module are 2 separate modules, so dependencies in `go.mod` are separated from each other
- Main module and tools module can use mix of vendoring and GOPATH. E.g. for tool module, since it's used less frequent, people might use GOPATH instead of vendoring to avoid bloating repository size. The setup in this repo uses vendoring for main module and GOPATH for tool module

### Disadvantages
- All go commands related to tooling must be run in `tools` sub-directory, .e.g. `cd tools && go install github.com/golang/mock/mockgen`

### What goes where?
- Tools/libraries that are required for compilation of main module go to main module `go.mod`
- Other tools (.e.g. `golangci-lint`) go to tools module `go.mod`
- Special case: `github.com/golang/mock` module includes `github.com/golang/mock/gomock` and `github.com/golang/mock/mockgen` packages.
  - `gomock` is required for compilation of main module (since it's imported and used in main module tests), while `mockgen` is only used as a code generation tool and is not needed when compiling main module
  - in this case, `github.com/golang/mock` can be included as dependency to both main module and tool module. Their version will be tracked separately from each other. `github.com/golang/mock/gomock` will be imported and vendored with main module. `github.com/golang/mock/mockgen` is only in tool module, not vendored together with main module
