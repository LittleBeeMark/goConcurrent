package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"sync"
	"time"
)

type Tracker struct {
   wg   sync.WaitGroup
}

type App struct {
	track Tracker
}

func (a *App)Handle(w http.ResponseWriter,r *http.Request){
	w.WriteHeader(http.StatusCreated)

	go a.track.Event("this event")
}

func (t *Tracker)Event(data string){
	t.wg.Add(1)

	go func() {
		defer t.wg.Done()

		time.Sleep(time.Millisecond)
		log.Println(data)

	}()

}

func (t *Tracker)Shutdown(ctx context.Context)error{
	ch := make(chan struct{})


	go func() {
		t.wg.Wait()
		close(ch)
	}()

	select {
	case <-ch:
		return nil
	case <-ctx.Done():
		return errors.New("timeout")
		
	 }
}

func main(){
	var a App

	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	err := a.track.Shutdown(ctx)
	if err != nil{
		log.Println("shutdown err :",err)
	}
}