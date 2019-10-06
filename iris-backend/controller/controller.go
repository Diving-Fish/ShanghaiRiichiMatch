
package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/errors"
	"io/ioutil"
	"net/http"
	"time"
)

type JSON map[string]interface{}

type Player struct {
	ID		int
	Name	string
	School  string
	Sid		int
	Status  int
	Password string
}

var db *gorm.DB

// Admin Routers

// Player Routers

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
	id, _ := ctx.Values().GetInt("sid")
	school := ctx.Values().GetString("school")
	_, _ = ctx.JSON(JSON{
		"id": id,
		"school": school,
	})
}

func Bind(ctx iris.Context) {
	j := JSON{}
	err := ctx.ReadJSON(&j)
	if err != nil || j["id"] == nil || j["name"] == nil {
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
	player.ID = int(j["id"].(float64))
	player.Name = j["name"].(string)
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

func Login(ctx iris.Context) {
	j := JSON{}
	err := ctx.ReadJSON(&j)
	if err != nil || j["sid"] == nil || j["school"] == nil || j["password"] == nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(JSON{
			"msg": "json error",
		})
		return
	}
	player, err := queryPlayerBySidAndSchool(int(j["sid"].(float64)), j["school"].(string))
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(JSON{
			"msg": "player info error",
		})
		return
	}
	if player.Password == j["password"].(string) {
		token := BuildToken(int(j["sid"].(float64)), j["school"].(string))
		_, _ = ctx.JSON(JSON{
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

func queryPlayerBySidAndSchool(id int, school string) (Player, error) {
	player := Player{}
	db.Where("id = ? and school = ?", id, school).First(&player)
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
}
