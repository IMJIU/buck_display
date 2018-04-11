package main

import (
	"fmt"
	"time"
	"sync"
	"context"
)

func main() {
	//test_prime()
	//test_wait()
	//test_cancel()
	test_timeout()
}
func test_prime() {
	time.Now()

	origin, wait := make(chan int), make(chan struct{})
	Processor(origin, wait)
	for num := 2; num < 1000; num++ {
		origin <- num
	}
	close(origin)
	<-wait
}

func Processor(seq chan int, wait chan struct{}) {
	go func() {
		prime, ok := <-seq
		if !ok {
			close(wait)
			return
		}
		fmt.Println(prime)
		out := make(chan int)
		Processor(out, wait)
		for num := range seq {
			if num % prime != 0 {
				out <- num
			}
		}
		close(out)
	}()
}
func test_wait() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("before 1")
		time.Sleep(1000 * time.Microsecond)
		fmt.Println("after 1")
		//do...
	}()

	go func() {
		defer wg.Done()
		fmt.Println("before 2")
		time.Sleep(1000 * time.Microsecond)
		fmt.Println("after 2")
	}()
	fmt.Println("before wait")
	wg.Wait()
	fmt.Println("after wait")
}
func test_cancel() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go Proc(ctx)
	go Proc(ctx)
	fmt.Println("before sleep")
	time.Sleep(time.Second)
	fmt.Println("before cancel")
	cancel()
	fmt.Println("after cancel")
}
func Proc(ctx context.Context) {
	fmt.Println("before Proc")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			return
		default:
		//fmt.Println("default")
		}
	}
	time.Sleep(1000 * time.Microsecond)
	fmt.Println("after Proc")
}

func test_timeout() {
	timeout := time.Microsecond * 2000
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	done := make(chan int, 1)
	go func() {
		//RPC(ctx,...)
		fmt.Println("before sleep")
		time.Sleep(time.Microsecond * 1000)
		fmt.Println("after sleep")
		done <- 1
	}()
	select {
	case <-done:
		fmt.Println("done")
	//nice
	case <-ctx.Done():
		fmt.Println("ctx done")
	}
}
