
package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type JSON map[string]interface{}

type Player struct {
	ID		int
	GameID	int
	GameName	string
	Nickname string
	School  string
	Sid		int
	Status  int
	Password string
}

type Score struct {
	ID		int
	GameID	int
	Phase	int
	Round	int
	Point 	float64
}

type Group struct {
	ID		int
	GameID	int
	Round	int
	GroupID	int
}

var db *gorm.DB

// Admin Routers

func AdminHandler(ctx iris.Context) {
	sid, school := GetSidAndSchool(ctx.GetHeader("Authorization"))
	if sid != 0 {
		ctx.StatusCode(iris.StatusUnauthorized)
		_, _ = ctx.JSON(JSON{
			"msg": "not admin",
		})
		return
	}
	ctx.Values().Set("sid", sid)
	ctx.Values().Set("school", school)
	ctx.Next()
}

func GetPlayers(ctx iris.Context) {
	school := ctx.Values().GetString("school")
	var players []Player
	db.Where("school = ?", school).Find(&players)
	var j []JSON
	for _, player := range players {
		bound := player.Nickname != ""
		ele := JSON{
			"sid": player.Sid,
			"bound": bound,
			"nickname": player.Nickname,
			"game_id": player.GameID,
			"game_name": player.GameName,
		}
		if !bound {
			ele["password"] = player.Password
		}
		j = append(j, ele)
	}
	_, _ = ctx.JSON(j)
}

func ResetPlayer(ctx iris.Context) {
	id, err := ctx.URLParamInt("sid")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(JSON{
			"msg": "json error",
		})
		return
	}
	school := ctx.Values().GetString("school")
	player := Player{}
	db.Where("school = ? and sid = ?", school, id).First(&player)
	player.GameID = 0
	player.GameName = ""
	player.Nickname = ""
	player.Password = randomPassword()
	db.Save(&player)
	_, _ = ctx.JSON(JSON{
		"username": stringify(player.Sid, school),
		"password": player.Password,
	})
}

func ApplyNewPlayer(ctx iris.Context) {
	school := ctx.Values().GetString("school")
	var players []Player
	db.Where("school = ?", school).Find(&players)
	if len(players) > 10 {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(JSON{
			"msg": "full team",
		})
		return
	}
	sid := len(players)
	password := randomPassword()
	player := Player{
		School: school,
		Sid: sid,
		Password: password,
	}
	db.Create(&player)
	_, _ = ctx.JSON(JSON{
		"username": stringify(sid, school),
		"password": password,
	})
}

// Player Routers

func CheckInStatus(ctx iris.Context) {
	sid, _ := ctx.Values().GetInt("sid")
	school := ctx.Values().GetString("school")
	player, _ := queryPlayerBySidAndSchool(sid, school)
	_, _ = ctx.JSON(JSON{
		"status": player.Status,
	})
}

func CheckIn(ctx iris.Context) {
	sid, _ := ctx.Values().GetInt("sid")
	school := ctx.Values().GetString("school")
	player, _ := queryPlayerBySidAndSchool(sid, school)
	player.Status = 1
	db.Save(&player)
	_, _ = ctx.JSON(JSON{
		"match_name": "128进96",
	})
}

func AuthHandler(ctx iris.Context) {
	sid, school := GetSidAndSchool(ctx.GetHeader("Authorization"))
	if sid == -1 {
		ctx.StatusCode(iris.StatusUnauthorized)
		_, _ = ctx.JSON(JSON{
			"msg": "not login",
		})
		return
	}
	ctx.Values().Set("sid", sid)
	ctx.Values().Set("school", school)
	ctx.Next()
}

func Status(ctx iris.Context) {
	sid, _ := ctx.Values().GetInt("sid")
	school := ctx.Values().GetString("school")
	player, _ := queryPlayerBySidAndSchool(sid, school)
	_, _ = ctx.JSON(JSON{
		"sid": sid,
		"school": school,
		"nickname": player.Nickname,
		"game_id": player.GameID,
		"game_name": player.GameName,
		"check_in": player.Status,
	})
}

func Bind(ctx iris.Context) {
	j := JSON{}
	err := ctx.ReadJSON(&j)
	if err != nil || j["id"] == nil || j["game_name"] == nil || j["nickname"] == nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(JSON{
			"msg": "json error",
		})
		return
	}
	sid, _ := ctx.Values().GetInt("sid")
	school := ctx.Values().GetString("school")
	player, err := queryPlayerBySidAndSchool(sid, school)
	if err != nil || sid == 0 {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(JSON{
			"msg": "player info error",
		})
		return
	}
	player.GameID = int(j["id"].(float64))
	player.GameName = j["game_name"].(string)
	player.Nickname = j["nickname"].(string)
	db.Save(&player)
	_, _ = ctx.JSON(JSON{
		"msg": "success",
	})
}

func ChangePwd(ctx iris.Context) {
	j := JSON{}
	err := ctx.ReadJSON(&j)
	if err != nil || j["new_pwd"] == nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(JSON{
			"msg": "json error",
		})
		return
	}
	sid, _ := ctx.Values().GetInt("sid")
	school := ctx.Values().GetString("school")
	player, err := queryPlayerBySidAndSchool(sid, school)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(JSON{
			"msg": "player info error",
		})
		return
	}
	player.Password = j["new_pwd"].(string)
	db.Save(&player)
	_, _ = ctx.JSON(JSON{
		"msg": "success",
	})
}

// Public Routers

func PlayerScores(ctx iris.Context) {
	round, _ := ctx.URLParamInt("round")
	name := ctx.URLParam("name")
	player := Player{}
	db.First(&player, "game_name = ?", name)
	var scores []Score
	db.Where("round = ? and game_id = ?", round, player.GameID).Find(&scores)
	var s []float64
	for _, score := range scores {
		s = append(s, score.Point)
	}
	var id int
	if player.Status == -1 {
		id = 0
	} else {
		id = player.GameID
	}
	_, _ = ctx.JSON(JSON{
		"id": id,
		"school": player.School,
		"sid": player.Sid,
		"nick_name": player.Nickname,
		"scores": s,
	})
}

func AllScores(ctx iris.Context) {
	round, _ := ctx.URLParamInt("round")
	var scores []Score
	db.Where("round = ?", round).Find(&scores)
	scoreMap := map[int][]float64{}
	for _, score := range scores {
		if scoreMap[score.GameID] == nil {
			scoreMap[score.GameID] = []float64{}
		}
		scoreMap[score.GameID] = append(scoreMap[score.GameID], score.Point)
	}
	var scoreJson []JSON
	for k, v := range scoreMap {
		player := Player{}
		db.First(&player, "game_id = ?", k)
		scoreJson = append(scoreJson, JSON{
			"id": k,
			"school": player.School,
			"sid": player.Sid,
			"nick_name": player.Nickname,
			"game_name": player.GameName,
			"scores": v,
			"check_in": player.Status,
		})
	}
	_, _ = ctx.JSON(scoreJson)
}

func PushScore(ctx iris.Context) {
	j := JSON{}
	err := ctx.ReadJSON(&j)
	if err != nil || j["name"] == nil || j["point"] == nil || j["round"] == nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(JSON{
			"msg": "json error",
		})
		return
	}
	name := j["name"].(string)
	point := j["point"].(float64)
	round := int(j["round"].(float64))
	player := Player{}
	var scores []Score
	db.First(&player, "game_name = ?", name)
	db.Find(&scores, "round = ? and game_id = ?", round, player.GameID)
	if player.School == "PD" {
		_, _ = ctx.JSON(JSON{
			"msg": "success",
		})
		return
	}
	phase := len(scores)
	score := Score{
		GameID: player.GameID,
		Round: round,
		Point: point,
		Phase: phase,
	}
	db.Create(&score)
	_, _ = ctx.JSON(JSON{
		"msg": "success",
	})
}

func FindAll(ctx iris.Context) {
	var players []Player
	var j []JSON
	db.Where("nickname != '' and sid != 0").Find(&players)
	for _, p := range players {
		j = append(j, JSON{
			"school": p.School,
			"nickname": p.Nickname,
			"game_name": p.GameName,
		})
	}
	_, _ = ctx.JSON(j)
}

func Login(ctx iris.Context) {
	j := JSON{}
	err := ctx.ReadJSON(&j)
	if err != nil || j["username"] == nil || j["password"] == nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(JSON{
			"msg": "json error",
		})
		return
	}
	sid, school := parseString(j["username"].(string))
	player, err := queryPlayerBySidAndSchool(sid, school)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(JSON{
			"msg": "player info error",
		})
		return
	}
	if player.Password == j["password"].(string) {
		token := BuildToken(sid, school)
		_, _ = ctx.JSON(JSON{
			"sid": sid,
			"school": school,
			"token": token,
		})
	} else {
		ctx.StatusCode(iris.StatusUnauthorized)
		_, _ = ctx.JSON(JSON{
			"msg": "bad credential",
		})
		return
	}
}

func SearchPlayerById(ctx iris.Context) {
	id, _ := ctx.URLParamInt("id")
	s, err := findPlayer(id)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(JSON{
			"msg": "can't find player",
		})
		return
	}
	_, _ = ctx.JSON(JSON{
		"name": s,
	})
}

func QueryPlayer(ctx iris.Context) {
	id, _ := ctx.URLParamInt("id")
	player := Player{}
	db.Where("id = ?", id).First(&player)
	if player.ID != 0 {
		_, _ = ctx.JSON(map[string]interface{} {
			"reg": 1,
		})
	} else {
		name, err := findPlayer(id)
		if err != nil {
			_, _ = ctx.JSON(map[string]interface{} {
				"err": 1,
			})
		} else {
			_, _ = ctx.JSON(map[string]interface{} {
				"name": name,
			})
		}
	}
}

// Functions

func randomPassword() string {
	charList := []string {"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	s := ""
	for i := 0; i < 6; i++ {
		s += charList[rand.Intn(len(charList))]
	}
	return s
}

func stringify(sid int, school string) string {
	return fmt.Sprintf("%s%03d", school, sid)
}

func parseString(s string) (int, string) {
	var idx int
	var ch rune
	for idx, ch = range s {
		if string(ch) == "0" {
			break
		}
	}
	i, _ := strconv.Atoi(s[idx:])
	s = s[0:idx]
	return i, s
}

func queryPlayerBySidAndSchool(id int, school string) (Player, error) {
	player := Player{}
	db.Where("sid = ? and school = ?", id, school).First(&player)
	if player.Password == "" {
		return player, errors.New("Can't find player")
	}
	return player, nil
}

func findPlayer(id int) (name string, err error) {
	resp, _ := http.Get(fmt.Sprintf("http://localhost:5000/get_username/%d", id))
	b, _ := ioutil.ReadAll(resp.Body)
	name = string(b)
	if len(name) >= 30 || name == "获取角色信息出错" || name == "busy" {
		return "", errors.New("can't find")
	}
	return name, nil
}

func BuildToken(sid int, school string) string {
	claims := make(jwt.MapClaims)
	claims["sid"] = sid
	claims["school"] = school
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["iat"] = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("lqynb"))
	return tokenString
}

func GetSidAndSchool(tokenString string) (int, string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("lqynb"), nil
	})
	if err != nil {
		return -1, ""
	} else if !token.Valid {
		return -1, ""
	} else {
		return int(token.Claims.(jwt.MapClaims)["sid"].(float64)), token.Claims.(jwt.MapClaims)["school"].(string)
	}
}

func init() {
	var err error
	db, err = gorm.Open("mysql", "dfy:woshisb@tcp(localhost:3306)/srm")
	if err != nil {
		fmt.Println(err)
		return
	}
	if !db.HasTable(&Player{}) {
		db.CreateTable(&Player{})
	}
	if !db.HasTable(&Score{}) {
		db.CreateTable(&Score{})
	}
}
