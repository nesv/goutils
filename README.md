goutils
=======

Some utility libraries for Go that aren't special enough to warrant their own repository.

`goutils/daemon/mainloop`
-------------------------

A really simple mainloop that listens for SIGTERM, SIGKILL and SIGHUP signals
from the operating system.

To use this simple main-loop, after `go get`ting the package, just put
`mainloop.Start()` at the end of your `main()` function.

```go
import "github.com/nesv/goutils/daemon/mainloop"

func main() {
	...
	mainloop.Start()
	return
}
```
