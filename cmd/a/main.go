package main

import (
	"context"
	"fmt"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

var ch = make(chan int, 100)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go tik(ctx, wg)
	go readChan(ctx, wg)

	<-ctx.Done()
	wg.Wait()

	fmt.Println("Done...")
}

func tik(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("tik stop")
			return
		case <-ticker.C:
			ch <- 1
			fmt.Println("tik")
		}
	}
}

func readChan(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	data := make([]byte, 0, 10)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("read chan stop")
			return
		case d := <-ch:
			data = strconv.AppendInt(data, int64(d), 10)
			if len(data) > 4 {
				fmt.Println("data =", data)
				data = nil
			}
		}
	}
}
