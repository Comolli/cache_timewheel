package collection

import (
	"cache_timewheel/config"
	"sync/atomic"
	"time"
)

type Stat struct {
	name     string
	hit      uint64
	miss     uint64
	Callback func() int
}

func NewState(name string, callBack func() int) *Stat {
	st := &Stat{
		name:     name,
		Callback: callBack,
	}
	go st.StartLoop()

}

func (state *Stat) IncreamentHit() {
	atomic.AddUint64(&state.hit, 1)
}

func (state *Stat) IncrementMiss() {
	atomic.AddUint64(&state.miss, 1)
}

func (state *Stat) StartLoop() {
	ticker := time.NewTicker(config.StatInterval)
	defer ticker.Stop()
	for range ticker.C {
		hit := atomic.SwapUint64(&state.hit, 0)
		miss := atomic.SwapUint64(&state.miss, 0)
		total := hit + miss
		if total == 0 {
			continue
		}
	}
}
