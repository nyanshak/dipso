package dipso

import (
	"net/http"
	"errors"
	"io/ioutil"
	"encoding/json"
	"strings"
	"bytes"
)

const (
	BaseUrl		= "http://services.wine.com/api/beta2/service.svc/json"
	CatalogUrl	= BaseUrl + "/catalog"
	CategoryUrl	= BaseUrl + "/categoryMap"
)

type WineApi struct {
	apiKey string
}

func NewWineApi(apiKey string) *WineApi {
	api := &WineApi{apiKey: apiKey};
	return api
}

func (self WineApi) SetApiKey(wineApiKey string) {
	self.apiKey = wineApiKey
}

func (self WineApi) Search(params string) (WineList, error) {
	var wineList WineList
	if params == "" {
		return wineList, errors.New("params cannot be empty")
	}
	requestUrl := CatalogUrl + "?apikey=" + self.apiKey + "&" + params
	resp, err := http.Get(requestUrl)

	if err != nil {
		return wineList, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return wineList, err
	}
	body = fixJson(body)

	var productResponse ProductResponse
	err = json.Unmarshal(body, &productResponse)

	if err != nil {
		return wineList, err
	} else if productResponse.Status.ReturnCode != 0 {
		return wineList, errors.New(strings.Join(productResponse.Status.Messages, ","))
	} else {
		wineList = productResponse.Wines
		return wineList, nil
	}
}

// This is a somewhat ugly hack, because I was having issues with the unquote solution
// found at http://golang.org/src/encoding/json/decode.go. If anyone has a better solution,
// I would gladly accept a PR
func fixJson(jsonArr []byte) []byte {
	return bytes.Replace(jsonArr, []byte{'&', 'a', 'm', 'p', ';'}, []byte{'&'}, -1)
}


