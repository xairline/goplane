module xairline/goplane/example

go 1.17

require github.com/xairline/goplane v3.0.1

// NOTE: next line map to local, you may want to remove it in your actual plugin
replace github.com/xairline/goplane v3.0.1 => ../

require github.com/go-errors/errors v1.0.1 // indirect
