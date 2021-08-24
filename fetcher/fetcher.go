package fetcher

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/bontaurean/go-test-proxy/models"
	"github.com/bontaurean/go-test-proxy/storage"
)

var client = &http.Client{Timeout: 10 * time.Second}

func Fetch(preq models.ProxyRequest) (*models.ProxyResponse, error) {
	req, err := http.NewRequest(preq.Method, preq.URL, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range preq.Headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	presp := models.ProxyResponse{
		Status:  resp.StatusCode,
		Headers: flattenHeaders(resp.Header),
		Length:  getContentLen(*resp),
	}

	presp.ID = storage.History.Add(preq, presp)
	return &presp, nil
}

func flattenHeaders(h http.Header) (p models.PlainHeaders) {
	p = make(map[string]string, len(h))

	for k, v := range h {
		p[k] = strings.Join(v, ", ")
	}

	return
}

func getContentLen(resp http.Response) int64 {
	if resp.ContentLength >= 0 {
		return resp.ContentLength
	}

	cl, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		return -1
	}

	return cl
}
