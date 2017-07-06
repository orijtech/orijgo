// Copyright 2017 orijtech. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package archomp

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync"
)

const (
	apiURL = "https://archomp.orijtech.com/v1"
)

type Client struct {
	sync.RWMutex

	apiKey string

	Transport http.RoundTripper `json:"-"`
}

type Resource struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

type Request struct {
	Resources []*Resource `json:"files"`
}

var (
	errNilRequest     = errors.New("a non-nil request is needed")
	errEmptyResources = errors.New("\"resources\" must contain at least one resource")
	errBlankResources = errors.New("expecting atleast one resource ie with a non-empty URL")
)

func atLeastOneNonBlankURI(items []*Resource) bool {
	nonBlankResourceCount := 0
	for _, item := range items {
		if item != nil && strings.TrimSpace(item.URL) != "" {
			nonBlankResourceCount += 1
		}
	}
	return nonBlankResourceCount >= 1
}

func (req *Request) Validate() error {
	if req == nil {
		return errNilRequest
	}

	if len(req.Resources) < 1 {
		return errEmptyResources
	}

	if !atLeastOneNonBlankURI(req.Resources) {
		return errBlankResources
	}

	return nil
}

func (c *Client) Compress(req *Request) (io.ReadCloser, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	blob, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hReq, err := http.NewRequest("POST", apiURL, bytes.NewReader(blob))
	if err != nil {
		return nil, err
	}
	if c.apiKey != "" {
		hReq.Header.Set("ARCHOMP-API-KEY", c.apiKey)
	}

	res, err := c.httpClient().Do(hReq)
	if err != nil {
		return nil, err
	}

	if !statusOK(res.StatusCode) {
		return nil, errors.New(res.Status)
	}

	return res.Body, nil
}

func (c *Client) httpClient() *http.Client {
	c.RLock()
	defer c.RUnlock()

	if c.Transport == nil {
		return http.DefaultClient
	}

	client := &http.Client{
		Transport: c.Transport,
	}

	return client
}

func statusOK(code int) bool { return code >= 200 && code <= 299 }
