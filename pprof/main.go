package main

import (
	"fmt"
	_ "net/http/pprof"
)
func main(){
   done := make(chan error,2)
   stop := make(chan struct{})

   go func() {
   	done <- serveDebug(stop)
   }()

	go func() {
		done <- serveAPP(stop)
	}()

   var stopped bool
   for i:=0; i<2;i++{
   	if err := <- done ;err != nil{
   		fmt.Printf("error %v \n",err)
	}

	if ! stopped{
		stopped = true
		close(stop)
	}
   }

}


