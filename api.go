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
	BaseUrl			= "http://services.wine.com/api/beta2/service.svc/json"
	CatalogUrl		= BaseUrl + "/catalog"
	CategoryUrl		= BaseUrl + "/categoryMap"
	ReferenceUrl	= BaseUrl + "/reference"
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

func (self WineApi) SearchCatalog(params string) (ProductList, error) {
	var wineList ProductList
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
		wineList = productResponse.Products
		return wineList, nil
	}
}

func (self WineApi) SearchCategory(params string) ([]Category, error) {
	var categoryList []Category
	if params == "" {
		return categoryList, errors.New("params cannot be empty")
	}
	requestUrl := CategoryUrl + "?apikey=" + self.apiKey + "&" + params
	resp, err := http.Get(requestUrl)

	if err != nil {
		return categoryList, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return categoryList, err
	}
	body = fixJson(body)

	var categoryResponse CategoryResponse
	err = json.Unmarshal(body, &categoryResponse)

	if err != nil {
		return categoryList, err
	} else if categoryResponse.Status.ReturnCode != 0 {
		return categoryList, errors.New(strings.Join(categoryResponse.Status.Messages, ","))
	} else {
		categoryList = categoryResponse.Categories
		return categoryList, nil
	}
}

func (self WineApi) FilterReference(params string) (blogs []string, books []Book, vineyards []Vineyard, err error) {
	if params == "" {
		err = errors.New("params cannot be empty")
		return
	}

	requestUrl := ReferenceUrl + "?apikey=" + self.apiKey + "&" + params
	resp, err := http.Get(requestUrl)

	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	body = fixJson(body)

	var referenceResponse ReferenceResponse
	err = json.Unmarshal(body, &referenceResponse)

	if err != nil {
		return
	} else if referenceResponse.Status.ReturnCode != 0 {
		err = errors.New(strings.Join(referenceResponse.Status.Messages, ","))
		return
	} else {
		blogs = referenceResponse.Blogs
		books = referenceResponse.Books
		vineyards = referenceResponse.Vineyards
		err = nil
		return
	}
}

// This is a somewhat ugly hack, because I was having issues with the unquote solution
// found at http://golang.org/src/encoding/json/decode.go. If anyone has a better solution,
// I would gladly accept a PR
func fixJson(jsonArr []byte) []byte {
	return bytes.Replace(jsonArr, []byte{'&', 'a', 'm', 'p', ';'}, []byte{'&'}, -1)
}


