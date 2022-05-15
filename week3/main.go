package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

func myfunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func startServer(srv *http.Server) error {
	http.HandleFunc("/1", myfunc)
	fmt.Println("http /1 server start port 8083")
	error := srv.ListenAndServe()
	return error
}

// 基于 errgroup 实现一个 http server 的启动和关闭
// 以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出
func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	group, errCtx := errgroup.WithContext(ctx)

	srv := &http.Server{Addr: ":8083"}

	group.Go(func() error {
		return startServer(srv)
	})

	group.Go(func() error {
		<-errCtx.Done()
		fmt.Println("http server stop")
		return srv.Shutdown(errCtx)
	})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit)

	group.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				return errCtx.Err()
			case <-quit:
				cancel()
			}
		}
		return nil
	})

	fmt.Printf("errgroup exiting: %+v\n", group.Wait())
}
