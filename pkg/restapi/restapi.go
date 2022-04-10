package restapi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// Header ...
type Header struct {
	Key   string
	Value string
}

// Call ...
func Call(method, url string, headers []Header, payload []byte) (res map[string]interface{}, err error) {
	httpBody := bytes.NewBuffer(payload)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	r, _ := http.NewRequest(method, url, httpBody)

	for _, h := range headers {
		r.Header.Add(h.Key, h.Value)
	}

	resp, err := client.Do(r)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		return res, errors.New(url + " " + string(body))
	}

	return res, err
}
