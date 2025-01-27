package closer

import (
	"log"
	"os"
	"os/signal"
	"sync"
)

type Closer struct {
	mu        sync.Mutex
	once      sync.Once
	done      chan struct{}
	functions []func() error
}

var globalCloser = New()

func Add(f ...func() error) {
	globalCloser.Add(f...)
}

func CloseAll() {
	globalCloser.CloseAll()
}

func Wait() {
	globalCloser.Wait()
}

// New returns new Closer, if []os.Signal is specified Closer will automatically call CloseAll when one of signals is received from OS
func New(sig ...os.Signal) *Closer {
	c := &Closer{
		done: make(chan struct{}),
	}
	if len(sig) > 0 {
		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, sig...)
			<-ch
			signal.Stop(ch)
			c.CloseAll()
		}()
	}

	return c
}

func (c *Closer) Add(f ...func() error) {
	c.mu.Lock()
	c.functions = append(c.functions, f...)
	c.mu.Unlock()
}

// CloseAll calls all closer functions
func (c *Closer) CloseAll() {
	c.once.Do(func() {
		defer close(c.done)

		c.mu.Lock()
		funcs := c.functions
		c.functions = nil
		c.mu.Unlock()

		errors := make(chan error, len(funcs))
		for _, f := range funcs {
			go func(f func() error) {
				errors <- f()
			}(f)
		}
		for i := 0; i < cap(errors); i++ {
			if err := <-errors; err != nil {
				log.Println("error returned from Closer")
			}
		}
	})
}

// Wait blocks until all closer functions are done
func (c *Closer) Wait() {
	<-c.done
}
