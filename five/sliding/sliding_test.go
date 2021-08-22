package sliding

import (
	"testing"
	"time"
	// "go-lesson/five/sliding"
)

func TestSliding(t *testing.T) {
	s := NewSliding(BucketDuration(time.Second))
	for i := 0; i < 4; i++ {
		s.Add(1)
		time.Sleep(time.Second)
	}
	time.Sleep(time.Second)
	got := s.Sum()
	if got != 3 {
		var buckets []int
		for _, b := range s.buckets {
			buckets = append(buckets, int(b.val))
		}
		t.Errorf("s.Sum() = %d, want %d, %v\n", got, 3, buckets)
	}
}

func TestSlidingEmpty(t *testing.T) {
	s := NewSliding(BucketDuration(time.Millisecond))
	for i := 0; i < 4; i++ {
		s.Add(1)
		time.Sleep(time.Millisecond)
	}
	time.Sleep(10 * time.Millisecond)
	got := s.Sum()
	if got != 0 {
		var buckets []int
		for _, b := range s.buckets {
			buckets = append(buckets, int(b.val))
		}
		t.Errorf("s.Sum() = %d, want %d, %v\n", got, 0, buckets)
	}
}

func TestSlidingOne(t *testing.T) {
	s := NewSliding(BucketDuration(time.Second))
	s.Add(10)
	got := s.Sum()
	if got != 10 {
		var buckets []int
		for _, b := range s.buckets {
			buckets = append(buckets, int(b.val))
		}
		t.Errorf("s.Sum() = %d, want %d, %v\n", got, 10, buckets)
	}
}
