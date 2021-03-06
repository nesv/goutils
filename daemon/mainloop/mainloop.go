/*
Creates a mainloop, that handles signals received from the operating system.
*/
package mainloop

import (
	"errors"
	"os"
	"os/signal"
)

var (
	SignalAlreadyBoundError = errors.New("Signal is already bound to a function")
)

/*
This struct represents the mainloop for catching operating system signals in
your application.
*/
type Mainloop struct {
	sigchan  chan os.Signal
	termchan chan int
	Bindings map[os.Signal]func()
}

/*
Initializes and returns a pointer to a new Mainloop.
*/
func New() *Mainloop {
	m := Mainloop{
		sigchan:  make(chan os.Signal),
		termchan: make(chan int),
		Bindings: make(map[os.Signal]func())}
	return &m
}

/*
Bind an operating system signal to a handler function, prior to calling
Mainloop.Start().

You cannot bind multiple functions to the same signal, and any attempt to
do so will raise an error.
*/
func (m *Mainloop) Bind(sig os.Signal, f func()) (err error) {
	for s, _ := range m.Bindings {
		if sig == s {
			err = SignalAlreadyBoundError
			return
		}
	}
	m.Bindings[sig] = f
	return nil
}

/*
Unbind the currently-bound function from the specified operating system signal.

If the signal is not bound to a handler function, then this method is
effectively a no-op.
*/
func (m *Mainloop) Unbind(sig os.Signal) {
	delete(m.Bindings, sig)
	return
}

/*
Start the mainloop.

This method will block its current thread. The best spot for calling this
method is right near the bottom of your application's main() function.
*/
func (m *Mainloop) Start() {
	sigs := make([]os.Signal, len(m.Bindings))
	for s, _ := range m.Bindings {
		sigs = append(sigs, s)
	}
	signal.Notify(m.sigchan, sigs...)
	for {
		select {
		case sig := <-m.sigchan:
			m.Bindings[sig]()
		case _ = <-m.termchan:
			break
		}
	}
	return
}

/*
Stops the mainloop.
*/
func (m *Mainloop) Stop() {
	go func() { m.termchan <- 1 }()
	return
}
