package main

import (
	"context"
	"fmt"
	"time"
)

type Tracker2 struct {
	ch chan string
	stop chan  struct{}
}

func NewTracker()*Tracker2{
	return &Tracker2{
	ch:	make(chan string, 100),
	}
}

func main(){
	tr := NewTracker()
	go tr.Run()
	_ = tr.Event(context.Background(),"test1")
	_ = tr.Event(context.Background(),"test2")
	_ = tr.Event(context.Background(),"test3")
	ctx ,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	tr.Shutdown(ctx)
}

func (t *Tracker2)Event(ctx context.Context,data string)error{
	select {
	case t.ch <- data:
		return nil
	case <- ctx.Done():
		return ctx.Err()
	}
}

func (t *Tracker2)Run(){
	for data := range t.ch{
		time.Sleep(time.Second)
		fmt.Println(data)
	}

	t.stop <- struct{}{}
}


func (t *Tracker2)Shutdown(ctx context.Context) {
	close(t.ch)
	select {
	case <-t.stop:
	case <-ctx.Done():

	}
}