package authenticate

import (
	"fmt"
	"github.com/EZVIK/psn-go"
	"github.com/gorilla/schema"
	"io/ioutil"
	"net/http"
	urlx "net/url"
	"strings"
)

type CodeResponse struct {
	Code string `schema:"code"`
	Cid  string `schema:"cid"`
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
		return "", psn_go.ErrNPSSO
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
