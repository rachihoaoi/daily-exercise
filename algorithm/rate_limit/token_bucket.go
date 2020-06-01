package rate_limit

import (
	"sync"
	"time"
)

const (
	DEFAULT_BUCKET_CAP   = 20
	DEFAULT_ADD_INTERVAL = time.Millisecond * 20
)

var bucket sync.Map

type TokenBucket struct {
	cap      int // 桶的容量
	ch       chan bool
	timeRate *time.Ticker // 放入令牌的速率
	mx       sync.Mutex   // 互斥锁
}

func (b *TokenBucket) InitCap() {
	for i := len(b.ch); i < b.cap; i++ {
		b.ch <- true
	}
	return
}

func (b *TokenBucket) AutoAddToken() {
	for {
		select {
		case <-b.timeRate.C:
			if len(b.ch) < b.cap {
				b.ch <- true
			}
		}
	}
}

func (b *TokenBucket) GetToken() bool {
	select {
	case <-b.ch:
		return true
	default:
		return false
	}
}

func (b *TokenBucket) ReleaseToken() {
	b.mx.Lock()
	if len(b.ch) < b.cap {
		b.ch <- true
	}
	b.mx.Unlock()
}

func NewTokenBucket(bucketName string, cap int, internal time.Duration) {
	if _, ok := bucket.Load(bucketName); ok {
		return
	}
	_bucket := &TokenBucket{
		cap:      cap,
		ch:       make(chan bool, cap),
		timeRate: time.NewTicker(internal),
	}
	bucket.Store(bucketName, _bucket)
	b := GetBucket(bucketName)
	b.InitCap()
	go b.AutoAddToken()
}

func GetBucket(bucketName string) *TokenBucket {
	var b interface{}
	var ok bool
	if b, ok = bucket.Load(bucketName); ok || b == nil {
		NewTokenBucket(bucketName, DEFAULT_BUCKET_CAP, DEFAULT_ADD_INTERVAL)
		b, _ = bucket.Load(bucketName)
		return b.(*TokenBucket)
	}
	return b.(*TokenBucket)
}
