package writer

import (
	"bytes"
	"io"
	"sync"
)

type uniqueWriter struct {
	mu   sync.Mutex
	w    io.Writer
	last []byte
	cur  []byte
	cb   func(line []byte)
}

func UniqueWriter(w io.Writer, cb func([]byte)) *uniqueWriter {
	return &uniqueWriter{
		w:  w,
		cb: cb,
	}
}

func (w *uniqueWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	n, err := w.w.Write(p)
	p = p[:n]
	for {
		i := bytes.IndexByte(p, '\n')
		if i < 0 {
			w.cur = append(w.cur, p...)
			break
		}
		w.cur = append(w.cur, p[:i]...)
		if bytes.Equal(w.cur, w.last) {
			w.cur = w.cur[:0]
		} else {
			w.cb(w.cur)
			w.cur, w.last = w.last[:0], w.cur
		}
		p = p[i+1:]
	}
	return n, err
}
