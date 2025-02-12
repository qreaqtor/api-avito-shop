package appserver

import "io"

type noErrorClose interface {
	Close()
}

type closer struct {
	noErrCloser noErrorClose
}

func newCloser(noErrcloser noErrorClose) *closer {
	return &closer{
		noErrCloser: noErrcloser,
	}
}

func (c *closer) Close() error {
	c.noErrCloser.Close()
	return nil
}

func (a *AppServer) WithClosers(closers ...noErrorClose) *AppServer {
	for _, closer := range closers {
		a.closers = append(a.closers, newCloser(closer))
	}
	return a
}

func (a *AppServer) WithErrorClosers(closers ...io.Closer) *AppServer {
	a.closers = append(a.closers, closers...)
	return a
}
