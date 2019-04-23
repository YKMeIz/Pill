package pill

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// fetch send "GET" request with given url and custom headers, and return the response body.
func fetch(u string, headers []header) ([]byte, error) {
	c := http.DefaultClient

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	for _, v := range headers {
		req.Header.Set(v.key, v.val)
	}

	resp, err := c.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("expect status ok (200), get: " + resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}

// headStatusOk send "HEAD" request with given url and custom headers in order to check if url is existed (200).
// In this package, it is utilized to check if assumed illustration url has content without transmitting response body.
func headStatusOk(u string, headers []header) bool {
	c := http.DefaultClient

	req, err := http.NewRequest("HEAD", u, nil)
	if err != nil {
		return false
	}

	for _, v := range headers {
		req.Header.Set(v.key, v.val)
	}

	resp, err := c.Do(req)

	if err != nil {
		return false
	}

	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}
