package sdk

import (
	"net/http"
	"crypto/tls"
	"net/url"
	"io"
	"bytes"
	"io/ioutil"
)

func httpClient() *http.Client  {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	return client
}

func (conn ActConnection) httpGet(url string, data *url.Values) (string, error) {
	return conn.httpCall(url, "GET", data)
}

func (conn ActConnection) httpPost(url string, data *url.Values) (string, error) {
	return conn.httpCall(url, "POST", data)
}

func (conn ActConnection) httpCall(url string, method string, data *url.Values) (string, error) {
	client := httpClient()
	var payload io.Reader
	if data != nil {
		payload = bytes.NewBufferString(data.Encode())
	}

	request, err := http.NewRequest(method, url, payload)
	response, err := client.Do(request)

	if err != nil {
		return "", err
	}

	body, _ := ioutil.ReadAll(response.Body)

	response.Body.Close()
	return string(body), nil
}
