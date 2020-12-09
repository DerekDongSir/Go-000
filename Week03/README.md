#### package errgroup
import "golang.org/x/sync/errgroup"

##### func WithContext
```go
func WithContext(ctx context.Context) (*Group, context.Context)
```
WithContext returns a new Group and an associated Context derived from ctx.
The derived Context is canceled the first time a function passed to Go returns a non-nil error or the first time Wait returns, whichever occurs first.

##### func (*Group) Go
```go
func (g *Group) Go(f func() error)
```
Go calls the given function in a new goroutine.
The first call to return a non-nil error cancels the group; its error will be returned by Wait.

##### func (*Group) Wait
```go
func (g *Group) Wait() error
```
Wait blocks until all function calls from the Go method have returned, then returns the first non-nil error (if any) from them.


#### package signal
import "os/signal"

##### func Ignore
```go
func Ignore(sig ...os.Signal)
```
Ignore causes the provided signals to be ignored. If they are received by the program, nothing will happen. Ignore undoes the effect of any prior calls to Notify for the provided signals. If no signals are provided, all incoming signals will be ignored.

##### func Ignored
```go
func Ignored(sig os.Signal) bool
```
Ignored reports whether sig is currently ignored.

##### func Notify
```go
func Notify(c chan<- os.Signal, sig ...os.Signal)
```
Notify causes package signal to relay incoming signals to c. If no signals are provided, all incoming signals will be relayed to c. Otherwise, just the provided signals will.

```go
// Set up channel on which to send signal notifications.
// We must use a buffered channel or risk missing the signal
// if we're not ready to receive when the signal is sent.
c := make(chan os.Signal, 1)
signal.Notify(c, os.Interrupt)

// Block until a signal is received.
s := <-c
fmt.Println("Got signal:", s)
```
##### func Reset
```go
func Reset(sig ...os.Signal)
```
Reset undoes the effect of any prior calls to Notify for the provided signals. If no signals are provided, all signal handlers will be reset.

##### func Stop
```go
func Stop(c chan<- os.Signal)
```
Stop causes package signal to stop relaying incoming signals to c. It undoes the effect of all prior calls to Notify using c. When Stop returns, it is guaranteed that c will receive no more signals.
