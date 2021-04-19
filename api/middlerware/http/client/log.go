package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"unsafe"

	"github.com/dghubble/sling"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/timex"
)

type LogDoer struct {
	doer sling.Doer
}

func (l LogDoer) Do(req *http.Request) (*http.Response, error) {
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

	logx.WithContext(req.Context()).WithDuration(timex.Since(startTime)).Infof("发送请求[HTTP]: [method: %s, url: %s, params: %s]", req.Method, req.URL.String(), *(*string)(unsafe.Pointer(&reqBody)))

	resp, respErr := l.doer.Do(req)

	var respBody []byte
	if resp != nil && resp.Body != nil {
		if b, err := ioutil.ReadAll(resp.Body); err != nil {
			logx.Error("读取reqBody失败 ", err)
		} else {
			respBody = b
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
		}
	}

	var level string
	if resp != nil {
		var build strings.Builder
		build.WriteString(resp.Status)
		level = build.String()
	} else {
		level = "nil"
	}

	logx.WithContext(req.Context()).
		WithDuration(timex.Since(startTime)).
		Infof("接收响应[HTTP]:[level: %s, method: %s, url: %s, resp: %s]", level, req.Method, req.URL.String(), *(*string)(unsafe.Pointer(&respBody)))
	if respErr != nil {
		return nil, respErr
	}
	return resp, nil
}
