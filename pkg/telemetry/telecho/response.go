package telecho

import (
	"bytes"
	"io"
	"net/http"
)

// responseDumpWriter is a response multiwriter.
type responseDumpWriter struct {
	http.ResponseWriter
	mw  io.Writer
	buf *bytes.Buffer
}

func newResponseDumpWriter(respW http.ResponseWriter) *responseDumpWriter {
	buf := new(bytes.Buffer)

	return &responseDumpWriter{
		ResponseWriter: respW,
		mw:             io.MultiWriter(respW, buf),
		buf:            buf,
	}
}

func (d *responseDumpWriter) Write(b []byte) (int, error) {
	return d.mw.Write(b)
}

func (d *responseDumpWriter) response() *bytes.Buffer {
	return d.buf
}
