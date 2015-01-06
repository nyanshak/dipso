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

func (self WineApi) SetApiKey(apiKey string) {
	self.apiKey = apiKey
}

// This is a somewhat ugly hack, because I was having issues with the unquote solution
// found at http://golang.org/src/encoding/json/decode.go. If anyone has a better solution,
// I would gladly accept a PR
func fixJson(jsonArr []byte) []byte {
	return bytes.Replace(jsonArr, []byte{'&', 'a', 'm', 'p', ';'}, []byte{'&'}, -1)
}

func (self WineApi) Search(params string) (ProductList, error) {
	var prodList ProductList
	if params == "" {
		return prodList, errors.New("params cannot be empty")
	}
	requestUrl := CatalogUrl + "?apikey=" + self.apiKey + "&" + params
	resp, err := http.Get(requestUrl)

	if err != nil {
		return prodList, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return prodList, err
	}
	body = fixJson(body)

	var productResponse ProductResponse
	err = json.Unmarshal(body, &productResponse)

	if err != nil {
		return prodList, err
	} else if productResponse.Status.ReturnCode != 0 {
		return prodList, errors.New(strings.Join(productResponse.Status.Messages, ","))
	} else {
		prodList = productResponse.Products
		return prodList, nil
	}

}



