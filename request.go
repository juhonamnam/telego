package telego

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

const baseUrl = "https://api.telegram.org/bot"

type baseResponse struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description"`
}

func (telego *telegoStruct) Request(endpoint string, data any) (*[]byte, error) {
	pbytes, err := json.Marshal(data)
	if err != nil {
		telego.logger.Error(endpoint, "Request:", data, "Error:", err.Error())
		return nil, err
	}

	buff := bytes.NewBuffer(pbytes)
	resp, err := telego.httpClient.Post(baseUrl+telego.apiKey+"/"+endpoint, "application/json", buff)

	if err != nil {
		telego.logger.Error(endpoint, "Request:", string(pbytes), "Error:", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		telego.logger.Error(endpoint, "Request:", string(pbytes), "Error:", err.Error())
		return nil, err
	}

	telego.logger.Info(endpoint, "Request:", string(pbytes), "Response:", string(body))
	return &body, nil
}
