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
