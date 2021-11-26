package controllers

import "Linkux/models"

type _ResponsePostList struct {
	code ResCode                 `json:"code"`
	msg  string                  `json:"msg"`
	data *[]models.ApiPostDetail `json:"data"`
}

type _ResponseTransList struct {
	code ResCode         `json:"code"`
	msg  string          `json:"msg"`
	data *[]models.Trans `json:"data"`
}

type _ResponseFollowList struct {
	code ResCode        `json:"code"`
	msg  string         `json:"msg"`
	data *[]models.User `json:"data"`
}

type _ResponseMsg struct {
	code ResCode       `json:"code"`
	msg  string        `json:"msg"`
	data []interface{} `json:"data"`
}

type _RankResponseMsg struct {
	code ResCode       `json:"code"`
	msg  string        `json:"msg"`
	data []interface{} `json:"data"`
}

type _ResponseUsr struct {
	code  ResCode `json:"code"`
	msg   string  `json:"msg"`
	token string  `json:"token"`
}

type _ResponseVC struct {
	code ResCode         `json:"code"`
	msg  string          `json:"msg"`
	vc   *models.VCorNot `json:"vc"`
}
