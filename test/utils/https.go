package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

func GetRequest(u string, headers http.Header, query url.Values, rs any) error {
	req, err := http.NewRequest(http.MethodGet, u+"?"+query.Encode(), nil)
	if err != nil {
		return err
	}
	if headers != nil {
		for key, hs := range headers {
			for _, h := range hs {
				req.Header.Add(key, h)
			}
		}
	}
	var client = &http.Client{
		Timeout: time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bodyBytes, rs)
	return err
}

func PostRequest(u string, headers http.Header, body, rs any) error {
	reqByte, _ := json.Marshal(body)
	req, err := http.NewRequest(http.MethodPost, u, bytes.NewReader(reqByte))
	if err != nil {
		return err
	}
	if headers != nil {
		for key, hs := range headers {
			for _, h := range hs {
				req.Header.Add(key, h)
			}
		}
	}
	var client = &http.Client{
		Timeout: time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bodyBytes, rs)
	return err
}
