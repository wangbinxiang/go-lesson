package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func main() {

	srv := &http.Server{Addr: ":8000"}

	// 注册http服务 / /stop
	// server监听

	// 启动

	g, ctx := errgroup.WithContext(context.Background())

	// ctx := context.Background()

	// ctx, cancel := context.WithCancel(ctx)

	// g, errCtx := errgroup.WithContext(ctx)

	// 一个http服务
	serverChan := make(chan struct{})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})
	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "server shut down")
		serverChan <- struct{}{}
	})

	srv.RegisterOnShutdown(func() {
		fmt.Println("server on shutdown")
	})

	g.Go(func() error {
		return srv.ListenAndServe()
	})

	g.Go(func() error {
		select {
		case <-serverChan:
			fmt.Println("server stop with serverChan")
		case <-ctx.Done():
			fmt.Println("server stop with ctx.Done")
		}

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		fmt.Println("http server shutdown")
		return srv.Shutdown(timeoutCtx)
	})

	g.Go(func() error {
		c := make(chan os.Signal, 4)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
		select {
		case sig := <-c:
			fmt.Printf("signal stop with signal: %v\n", sig)
			return errors.Errorf("get os signal: %v", sig)
		case <-ctx.Done():
			fmt.Println("signal stop with ctx.Done")
			return ctx.Err()
		}
	})

	g.Wait()
	fmt.Printf("errgroup done: %v\n", g.Wait())
}
