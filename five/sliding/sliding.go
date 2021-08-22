package sliding

import (
	"sync"
	"time"
)

type Sliding struct {
	buckets        []*bucket //桶集合
	size           int       // 桶数量
	window         int       // 窗口长度 必须小于等于size
	current        int       // 当前桶
	last           time.Time //
	lastDuration   time.Duration
	bucketDuration time.Duration
	mu             sync.RWMutex
}

type bucket struct {
	val int64
}

func (b *bucket) Reset() {
	b.val = 0
}

type SlidingOption func(opt *Sliding)

func BucketDuration(t time.Duration) SlidingOption {
	return func(opt *Sliding) {
		opt.bucketDuration = t
	}
}

func NewSliding(opts ...SlidingOption) *Sliding {
	now := time.Now()
	sliding := &Sliding{
		bucketDuration: time.Second,
		lastDuration:   time.Duration(0),
		last:           now,
		size:           10,
		window:         5,
	}

	for _, opt := range opts {
		opt(sliding)
	}

	sliding.buckets = make([]*bucket, sliding.size)
	for i := range sliding.buckets {
		sliding.buckets[i] = &bucket{}
	}
	return sliding
}

func (s *Sliding) currentBucket() *bucket {
	old := s.current
	nowDuration := time.Since(s.last)
	d := int((time.Since(s.last) + s.lastDuration%s.bucketDuration) / s.bucketDuration)
	if d > 0 {
		s.last = time.Now()
		s.lastDuration = nowDuration
	}

	s.current = (old + d) % s.size

	if d > s.size {
		d = s.size
	}

	for i := 1; i <= d; i++ {
		s.buckets[(old+i)%s.size].Reset()
	}
	return s.buckets[s.current]
}

func (s *Sliding) Add(val int64) {
	if val == 0 {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	s.currentBucket().val += val
}

func (s *Sliding) Sum() int64 {
	var sum int64

	s.mu.RLock()
	defer s.mu.RUnlock()
	old := s.current
	d := int((time.Since(s.last) + s.lastDuration%s.bucketDuration) / s.bucketDuration)

	n := (old + d) % s.size
	// fmt.Printf("d = %d\n", d)
	start := (n-s.window)%s.size + s.size
	// fmt.Printf("start = %d\n", start)
	for i := 0; i < s.window-d; i++ {
		key := start + i + 1
		// fmt.Printf("n = %d; i = %d; key = %d\n", n, i, key)
		sum += s.buckets[key%s.size].val
	}
	return sum
}
