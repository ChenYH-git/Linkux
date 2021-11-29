package models

const (
	OrderTime  = "time"  //时间排序
	OrderScore = "score" //分数排序
)

//定义请求参数的结构体
// 投票
type ParamVoteData struct {
	PostID    int64 `json:"post_id,string" binding:"required"`                   // 帖子id
	Direction int8  `json:"direction,string" binding:"oneof=1 0 -1" example:"1"` // 赞同（1） 反对（-1） 取消（0）(required在默认为0时，会主动过滤）
}

// ParamPostList 获取帖子列表query string参数
type ParamPostList struct {
	LabelID int64  `json:"label_id" form:"label_id"`           // 帖子社区标签id，可以为空
	TransID int64  `json:"trans_id" form:"trans_id"`           // 帖子对应的翻译任务帖id，可以为空，查看已有翻译时必须保证不为0
	Page    int64  `json:"page" form:"page"`                   // 分页信息，可以为空，默认从1开始
	Size    int64  `json:"size" form:"size"`                   // 分页大小，可以为空，默认大小10
	Order   string `json:"order" form:"order" example:"score"` // 排序方式，可以为空，默认为score，可以为time
	Search  string `json:"search" form:"search"`               // 搜索内容，可以为空，搜索时必填
}

type VCorNot struct {
	Voted     bool `json:"voted"`     // 1为点过赞，0为没有
	Collected bool `json:"collected"` // 1为收藏过，0为没有
}

type PUorNot struct {
	Qualified bool `json:"qualified"` // 1为加过精，0为没有
}
