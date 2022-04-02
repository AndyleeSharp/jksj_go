package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

type MyHandler struct {
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	l, err := net.ListenTCP("tcp", &net.TCPAddr{Port: 8080})
	if err != nil {
		panic(errors.Wrap(err, "listen fail"))
	}
	s := http.Server{Handler: &MyHandler{}}
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		return s.Serve(l)
	})

	g.Go(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			fmt.Println("get signal", sig)
			return s.Shutdown(ctx)
		case <-ctx.Done():
			return ctx.Err()
		}

	})
	fmt.Printf("errgroup exiting: %+v\n", g.Wait())

}
