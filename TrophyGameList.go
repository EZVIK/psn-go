package main

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

const TROPHY_USER_URL = `https://m.np.playstation.net/api/trophy/v1`

// GetTrophyGameList List of all recently played games
func GetTrophyGameList(accessToken string) (p []PlayerGameTrophyInfo, err error) {

	url := TROPHY_USER_URL + "/users/me/trophyTitles"

	authorization := AUTHORIZATION_BEARER + accessToken

	client := resty.New()

	resp, err := client.R().
		SetHeader("Authorization", authorization).
		Get(url)

	if err != nil {
		return
	}

	lp := new(PlayerTrophyList)
	if err = json.Unmarshal(resp.Body(), lp); err != nil {
		return nil, err
	}

	return lp.TrophyTitles, nil

}

// GetPlayTrophyStatus Find a list of game achievements by npCommunicationId
// contain the trophy info
func GetPlayTrophyStatus(accessToken string, npCommunicationId string, ifPs5 bool) (gl GameTrophyStatusList, err error) {

	np := ""
	if !ifPs5 {
		np = "?npServiceName=trophy"
	}

	url := TROPHY_USER_URL + "/users/me/npCommunicationIds/" + npCommunicationId + "/trophyGroups/all/trophies" + np

	authorization := AUTHORIZATION_BEARER + accessToken

	client := resty.New()

	resp, err := client.R().
		SetHeader("Authorization", authorization).
		Get(url)

	if err != nil {
		return
	}

	lp := GameTrophyStatusList{}
	if err = json.Unmarshal(resp.Body(), &lp); err != nil {
		return
	}

	return lp, nil
}

// GetGamesTrophies 查询游戏内所有成就详细信息
func GetGamesTrophies(accessToken string, npCommunicationIds string, ifPs5 bool) (gl GameTrophyList, err error) {

	np := ""
	if !ifPs5 {
		np = "?npServiceName=trophy"
	}
	url := TROPHY_USER_URL + "/npCommunicationIds/" + npCommunicationIds + "/trophyGroups/all/trophies" + np

	authorization := AUTHORIZATION_BEARER + accessToken

	client := resty.New()

	resp, err := client.R().
		SetHeader("Authorization", authorization).
		Get(url)

	if err != nil {
		return
	}

	lp := GameTrophyList{}
	if err = json.Unmarshal(resp.Body(), &lp); err != nil {
		return
	}

	return lp, nil
}

// AggregatePlayerTopTrophies aggregate trophy and have earned and trophyDetail
func AggregatePlayerTopTrophies(accessToken string, limit int) (m []UserTrophy, err error) {

	m = make([]UserTrophy, 0)
	gameTrophyList, err := GetTrophyGameList(accessToken)
	if err != nil {
		return nil, err
	}

	for i, game := range gameTrophyList {
		if i >= limit {
			break
		}

		tlist, err := GetGamesTrophies(accessToken, game.NpCommunicationID, game.TrophyTitlePlatform == "PS5")
		if err != nil {
			return nil, err
		}

		tl, err := GetPlayTrophyStatus(accessToken, game.NpCommunicationID, game.TrophyTitlePlatform == "PS5")
		if err != nil {
			return nil, err
		}

		tlMap := make(map[int]Trophy)
		for _, t := range tlist.Trophies {
			tlMap[t.TrophyID] = t
		}

		for _, tr := range tl.Trophies {
			if tr.Earned {
				trophyInfo := tlMap[tr.TrophyID]
				ut := UserTrophy{
					TrophyID:            trophyInfo.TrophyID,
					GameTitle:           game.TrophyTitleName,
					TrophyName:          trophyInfo.TrophyName,
					TrophyTitlePlatform: game.TrophyTitlePlatform,
					TrophyType:          trophyInfo.TrophyType,
					TrophyDetail:        trophyInfo.TrophyDetail,
					TrophyIconURL:       trophyInfo.TrophyIconURL,
					EarnedDateTime:      tr.EarnedDateTime,
					TrophyEarnedRate:    tr.TrophyEarnedRate,
				}
				m = append(m, ut)
			}
		}

		//fmt.Println(game.TrophyTitleName, " done.")
	}

	return m, nil
}
