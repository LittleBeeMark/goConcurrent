package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

//
type Conf struct {
	a []int
}

func (c *Conf)test(){}


func BenchmarkAtomic(b *testing.B){
	var c atomic.Value
	c.Store(&Conf{})

	var cfg *Conf

	go func() {
		i:=0
		for {
			i++

			cfg = &Conf{
				a: []int{i,i+1,i+2,i+3,i+4,i+5},
			}

			c.Store(cfg)
		}

	}()

	var wg sync.WaitGroup
	for n:=0;n<4;n++{
		wg.Add(1)
		go func() {
			for n:=0;n<b.N;n++{
				cf :=	c.Load().(*Conf)
				//fmt.Printf("cfg: %+v \n",cf.a)
				cf.test()
			}
			wg.Done()

		}()

	}
	wg.Wait()

}


func BenchmarkMutex(b *testing.B){
	var r sync.RWMutex

	//cfg := &Conf{}
	var cfg *Conf
	go func() {

		i:= 0
		for {
			i++
			r.Lock()
			cfg = &Conf{a:[]int{i,i+1,i+2,i+3,i+4,i+5}}
			r.Unlock()
		}

	}()

	var wg sync.WaitGroup
	for n:=0;n<4;n++{
		wg.Add(1)
		go func() {
			r.RLock()
			for n:=0;n<b.N;n++{
				cfg.test()
				//fmt.Printf("cfg : %+v \n",cfg.a)
			}
			r.RUnlock()
			wg.Done()
		}()
	}

	wg.Wait()

}
