package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func main() {

	g, ctx := errgroup.WithContext(context.Background())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})
	serverChan := make(chan struct{})
	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		serverChan <- struct{}{}
		fmt.Fprint(w, "stop server ")
	})

	srv := &http.Server{Addr: ":8000"}

	g.Go(func() error {
		return srv.ListenAndServe()
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
			log.Println("errgroup shutdown")
		case <-serverChan:
			log.Println("server will stop")
		}
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		return srv.Shutdown(timeoutCtx)
	})

	// 一个sign信号
	g.Go(func() error {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

		select {
		case sig := <-c:
			return errors.Errorf("get os signal: %v", sig)
		case <-ctx.Done():
			return ctx.Err()
		}
	})

	fmt.Printf("errgroup done: %v\n", g.Wait())
}

// func operation1(ctx context.Context) error {
// 	// 让我们假设这个操作会因为某种原因失败
// 	// 我们使用time.Sleep来模拟一个资源密集型操作
// 	time.Sleep(100 * time.Millisecond)
// 	return errors.New("failed")
// }

// func operation2(ctx context.Context) {
// 	// 我们使用在前面HTTP服务器例子里使用过的类型模式
// 	select {
// 	case <-time.After(500 * time.Millisecond):
// 		fmt.Println("done")
// 	case <-ctx.Done():
// 		fmt.Println("halted operation2")
// 	}
// }

// func main() {
// 	// 新建一个上下文
// 	ctx := context.Background()
// 	// 在初始上下文的基础上创建一个有取消功能的上下文
// 	ctx, cancel := context.WithCancel(ctx)
// 	// 在不同的goroutine中运行operation2

// 	fmt.Println("start")
// 	go func() {
// 		operation2(ctx)
// 	}()

// 	err := operation1(ctx)
// 	// 如果这个操作返回错误，取消所有使用相同上下文的操作
// 	if err != nil {
// 		cancel()
// 	}
// 	time.Sleep(1000 * time.Millisecond)
// 	fmt.Println("end")
// }
