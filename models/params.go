package models

const (
	OrderTime  = "time"  //时间排序
	OrderScore = "score" //分数排序
)

//定义请求参数的结构体
// 投票
type ParamVoteData struct {
	// UserID 从请求中获取当前用户
	PostID    string `json:"post_id" binding:"required"`              // 帖子id
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` // 赞同（1） 反对（-1） 取消（0）(required在默认为0时，会主动过滤）
}

// ParamPostList 获取帖子列表query string参数
type ParamPostList struct {
	LabelID int64  `json:"label_id" form:"label_id"`
	TransID int64  `json:"trans_id" form:"trans_id"`
	Page    int64  `json:"page" form:"page"`
	Size    int64  `json:"size" form:"size"`
	Order   string `json:"order" form:"order"`
	Search  string `json:"search" form:"search"`
}
