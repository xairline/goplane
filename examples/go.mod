module xairline/goplane/example

go 1.21

toolchain go1.21.5

require github.com/xairline/goplane v1.301.0

// NOTE: next line map to local, you may want to remove it in your actual plugin
replace github.com/xairline/goplane v1.301.0 => ../

require github.com/go-errors/errors v1.0.1 // indirect
