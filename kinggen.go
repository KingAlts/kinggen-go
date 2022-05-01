package kinggen

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func NewKingGen(apiKey string) (*KingGen, error) {
	if len(strings.TrimSpace(apiKey)) == 0 {
		return nil, errors.New("invalid api key")
	}
	return &KingGen{apiKey: apiKey}, nil
}

func (kg *KingGen) GetProfile() (*Profile, error) {
	return makeRequest[Profile](kg, profileEndpoint)
}

func (kg *KingGen) GetAlt() (*Alt, error) {
	return makeRequest[Alt](kg, altEndpoint)
}

func makeRequest[T any](kg *KingGen, endpoint Endpoint) (*T, error) {
	res, err := http.Get(endpoint.build(kg.apiKey))
	if err != nil {
		return new(T), err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return new(T), err
	}
	if res.StatusCode >= 300 {
		var errRes map[string]interface{}
		err = json.Unmarshal(body, &errRes)
		if err != nil {
			return new(T), err
		}
		return new(T), errors.New(strings.ToLower(errRes["message"].(string)))
	}
	var resDeserialized T
	err = json.Unmarshal(body, &resDeserialized)
	return &resDeserialized, err
}
