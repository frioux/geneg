Example package using `go generate`.

Run `go build` and then the built binary.  Note the error.  Now run
`go generate`.  Note the new file in the repo.  Feel free to look at it's
contents.  Run `go build` again and then run the built binary.

By the way, setting a string like this with `go generate` is probably overkill,
since `go build -ldflags "-X 'main.compiledAt=whatever'"` would do that more
simply, but there are times when that isn't an option, or when the structure
being generated is more complex than a simple string.
