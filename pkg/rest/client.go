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
	protocol  = "https://"
	ingestAPI = "/api/v2/metrics/ingest"
)

type DTClient interface {
	PostMetrics(string)
}

type DTClientImpl struct {
	baseURL string
	token   string
}

func (dtc DTClientImpl) PostMetrics(data string) {
	payload := bytes.NewBuffer([]byte(data))

	req, err := http.NewRequest("POST", dtc.baseURL, payload)
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

func NewDTClient(tenant, token string) DTClient {
	return &DTClientImpl{
		baseURL: protocol + tenant + ingestAPI,
		token:   token,
	}
}
