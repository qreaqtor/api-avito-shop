package appserver

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"sync/atomic"

	httpserver "github.com/qreaqtor/api-avito-shop/pkg/httpServer"
)

type AppServer struct {
	started atomic.Bool

	ctx context.Context

	server *httpserver.HTTPServer

	port int

	waitErrChan  chan error
	serveErrChan chan error

	closers []io.Closer
}

func NewAppServer(ctx context.Context, handler http.Handler, env string, port int) *AppServer {
	return &AppServer{
		ctx:          ctx,
		port:         port,
		server:       httpserver.NewHTTPServer(handler, env),
		waitErrChan:  make(chan error),
		serveErrChan: make(chan error, 1),
		closers:      make([]io.Closer, 0),
	}
}

// start listen
func (a *AppServer) Start() error {
	if a.started.Swap(true) {
		return ErrAlreadyStarted
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return err
	}

	go a.serve(l)

	go func(ctx context.Context, l net.Listener) {
		defer a.close(l)

		select {
		case <-ctx.Done():
			return
		case err := <-a.serveErrChan:
			a.waitErrChan <- err
			return
		}
	}(a.ctx, l)

	return nil
}

func (a *AppServer) serve(l net.Listener) {
	defer close(a.serveErrChan)

	err := a.server.Serve(l)
	if err != nil {
		a.serveErrChan <- err
	}
}

func (a *AppServer) close(l net.Listener) {
	err := l.Close()
	if err != nil {
		a.waitErrChan <- err
	}

	err = a.server.Close()
	if err != nil {
		a.waitErrChan <- err
	}

	close(a.waitErrChan)
}

// waiting when all goroutines is done and return serve errors
func (a *AppServer) waitErrors() []error {
	errs := make([]error, 0)

	for err := range a.waitErrChan {
		errs = append(errs, err)
	}

	return errs
}

// waiting when all goroutines is done and return close and serve erros
func (a *AppServer) WaitAndClose() error {
	errs := a.waitErrors()

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	wg.Add(len(a.closers))

	for _, closer := range a.closers {
		go func(closer io.Closer) {
			defer wg.Done()

			err := closer.Close()
			if err != nil {
				mu.Lock()
				errs = append(errs, err)
				mu.Unlock()
			}
		}(closer)
	}

	wg.Wait()

	return errors.Join(errs...)
}
