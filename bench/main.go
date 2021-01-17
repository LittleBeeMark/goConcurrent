package main

import (
	"fmt"
	"sync"
	"time"
)

type Config struct {
	a []int
}

func main() {
	//cfg := &Config{}
	//
	//go func() {
	//	i:= 0
	//	for {
	//
	//		i++
	//		cfg.a = []int{i,i+1,i+2,i+3,i+4,i+5,i+6,i+7,i+8,i+9}
	//	}
	//
	//}()
	//
	//var wg sync.WaitGroup
	//for n:=0;n<20;n++{
	//	wg.Add(1)
	//	go func() {
	//		for n:=0;n<100 ; n++{
	//			fmt.Printf("%v \n",cfg.a)
	//		}
	//		wg.Done()
	//	}()
	//
	//}
	//wg.Wait()

	//t1 := time.NewTicker(time.Second)
	//t2  := time.NewTicker(5*time.Second)
	//
	//for {
	//	select {
	//	case <- t1.C:
	//		DealTask()
	//		fmt.Println("1 秒")
	//	case <- t2.C:
	//		fmt.Println("5 秒")
	//
	//	}
	//}
	done := make(chan bool)
	var mu sync.Mutex

	go func() {
		a := 0
		for {
			select {
			case <-done:
				return

			default:
				mu.Lock()
				a++
				fmt.Println("task 1:", a)
				time.Sleep(100 * time.Millisecond)
				mu.Unlock()
			}
		}

	}()

	a := 0
	for i := 0; i < 20; i++ {
		time.Sleep(time.Millisecond)
		mu.Lock()
		a++
		fmt.Println("task 2:", a)
		mu.Unlock()
	}
	done <- true
}

func DealTask() {
	fmt.Println("task")
	time.Sleep(2 * time.Second)
}
