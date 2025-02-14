package appserver

import "io"

func (a *AppServer) WithClosers(closers ...io.Closer) *AppServer {
	a.closers = append(a.closers, closers...)
	return a
}
