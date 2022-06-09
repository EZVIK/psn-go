package trophy

import (
	auth "github.com/EZVIK/psn-go/authenticate"
	"testing"
)

var access = ``
var cid = `NPWR22859_00`

func Test_GetTrophyGameList(t *testing.T) {

	code, err := auth.ExchangeForCode(auth.NPSSO)
	if err != nil {
		t.Errorf("ExchangeForCode error: %v", err)
		return
	}
	//
	access, _, err := auth.ExchangeForAuthToken(code)
	if err != nil {
		t.Errorf("ExchangeForAuthToken error: %v", err)
		return
	}

	//t.Logf("AccessToken: %s,\n RefreshToken: %s", accessToken, refreshToken)

	list, err := GetTrophyGameList(access)
	if err != nil {
		t.Errorf("GetTrophyGameList error: %v", err)
		return
	}

	for _, l := range list {
		t.Logf("Game Title: %v, Platform: %v, npcid: %v", l.TrophyTitleName, l.TrophyTitlePlatform, l.NpCommunicationID)
	}
}

func Test_GetPlayTrophyStatus(t *testing.T) {

	code, err := auth.ExchangeForCode(auth.NPSSO)
	if err != nil {
		t.Errorf("ExchangeForCode error: %v", err)
		return
	}
	//
	access, _, err := auth.ExchangeForAuthToken(code)
	if err != nil {
		t.Errorf("ExchangeForAuthToken error: %v", err)
		return
	}

	//t.Logf("AccessToken: %s,\n RefreshToken: %s", accessToken, refreshToken)

	list, err := GetPlayTrophyStatus(access, cid, true)
	if err != nil {
		t.Errorf("GetTrophyGameList error: %v", err)
		return
	}

	for _, l := range list.Trophies {
		t.Logf("TrophyID: %v, EarnedDateTime: %v", l.TrophyID, l.EarnedDateTime)
	}
}

// 查询游戏内所有成就详细信息
func Test_GetGamesTrophies(t *testing.T) {

	tl, err := GetGamesTrophies(access, cid, true)
	if err != nil {
		t.Errorf("GetGamesTrophies error: %s", err)
		return
	}

	for _, trophy := range tl.Trophies {
		t.Logf("%v", trophy)
	}

}

func Test_AggregatePlayerTopTrophies(t *testing.T) {
	code, err := auth.ExchangeForCode(auth.NPSSO)
	if err != nil {
		t.Errorf("ExchangeForCode error: %v", err)
		return
	}
	//
	access, _, err := auth.ExchangeForAuthToken(code)
	if err != nil {
		t.Errorf("ExchangeForAuthToken error: %v", err)
		return
	}

	m, err := AggregatePlayerTopTrophies(access, 2)
	if err != nil {
		t.Errorf("GetGamesTrophies error: %s", err)
		return
	}

	t.Logf("%v", m)

}

func Test_AggregatePlayerTopTrophies_DB_INSERT(t *testing.T) {

	m, err := AggregatePlayerTopTrophies(access, 1)
	if err != nil {
		t.Errorf("GetGamesTrophies error: %s", err)
		return
	}
	t.Logf("%v", m)

	//mb, err := os.ReadFile("./game_data.txt")
	//if err != nil {
	//	t.Errorf("read data err: %v", err)
	//}
	//
	//m := new([]UserTrophy)
	//if err := json.Unmarshal(mb, m); err != nil {
	//	t.Errorf("json.Unmarshal error: %v", err)
	//}
	//
	//gormClient, err := gorm.Open(mysql.Open(datasource), &gorm.Config{})
	//if err != nil {
	//	t.Errorf("models.Setup err: %v", err)
	//}
	//
	//db := gormClient.Debug()
	//if err != nil {
	//	t.Errorf("db connection error: %v", err)
	//}
	//
	//if tx := db.Model(UserTrophy{}).Create(&m); tx.Error != nil {
	//	t.Errorf("batch insert error: %v", tx.Error)
	//}
}
