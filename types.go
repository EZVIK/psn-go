package psn_go

import "time"

type AuthTokensResponse struct {
	AccessToken           string `json:"access_token"`
	TokenType             string `json:"token_type"`
	ExpiresIn             int    `json:"expires_in"`
	Scope                 string `json:"scope"`
	IDToken               string `json:"id_token"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"`
}

type CodeResponse struct {
	Code string `schema:"code"`
	Cid  string `schema:"cid"`
}

// PlayerGameTrophyInfo
// struct of the latest trophies' player recently achieved
type PlayerGameTrophyInfo struct {
	NpServiceName       string `json:"npServiceName"`     // Service Type  PS5 Games: trophy2,  others : trophy
	NpCommunicationID   string `json:"npCommunicationId"` // kind of Game ID
	TrophySetVersion    string `json:"trophySetVersion"`
	TrophyTitleName     string `json:"trophyTitleName"`
	TrophyTitleIconURL  string `json:"trophyTitleIconUrl"`
	TrophyTitlePlatform string `json:"trophyTitlePlatform"`
	HasTrophyGroups     bool   `json:"hasTrophyGroups"`
	DefinedTrophies     struct {
		Bronze   int `json:"bronze"`
		Silver   int `json:"silver"`
		Gold     int `json:"gold"`
		Platinum int `json:"platinum"`
	} `json:"definedTrophies"`
	Progress       int `json:"progress"`
	EarnedTrophies struct {
		Bronze   int `json:"bronze"`
		Silver   int `json:"silver"`
		Gold     int `json:"gold"`
		Platinum int `json:"platinum"`
	} `json:"earnedTrophies"`
	HiddenFlag          bool      `json:"hiddenFlag"`
	LastUpdatedDateTime time.Time `json:"lastUpdatedDateTime"`
}

// PlayerTrophyList http Response
type PlayerTrophyList struct {
	TrophyTitles []PlayerGameTrophyInfo `json:"trophyTitles"`
}

type Trophy struct {
	TrophyID      int    `json:"trophyId"`
	TrophyHidden  bool   `json:"trophyHidden"`
	TrophyType    string `json:"trophyType"`
	TrophyName    string `json:"trophyName"`
	TrophyDetail  string `json:"trophyDetail"`
	TrophyIconURL string `json:"trophyIconUrl"`
	TrophyGroupID string `json:"trophyGroupId"`
}

type TrophyStatus struct {
	TrophyID         int    `json:"trophyId"`
	TrophyHidden     bool   `json:"trophyHidden"`
	Earned           bool   `json:"earned"`
	TrophyType       string `json:"trophyType"`
	TrophyRare       int    `json:"trophyRare"`
	TrophyEarnedRate string `json:"trophyEarnedRate"`
	EarnedDateTime   string `json:"earnedDateTime"`
}

type UserTrophy struct {
	ID                  int    `gorm:"primary_key" json:"id"`
	GameTitle           string `json:"game_title"`
	TrophyID            int    `json:"trophyId"`
	TrophyName          string `json:"trophyName"`
	TrophyTitlePlatform string `json:"trophy_title_platform"`
	TrophyType          string `json:"trophyType"`
	TrophyDetail        string `json:"trophyDetail"`
	TrophyIconURL       string `json:"trophyIconUrl"`
	EarnedDateTime      string `json:"earnedDateTime"`
	TrophyEarnedRate    string `json:"trophyEarnedRate"`
}

type GameTrophyStatusList struct {
	TrophySetVersion string         `json:"trophySetVersion"`
	HasTrophyGroups  bool           `json:"hasTrophyGroups"`
	TotalItemCount   int            `json:"totalItemCount"`
	Trophies         []TrophyStatus `json:"trophies"`
}

type GameTrophyList struct {
	TrophySetVersion string   `json:"trophySetVersion"`
	HasTrophyGroups  bool     `json:"hasTrophyGroups"`
	TotalItemCount   int      `json:"totalItemCount"`
	Trophies         []Trophy `json:"trophies"`
}

type UserTrophies struct {
	ID                string
	UserID            string
	NpCommunicationId string
	TrophyID          string
	EarnedDateTime    string
}

type Top10Trophy struct {
}
