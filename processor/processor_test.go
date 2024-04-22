package processor

import (
	"context"
	"math/rand"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestIntervalFrom(t *testing.T) {}

func TestCreateStatProcessor(t *testing.T) {}

func BenchmarkProcessMessages(b *testing.B) {
	b.StopTimer()

	ctx := context.Background()

	var wg sync.WaitGroup
	wg.Add(1)

	receiveMessages := make(chan UserMsg, 100)
	messagesToSend := make(chan StatMsg, 100)

	go func() {
		ts := uint64(time.Now().Unix())
		for i := 0; i < b.N; i++ {
			receiveMessages <- UserMsg{
				Uid: "user" + strconv.Itoa(rand.Intn(100)),
				Ts:  ts,
			}
			ts += uint64(rand.Intn(100))
		}
		close(receiveMessages)
	}()

	b.StartTimer()

	go ProcessMessages(ctx, &wg, receiveMessages, messagesToSend)

	for range messagesToSend {
	}
	wg.Wait()
}
