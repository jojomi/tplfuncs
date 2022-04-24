[![Go Report Card](https://goreportcard.com/badge/github.com/jojomi/tplfuncs)](https://goreportcard.com/report/github.com/jojomi/tplfuncs)

# tplfuncs
Go module with some useful template.FuncMap goodies

[Documentation](https://pkg.go.dev/github.com/jojomi/tplfuncs)

## Areas of concern

* **[Spacing](spacing.go)** (`newline` and `space`)
* **[Lines](lines.go)** (line-wise processing of input string using `trim`, `head`, `tail`, and more)
* **[Filesystem](fs.go)** (`include` and `glob`)
* **[Exec](exec.go)** (command execution and output capturing)

## Who uses it?

[io](https://github.com/jojomi/io)