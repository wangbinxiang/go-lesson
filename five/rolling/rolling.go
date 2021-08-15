package rolling

import (
	"sync"
	"time"
)

type Rolling struct {
	buckets        []*bucket
	last           time.Time
	current        int
	size           int
	bucketDuration time.Duration
	mu             sync.RWMutex
}

type bucket struct {
	val float64
}

func (b *bucket) Reset() {
	b.val = 0
}

type RollingOption func(opt *Rolling)

func WithBucketDuration(t time.Duration) RollingOption {
	return func(opt *Rolling) {
		opt.bucketDuration = t
	}
}

func NewRolling(opts ...RollingOption) *Rolling {
	rolling := &Rolling{
		size:           10,
		bucketDuration: time.Second,
		last:           time.Now(),
	}

	for _, opt := range opts {
		opt(rolling)
	}

	rolling.buckets = make([]*bucket, rolling.size)
	for i := range rolling.buckets {
		rolling.buckets[i] = &bucket{}
	}
	return rolling
}

func (r *Rolling) currentBucket() *bucket {
	old := r.current
	s := int(time.Since(r.last) / r.bucketDuration)
	if s > 0 {
		r.last = time.Now()
	}

	r.current = (old + s) % r.size

	if s > r.size {
		s = r.size
	}

	for i := 1; i <= s; i++ {
		r.buckets[(old+i)%r.size].Reset()
	}
	return r.buckets[r.current]
}

func (r *Rolling) Add(val float64) {
	if val == 0 {
		return
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.currentBucket().val += val
}

func (r *Rolling) Sum() float64 {
	var sum float64

	r.mu.RLock()
	defer r.mu.RUnlock()
	old := r.current
	s := int(time.Since(r.last) / r.bucketDuration)
	n := (old + s) % r.size

	for i := 0; i < r.size-s; i++ {
		sum += r.buckets[(n+i+1)%r.size].val
	}
	return sum
}
