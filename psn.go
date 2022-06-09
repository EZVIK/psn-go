package psn_go

import (
	"github.com/EZVIK/psn-go/authenticate"
	"github.com/EZVIK/psn-go/trophy"
)

func GetAuthCode(npsso string) (string, error) {
	return authenticate.ExchangeForCode(npsso)
}

func Login(code string) (access, refresh string, err error) {
	return authenticate.ExchangeForAuthToken(code)
}

func GetTrophyGameList(accessToken string) (p []trophy.PlayerGameTrophyInfo, err error) {
	return trophy.GetTrophyGameList(accessToken)
}

func GetPlayTrophyStatus(accessToken string, npCommunicationId string, ifPs5 bool) (gl trophy.GameTrophyStatusList, err error) {
	return trophy.GetPlayTrophyStatus(accessToken, npCommunicationId, ifPs5)
}

func GetGamesTrophies(accessToken string, npCommunicationIds string, ifPs5 bool) (gl trophy.GameTrophyList, err error) {
	return trophy.GetGamesTrophies(accessToken, npCommunicationIds, ifPs5)
}

func AggregatePlayerTopTrophies(accessToken string, limit int) (m []trophy.UserTrophy, err error) {
	return trophy.AggregatePlayerTopTrophies(accessToken, limit)
}
