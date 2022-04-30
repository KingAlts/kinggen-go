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
	res, err := makeRequest[Profile](kg, ProfileEndpoint)
	return &res, err
}

func (kg *KingGen) GetAlt() (*Alt, error) {
	res, err := makeRequest[Alt](kg, AltEndpoint)
	return &res, err
}

func makeRequest[T any](kg *KingGen, endpoint Endpoint) (T, error) {
	res, err := http.Get(endpoint.Build(kg.apiKey))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 300 {
		var errRes map[string]interface{}
		err = json.Unmarshal(body, &errRes)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(strings.ToLower(errRes["message"].(string)))
	}
	var resDeserialized T
	err = json.Unmarshal(body, &resDeserialized)
	return resDeserialized, err
}
