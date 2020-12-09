//  基于 errgroup 实现一个 http server 的启动和关闭
// 以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	var ctx context.Context
	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World"))
	})

	g.Go(func() error {
		s := http.Server{Addr: ":8080"}

		go func() {
			<-ctx.Done()
			s.Shutdown(context.TODO())
		}()
		return s.ListenAndServe()
	})

	g.Go(func() error {
		s := http.Server{Addr: ":8081"}

		go func() {
			<-ctx.Done()
			s.Shutdown(context.TODO())
		}()
		return s.ListenAndServe()
	})

	g.Go(func() error {
		sig := make(chan os.Signal)
		signal.Notify(sig, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
		log.Println("exit signal: ", <-sig)
		cancel()
		return nil
	})
	log.Println(g.Wait())
}
