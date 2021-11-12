package logic

import (
	"Linkux/dao/mysql"
	"Linkux/models"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
)

var (
	id         = `wxb27cb3df6158fc0e`
	secret     = `001a15eeb9b2e60acfb21ce896e3885f`
	grant_type = `authorization_code`
)

func Login(code string) (user *models.User, err error) {
	resp, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=" + id + `&secret=` + secret + `&js_code=` + code + `&grant_type=` + grant_type)
	if err != nil {
		zap.L().Error("Get weixin API failed", zap.Error(err))
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		zap.L().Error("Read from body err: ", zap.Error(err))
		return
	}
	var res models.LoginResBody
	err = json.Unmarshal(body, &res)
	if err != nil {
		zap.L().Error("json.Unmarshal failed:", zap.Error(err))
		return
	}

	user = &models.User{
		UserID:       res.OpenID,
		Contribution: 0,
	}

	return user, mysql.InsertUser(user)
}
