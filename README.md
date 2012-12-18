goutils
=======

Some utility libraries for Go that aren't special enough to warrant their own repository.

`goutils/daemon/mainloop`
-------------------------

A really simple mainloop that allows you to bind OS signals to handler
functions so you can easily catch any signals that get sent to your daemon.
Here's an example:

```go
import (
    "fmt"
    "github.com/nesv/goutils/daemon/mainloop"
    "syscall"
)

func HupHandler() {
    fmt.Println("caught SIGHUP")
    return
}

func IntHandler() {
    fmt.Println("caugh SIGINT")
    return
}

func main() {
    m := mainloop.New()
	m.Bind(syscall.SIGHUP, HupHandler)
    m.Bind(syscall.SIGINT, IntHandler)
	m.Start()
	return
}
```
