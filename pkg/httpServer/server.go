package httpserver

import (
	"log/slog"
	"net"
	"net/http"

	comlog "github.com/qreaqtor/api-avito-shop/pkg/logging"
)

type HTTPServer struct {
	server *http.Server
}

// return http server with added recovery middleware
// used default slog.Logger{}
func NewHTTPServer(handler http.Handler, env string) *HTTPServer {
	comlog.SetLogger(env)

	return &HTTPServer{
		server: &http.Server{
			Handler: panicMiddleware(handler),
		},
	}
}

func (h *HTTPServer) Serve(l net.Listener) error {
	slog.Info("Start http server at " + l.Addr().String())
	return h.server.Serve(l)
}

func (h *HTTPServer) Close() error {
	slog.Info("Stop http server")
	return h.server.Close()
}
