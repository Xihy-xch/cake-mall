package server

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"unsafe"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/timex"
)

func LogServerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var reqBody []byte
		if req.Body != nil {
			if b, err := ioutil.ReadAll(req.Body); err != nil {
				logx.Error("读取reqBody失败 ", err)
			} else {
				reqBody = b
				req.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}
		}
		startTime := timex.Now()
		var buf bytes.Buffer

		logx.WithContext(req.Context()).WithDuration(timex.Since(startTime)).Infof("接收请求[HTTP]: [method: %s, url: %s, params: %s]", req.Method, req.URL.String(), *(*string)(unsafe.Pointer(&reqBody)))
		logWriter := NewLogWriter(w, &buf)
		next(logWriter, req)
		logx.WithContext(req.Context()).WithDuration(timex.Since(startTime)).Infof("发送响应[HTTP]: [method: %s, url: %s, resp: %q]", req.Method, req.URL.String(), buf.String())
	}
}

type LogWriter struct {
	w   http.ResponseWriter
	buf io.ReadWriter
}

func NewLogWriter(w http.ResponseWriter, buf *bytes.Buffer) *LogWriter {
	return &LogWriter{w: w, buf: buf}
}

func (l *LogWriter) Write(p []byte) (int, error) {
	_, err := l.buf.Write(p)
	if err != nil {
		return 0, err
	}

	return l.w.Write(p)
}

func (l *LogWriter) Header() http.Header {
	return l.w.Header()
}

func (l *LogWriter) WriteHeader(statusCode int) {
	l.w.WriteHeader(statusCode)
}
