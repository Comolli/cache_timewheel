package collection

import "sync"

type SharedCalls interface {
	Do(key string, fn func() (interface{}, error)) (interface{}, error)
	DoEx(key string, fn func() (interface{}, error)) (interface{}, bool, error)
}
type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}
type sharedGroup struct {
	calls map[string]*call
	lock  sync.Mutex
}

func NewSharedCalls() SharedCalls {
	return &sharedGroup{
		calls: make(map[string]*call),
	}
}

func (g *sharedGroup) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	c, done := g.createCall(key)
	if done {
		return c.val, c.err
	}

	g.makeCall(c, key, fn)
	return c.val, c.err
}

func (g *sharedGroup) createCall(key string) (c *call, done bool) {
	g.lock.Lock()
	if c, ok := g.calls[key]; ok {
		g.lock.Unlock()
		c.wg.Wait()
		return c, true
	}

	c = new(call)
	c.wg.Add(1)
	g.calls[key] = c
	g.lock.Unlock()

	return c, false
}
