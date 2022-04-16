package window

import (
	"sync"
	"time"
)

type Bucket struct {
	Count int64
}
type Window struct {
	buckets        []Bucket
	size           int
	offset         int
	bucketDuration time.Duration
	lastAppendTime time.Time
	mu             sync.RWMutex
}

type options struct {
	size     int
	duration time.Duration
}

type Option func(opt *options)

func Size(size int) Option {
	return func(opt *options) {
		opt.size = size
	}
}
func (b *Bucket) Reset() {
	b.Count = 0
}

func (b *Bucket) Add(count int64) {
	b.Count += count
}

func Duration(d time.Duration) Option {
	return func(opt *options) {
		opt.duration = d
	}
}

func NewWindow(opts ...Option) *Window {
	o := options{
		size:     10,
		duration: time.Second,
	}
	for _, opt := range opts {
		opt(&o)
	}

	w := Window{}
	w.size = o.size
	w.bucketDuration = o.duration
	w.buckets = make([]Bucket, w.size)
	w.lastAppendTime = time.Now()
	return &w
}

func (w *Window) timespan() int {
	v := int(time.Since(w.lastAppendTime) / w.bucketDuration)
	if v > -1 {
		return v
	}
	return w.size
}
func (w *Window) resetBucket(offset int) {
	w.buckets[offset%w.size].Reset()
}
func (w *Window) resetBuckets(offset int, count int) {
	for i := 0; i < count; i++ {
		w.resetBucket(offset + i)
	}
}

func (w *Window) Add(val int) {
	if val < 0 {
		return
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	timespan := w.timespan()
	//oriTimespan := timespan
	if timespan > 0 {
		start := (w.offset + 1) % w.size
		end := (w.offset + timespan) % w.size
		if timespan > w.size {
			timespan = w.size
		}
		w.resetBuckets(start, timespan)
		w.offset = end
		//w.lastAppendTime = time.Now()
		w.lastAppendTime = w.lastAppendTime.Add(time.Duration(timespan * int(w.bucketDuration)))
	}
	w.buckets[w.offset].Add(int64(val))

}

func (w *Window) sumByOffsetAndCount(offset, count int) int64 {
	var val int64
	for i := 0; i < count; i++ {
		val += w.buckets[(offset+i)%w.size].Count
	}
	return val
}

func (w *Window) Sum() int64 {
	w.mu.RLock()
	defer w.mu.RUnlock()
	var val int64
	timespan := w.timespan()
	if count := w.size - timespan; count > 0 {
		offset := (w.offset + timespan + 1) % w.size
		val = w.sumByOffsetAndCount(offset, count)

	}
	return val
}
