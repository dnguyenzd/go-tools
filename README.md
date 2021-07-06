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
