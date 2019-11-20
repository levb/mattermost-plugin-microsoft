// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package msgraph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	msgraph "github.com/yaegashi/msgraph.go/v1.0"
)

func (c *client) Call(method, path string, in, out interface{}) (responseData []byte, err error) {
	errContext := fmt.Sprintf("msgraph: Call failed: method:%s, path:%s", method, path)
	baseURL, err := url.Parse(c.rbuilder.URL())
	if err != nil {
		return nil, errors.WithMessage(err, errContext)
	}
	if len(path) > 0 && path[0] != '/' {
		path = "/" + path
	}
	path = baseURL.String() + path

	var inBody io.Reader
	if in != nil {
		buf := &bytes.Buffer{}
		err = json.NewEncoder(buf).Encode(in)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, path, inBody)
	if err != nil {
		return nil, err
	}
	if inBody != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	if c.ctx != nil {
		req = req.WithContext(c.ctx)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()

	responseData, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated:
		if out != nil {
			err = json.Unmarshal(responseData, out)
			if err != nil {
				return responseData, err
			}
		}
		return responseData, nil

	case http.StatusNoContent:
		return nil, nil
	}

	errResp := msgraph.ErrorResponse{Response: resp}
	err = json.Unmarshal(responseData, &errResp)
	if err != nil {
		return responseData, errors.WithMessagef(err, "status: %s", resp.Status)
	}
	if err != nil {
		return responseData, err
	}
	return responseData, &errResp
}
