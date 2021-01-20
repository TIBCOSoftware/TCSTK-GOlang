/*
* BSD 3-Clause License
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */
package idm

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/TIBCOSoftware/TCSTK-common-mods/utils/cookie"
	"github.com/TIBCOSoftware/TCSTK-common-mods/utils/rest"
	"github.com/TIBCOSoftware/TCSTK-common-mods/utils/server"
)

const (
	// For V3 authentication
	authURL = "/idm/v3/login-oauth"

	// For V2 authentication
	cookieURL = "/idm/v1/login-oauth"
	tokenURL  = "https://sso-ext.tibco.com/as/token.oauth2"
)

// AccessToken for TIBCO Cloud
type AccessToken struct {
	Token   string `json:"access_token"`
	Type    string `json:"token_type"`
	Expires int    `json:"expires_in"`
}

// AuthResponse for login.
type AuthResponse struct {
	UserName  string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserID    string `json:"userId"`
	Ts        int    `json:"ts"`
	OrgName   string `json:"orgName"`
	// AccounstInfo and LoggedInUserInfo is used on Auth V2 when user belongs to multiple subscriptions
	AccountsInfo []struct {
		AccountID          string `json:"accountId"`
		AccountDisplayName string `json:"accountDisplayName"`
		LoggedInUserRole   string `json:"loggedInUserRole"`
		OwnersInfo         []struct {
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
			Email     string `json:"email"`
			Role      string `json:"role"`
		} `json:"ownersInfo"`
		SubscriptionID string `json:"subscriptionId"`
	} `json:"accountsInfo"`
	LoggedInUserInfo struct {
		FirstName    string `json:"firstName"`
		LastName     string `json:"lastName"`
		Email        string `json:"email"`
		UserEntityID string `json:"userEntityId"`
	} `json:"loggedInUserInfo"`
}

type Error struct {
	ErrorCode         string   `json:"errorCode"`
	ErrorMsg          string   `json:"errorMsg"`
	ContextAttributes []string `json:"contextAttributes"`
}

// LoginOauth returns a cookie for the region, tenant, user
func LoginOauth(connection string, region string, tenant string, username string, password string, clientID string, version string) (response cookie.Details, err error) {
	// response, ok := cookie.Get(connection)

	// if ok {
	// 	return response, nil
	// }

	if version == "V3" {
		response, err = authV3(region, tenant, username, password, clientID)
	} else {
		response, err = authV2(region, tenant, username, password)
	}

	// cookie.Set(connection, response)

	return response, err
}

func authV3(region string, tenant string, username string, password string, clientID string) (cookieDetails cookie.Details, err error) {

	data := url.Values{}
	data.Set("Email", username)
	data.Set("Password", password)
	data.Set("ClientID", clientID)
	if tenant != "" {
		data.Set("TenantId", tenant)
	}

	request := rest.Input{
		URL:         server.GetServer(region, tenant) + authURL,
		HTTPMethod:  "POST",
		ContentType: "application/x-www-form-urlencoded",
		Data:        data.Encode(),
		Response:    &AuthResponse{},
		Error:       &Error{},
	}

	response, err := rest.MakeCall(request)
	if err != nil {

	}

	tsc, domain := getTcCookie(response.Headers, true)
	cookieDetails = cookie.New(region, tenant, username, clientID, "V3", tsc, domain)
	return
}

func authV2(region string, tenant string, username string, password string) (output cookie.Details, err error) {

	accessToken, err := getToken(username, password)
	if err != nil {
		return output, err
	}

	data := url.Values{}
	data.Set("AccessToken", accessToken)
	data.Set("TenantId", tenant)

	request := rest.Input{
		URL:         server.GetServer(region, "") + "/idm/v1/login-oauth",
		HTTPMethod:  "POST",
		ContentType: "application/x-www-form-urlencoded",
		Data:        data.Encode(),
		Response:    &AuthResponse{},
		Error:       &Error{},
	}

	response, err := rest.MakeCall(request)

	if response.Status == 300 {
		// NOT SUPPORTED: User is in multiple subscriptions.
		return output, nil
		// return output, activity.NewError("Failed to get TSC and Domain cookies as the user belong to multiple subscriptions", "getCookieV2-5001", nil)
	} else if response.Status != http.StatusOK {
		return output, nil
		// return output, activity.NewError(fmt.Sprintf("Failed to get TIBCO Cloud cookies due to error: %s", resp.Status), "getCookieV2-5002", nil)
	}

	tsc, domain := getTcCookie(response.Headers, true)
	output = cookie.New(region, tenant, username, "", "V2", tsc, domain)

	return output, nil
}

func getTcCookie(headers http.Header, getValue bool) (tsc string, domain string) {
	for _, cookie := range headers["Set-Cookie"] {
		if strings.HasPrefix(cookie, "tsc") {
			tsc = cookie
		}
		if strings.HasPrefix(cookie, "domain") {
			domain = cookie
		}
	}

	if getValue {
		tsc = cookie.ExtractValue(tsc, "tsc=")
		domain = cookie.ExtractValue(domain, "domain=")
	}

	return tsc, domain
}

func getToken(username string, password string) (token string, err error) {

	data := url.Values{}
	data.Set("client_id", "ropc_ipass")
	data.Set("grant_type", "password")
	data.Set("username", username)
	data.Set("password", password)

	request := rest.Input{
		URL:         tokenURL,
		HTTPMethod:  "POST",
		ContentType: "application/x-www-form-urlencoded",
		Data:        data.Encode(),
		Response:    &AccessToken{},
	}

	response, err := rest.MakeCall(request)
	if response.Status != http.StatusOK {
		return "", nil
		// return "", activity.NewError(fmt.Sprintf("Failed to get TIBCO Cloud token due to error: %s", resp.Status), "getToken-5001", nil)
	}

	output2 := (response.Data).(*AccessToken)

	return output2.Token, nil
}
