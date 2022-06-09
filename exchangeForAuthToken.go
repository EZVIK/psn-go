package main

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

//https://my.account.sony.com/central/signin/?access_type=offline&client_id=ac8d161a-d966-4728-b0ea-ffec22f69edc&redirect_uri=com.playstation.PlayStationApp%3A%2F%2Fredirect&response_type=code&scope=psn%3Amobile.v1+psn%3Aclientapp&auth_ver=v3&error=login_required&error_code=4165&error_description=User+is+not+authenticated&no_captcha=false&cid=b4cd63ef-029b-4173-a170-06ee79ba3b2c

type AuthTokensResponse struct {
	AccessToken           string `json:"access_token"`
	TokenType             string `json:"token_type"`
	ExpiresIn             int    `json:"expires_in"`
	Scope                 string `json:"scope"`
	IDToken               string `json:"id_token"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"`
}

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
