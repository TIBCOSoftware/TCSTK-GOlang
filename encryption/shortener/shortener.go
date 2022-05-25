/*
* BSD 3-Clause License
* Copyright © 2022. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
*/
package shortener

import (
	"time"

	"github.com/TIBCOSoftware/TCSTK-GOlang/tibcocloud/rest"
)

// type FirebaseShortLinksRequest struct {
// 	LongDynamicLink string `json:"longDynamicLink"`
// 	Suffix          struct {
// 		Option string `json:"option"`
// 	} `json:"suffix"`
// }

// type FirebaseShortLinksResponse struct {
// 	ShortLink   string `json:"shortLink"`
// 	PreviewLink string `json:"previewLink"`
// }

type BitlyShortLinksRequest struct {
	Domain  string `json:"domain"`
	LongURL string `json:"long_url"`
}

type BitlyShortLinksResponse struct {
	CreatedAt      string        `json:"created_at"`
	ID             string        `json:"id"`
	Link           string        `json:"link"`
	CustomBitlinks []interface{} `json:"custom_bitlinks"`
	LongURL        string        `json:"long_url"`
	Archived       bool          `json:"archived"`
	Tags           []interface{} `json:"tags"`
	DeepLinks      []interface{} `json:"deepLink"`
	References     struct {
		Group string `json:"group"`
	} `json:"references"`
}

// func firebaseRequest(apikey string, longURL string) (request rest.Input) {

// 	queryParams := url.Values{}
// 	queryParams.Set("apikey", apikey)

// 	var body ShortLinksRequest
// 	body.LongDynamicLink = longURL
// 	body.Suffix.Option = "SHORT"

// 	request = rest.Input{
// 		URL:         "https://firebasedynamiclinks.googleapis.com/v1/shortLinks",
// 		HTTPMethod:  "POST",
// 		ContentType: "application/json",
// 		QueryParams: queryParams,
// 		Data:        body,
// 		Response:    &FirebaseShortLinksResponse{},
// 		Error:       &FirebaseShortLinksResponse{},
// 	}

// }

func Shortener(security string, provider string, longURL string) (shortURL string, err error) {
	var request rest.Input
	// if a.settings.Provider == "Firebase" {
	// 	request = firebaseRequest(a.settings.SecurityToken, input.LongDynamicLink)
	// }

	if provider == "Bitly" {
		request = bitlyRequest(security, longURL)
	}

	if provider == "Rebrandly" {
		request = rebrandlyRequest(security, longURL, "")
	}

	restResponse, err := rest.MakeCall(request)
	if err != nil {
		return "", err
	}

	var shortLink string
	if provider == "Bitly" {
		shortLink = bitlyResponse(restResponse)
	}

	if provider == "Rebrandly" {
		shortLink = rebrandlyResponse(restResponse)
	}

	return shortLink, nil
}

func bitlyRequest(apikey string, longURL string) (request rest.Input) {

	var body BitlyShortLinksRequest
	body.Domain = "bit.ly"
	body.LongURL = longURL

	request = rest.Input{
		URL:            "https://api-ssl.bitly.com/v4/shorten",
		HTTPMethod:     "POST",
		Authentication: apikey,
		ContentType:    "application/json",
		Data:           body,
		Response:       &BitlyShortLinksResponse{},
		Error:          &BitlyShortLinksResponse{},
	}
	return
}

func bitlyResponse(restResponse rest.Output) (url string) {

	var response BitlyShortLinksResponse
	if restResponse.Status == 200 {
		response = *(restResponse.Data).(*BitlyShortLinksResponse)
	} else {
		response = *(restResponse.Data).(*BitlyShortLinksResponse)
	}
	return response.Link
}

type Rebrandly struct {
	ID          string          `json:"id"`
	Title       string          `json:"title"`
	Slashtag    string          `json:"slashtag"`
	Destination string          `json:"destination"`
	ShortURL    string          `json:"shortUrl"`
	Domain      RebrandlyDomain `json:"domain"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
	Clicks      int             `json:"clicks"`
	LastClickAt time.Time       `json:"lastClickAt"`
	Favourite   bool            `json:"favourite"`
}
type RebrandlyDomain struct {
	ID       string `json:"id"`
	FullName string `json:"fullName"`
}

func rebrandlyRequest(apikey string, longURL string, domain string) (request rest.Input) {

	var body Rebrandly
	body.Destination = longURL

	if domain != "" {
		var domainRequest RebrandlyDomain
		domainRequest.FullName = domain
		body.Domain = domainRequest
	}

	headers := make(map[string]string)
	headers["apikey"] = apikey

	request = rest.Input{
		URL:         "https://api.rebrandly.com/v1/links",
		HTTPMethod:  "POST",
		Headers:     headers,
		ContentType: "application/json",
		Data:        body,
		Response:    &Rebrandly{},
		Error:       &Rebrandly{},
	}
	return
}

func rebrandlyResponse(restResponse rest.Output) (url string) {

	var response Rebrandly
	if restResponse.Status == 200 {
		response = *(restResponse.Data).(*Rebrandly)
	} else {
		response = *(restResponse.Data).(*Rebrandly)
	}
	return response.ShortURL
}
