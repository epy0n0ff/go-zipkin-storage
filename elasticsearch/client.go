package elasticsearch

import (
	"io"
	"net/http"
	"net/url"
)

const userAgent = "go-zipkin-storage/0.1.0"

type Client struct {
	hostURL url.URL
}

func (c Client) Bulk(reader io.Reader) error {
	req, err := http.NewRequest("POST", c.hostURL.RawPath+"/_bulk", reader)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}
