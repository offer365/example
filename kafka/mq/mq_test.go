package mq

import "testing"

func TestSyncProducer(t *testing.T) {
	go Consumer("10.0.0.55:9092", "test1")
	SyncProducer("10.0.0.55:9092", "test1", "aa")
}

func TestAsyncProducer(t *testing.T) {
	go Consumer("10.0.0.55:9092", "test2")
	AsyncProducer("10.0.0.55:9092", "test2", "bb")
}
