package main

import (
	"context"
	"net/http"
)


// debug 服务
func serveDebug(stop <-chan struct{})error{
	return serve(":6060",nil,stop)
}

// app 服务
func serveAPP(stop <-chan struct{})error{
	return serve(":8080",nil,stop)
}

// 启动一个服务
func serve(addr string,handler http.Handler,stop <-chan struct{})error{
	s := http.Server{
		Addr: addr,
		Handler: handler,

	}

	go func() {
		<- stop
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}