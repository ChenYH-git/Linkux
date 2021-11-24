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

type _ResponseMsg struct {
	code ResCode `json:"code"`
	msg  string  `json:"msg"`
}

type _RankResponseMsg struct {
	code ResCode       `json:"code"`
	msg  string        `json:"msg"`
	data []interface{} `json:"data"`
}
