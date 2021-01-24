package collection

import "container/list"

type lru interface {
	add(key string)
	remove(key string)
}

type emptyLru struct{}

type KeyLru struct {
	limit    int
	evicts   *list.List
	elmments map[string]*list.Element
	onEvict  func(key string)
}

func newKeyLru(limit int, onEvict func(key string)) *KeyLru {
	return &KeyLru{
		limit:    limit,
		evicts:   list.New(),
		elmments: make(map[string]*list.Element),
		onEvict:  onEvict,
	}
}

func (kl *KeyLru) add(key string) {
	if elem, ok := kl.elmments[key]; ok {
		kl.evicts.MoveToFront(elem)
		return
	}

	elem := kl.evicts.PushFront(key)
	kl.elmments[key] = elem

	if kl.evicts.Len() > kl.limit {
		kl.removeOldest()
	}

}

func (kl *KeyLru) remove(key string) {
	if elem, ok := kl.elmments[key]; ok {
		kl.removeElm(elem)
	}
}

func (kl *KeyLru) removeOldest() {
	elem := kl.evicts.Back()
	if elem != nil {
		kl.removeElm(elem)
	}
}

func (kl *KeyLru) removeElm(e *list.Element) {
	kl.evicts.Remove(e)
	key := e.Value.(string)
	delete(kl.elmments, key)
	kl.onEvict(key)
}
