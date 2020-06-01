package rate_limit

import (
	"fmt"
	"testing"
	"time"
)

func TestNewTokenBucket(t *testing.T) {
	NewTokenBucket("test", 1, time.Second*1)
	for i := 1; i <= 100; i++ {
		time.Sleep(time.Millisecond * 100)
		go doWork(GetBucket("test"), 0)
	}
	select {}
}

func doWork(bucket *TokenBucket, i int) {
	if bucket.GetToken() {
		defer bucket.ReleaseToken()
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("*********** SUCCESS: %v \n", i)
	} else {
		fmt.Printf("*********** FAILED : %v \n", i)
	}
}
