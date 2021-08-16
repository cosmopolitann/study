package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func PostJson(apiUrl string, headers map[string][]string, data map[string]interface{}) (res interface{}, err error) {

	cli := http.Client{}

	url, err := url.Parse(apiUrl)
	if err != nil {
		return
	}

	header := http.Header{
		"Content-Type": {"application/json"},
	}
	for key, vals := range headers {
		header[key] = vals
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		err = fmt.Errorf("json.Marshal err: %w", err)
		return
	}

	req := http.Request{
		URL:    url,
		Header: header,
		Method: http.MethodPost,
		Body:   ioutil.NopCloser(bytes.NewBuffer(dataBytes)),
	}

	response, err := cli.Do(&req)
	if err != nil {
		err = fmt.Errorf("cli.Do: %w", err)
		return
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("ioutil.ReadAll err: %w", err)
		return
	}

	var bodyJson map[string]interface{}
	err = json.Unmarshal(bodyBytes, &bodyJson)
	if err != nil {
		err = fmt.Errorf("json.Unmarshal err: %w", err)
		return
	}

	code, exists := bodyJson["code"]
	if !exists {
		err = fmt.Errorf("body no code")
		return
	}

	if code.(float64) != 200 {
		var errMsg string
		_, exists := bodyJson["message"]
		if exists {
			errMsg = bodyJson["message"].(string)
		} else {
			errMsg = "fail"
		}
		err = fmt.Errorf(errMsg)
		return
	}

	res, exists = bodyJson["data"]
	if !exists {
		err = fmt.Errorf("body no data")
	}

	return
}
