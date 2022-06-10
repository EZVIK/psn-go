package psn_go

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/schema"
	"io/ioutil"
	"net/http"
	urlx "net/url"
	"strings"
)

const (
	NPSSO = ``

	AUTH_BASE_URL = `https://ca.account.sony.com/api/authz/v3/oauth`

	REDIRECT_URI = `com.playstation.PlayStationApp://redirect`

	AUTHORIZATION = `YWM4ZDE2MWEtZDk2Ni00NzI4LWIwZWEtZmZlYzIyZjY5ZWRjOkRFaXhFcVhYQ2RYZHdqMHY`

	AUTHORIZATION_BEARER = `bearer `

	AUTHORIZATION_BASIC = `Basic `
)

func ExchangeForAuthToken(code string) (Access string, Refresh string, err error) {

	url := AUTH_BASE_URL + "/token"
	authorization := AUTHORIZATION_BASIC + AUTHORIZATION

	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		SetHeader("Authorization", authorization).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"code":         code,
			"redirect_uri": REDIRECT_URI,
			"grant_type":   "authorization_code",
			"token_format": "jwt",
		}).
		Post(url)

	if err != nil {
		return
	}

	authToken := new(AuthTokensResponse)
	if err = json.Unmarshal(resp.Body(), authToken); err != nil {
		return
	}

	return authToken.AccessToken, authToken.RefreshToken, nil
}

func ExchangeForCode(npsso string) (code string, err error) {

	//requestUrl := fmt.Sprintf("%s/authorize?%s", AUTH_BASE_URL, "access_type=offline&client_id=ac8d161a-d966-4728-b0ea-ffec22f69edc&redirect_uri=com.playstation.PlayStationApp://redirect&response_type=code&scope=psn:mobile.v1 psn:clientapp")
	url := "https://ca.account.sony.com/api/authz/v3/oauth/authorize?access_type=offline&client_id=ac8d161a-d966-4728-b0ea-ffec22f69edc&redirect_uri=com.playstation.PlayStationApp%3A%2F%2Fredirect&response_type=code&scope=psn%3Amobile.v1%20psn%3Aclientapp"

	//url := requestUrl
	method := "GET"

	client := &http.Client{
		// manual redirect
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return
	}
	req.Header.Add("Cookie", fmt.Sprintf("npsso=%s;", npsso))

	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	location := res.Header.Get("Location")
	if !strings.Contains(location, "redirect/?") {
		return "", errors.New("invalid npsso code")
	}
	respInfo := strings.Split(location, "redirect/?")[1]

	var decoder = schema.NewDecoder()

	CodeInfo := new(CodeResponse)

	v, err := urlx.ParseQuery(respInfo)
	if err != nil {
		return "", err
	}

	if err = decoder.Decode(CodeInfo, v); err != nil {
		return "", err
	}

	return CodeInfo.Code, nil
}
