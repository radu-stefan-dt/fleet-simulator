/**
 * Copyright (c) 2021 Radu Stefan
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 **/

package rest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/radu-stefan-dt/fleet-simulator/pkg/util"
)

const (
	protocol        = "https://"
	metricIngestAPI = "/api/v2/metrics/ingest"
	logsIngestAPI   = "/api/v2/logs/ingest"
)

type DTClient interface {
	PostMetrics(string)
	PostLogEvent([]byte)
}

type DTClientImpl struct {
	baseURL string
	token   string
}

func (dtc DTClientImpl) PostMetrics(data string) {
	payload := bytes.NewBuffer([]byte(data))

	req, err := http.NewRequest("POST", dtc.baseURL+metricIngestAPI, payload)
	if err != nil {
		util.PrintError(err)
	}
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Authorization", "Api-Token "+dtc.token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		util.PrintError(err)
	}
	if resp.StatusCode != 202 {
		fmt.Println("Got unexpected response code:", resp.StatusCode, "(", resp.Status, ")")
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			util.PrintError(err)
		}
		fmt.Println(string(body))

		os.Exit(1)
	}
}

func (dtc DTClientImpl) PostLogEvent(content []byte) {
	payload := bytes.NewBuffer([]byte(content))

	req, err := http.NewRequest("POST", dtc.baseURL+logsIngestAPI, payload)
	if err != nil {
		util.PrintError(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Api-Token "+dtc.token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		util.PrintError(err)
	}
	if resp.StatusCode != 204 {
		fmt.Println("Got unexpected response code:", resp.StatusCode, "(", resp.StatusCode, ")")
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			util.PrintError(err)
		}
		fmt.Println(string(body))
	}
}

func NewDTClient(tenant, token string) DTClient {
	return &DTClientImpl{
		baseURL: protocol + tenant,
		token:   token,
	}
}
