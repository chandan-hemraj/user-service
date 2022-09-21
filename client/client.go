package client

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/go-chassis/go-chassis/client/rest"
	"github.com/go-chassis/go-chassis/v2/core"
	"github.com/go-chassis/go-chassis/v2/pkg/util/httputil"
	"github.com/go-chassis/openlog"
)

// /**
func MakeRequest(Url, method string, payload []byte, headers map[string]string) (map[string]interface{}, string, error) {
	openlog.Info("Making request to : " + Url)
	req, err := rest.NewRequest(method, Url, payload)
	if err != nil {
		openlog.Error("new request failed. " + err.Error())
		return nil, "", errors.New("internal server error")
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := core.NewRestInvoker().ContextDo(context.TODO(), req)
	if err != nil {
		openlog.Error("do request failed. : " + err.Error())
		return nil, "", errors.New("internal server error")
	}
	var res = make(map[string]interface{})
	_ = json.Unmarshal(httputil.ReadBody(resp), &res)

	return res, "", nil
}
