package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Client struct {
	Host    string
	Timeout int64
}

func (c *Client) Request(endpoint string) []byte {
	client := http.Client{
		Timeout: time.Duration(time.Duration(c.Timeout).Seconds()),
	}

	resp, err := client.Get(fmt.Sprintf("%s%s", c.Host, endpoint))
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
