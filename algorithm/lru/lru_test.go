package lru

import (
	"testing"
	"time"
)

func TestLRU(t *testing.T) {
	lru := new(LRU)
	lru.Init(11)

	lru.Visit("雫")
	lru.Visit("嫁")
	lru.Visit("おれの")
	lru.Visit("は")
	lru.Visit("るる")
	lru.Visit("雫")

	lru.Visit(time.Now().String())

	lru.Print()
}
