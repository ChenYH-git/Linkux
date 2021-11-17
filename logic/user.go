package logic

import (
	"Linkux/dao/mysql"
	"Linkux/dao/redis"
	"Linkux/models"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
)

const (
	id         = `wxb27cb3df6158fc0e`
	secret     = `001a15eeb9b2e60acfb21ce896e3885f`
	grant_type = `authorization_code`
)

func Login(p *models.User) (userID string, err error) {
	resp, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=" + id + `&secret=` + secret + `&js_code=` + p.Code + `&grant_type=` + grant_type)
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

	user := &models.User{
		UserID:       res.OpenID,
		Contribution: 0,
		Username:     p.Username,
		PicLink:      p.PicLink,
	}

	return user.UserID, mysql.InsertUser(user)
}

func GetUserConByID(p *models.ParamPostList, userID string) (data []*models.ApiPostDetail, err error) {
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder return 0 id")
		return
	}

	posts, err := mysql.GetPostListOfMy(ids, userID)
	if err != nil {
		return
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}
	for idx, post := range posts {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID() failed",
				zap.String("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}

		label, err := mysql.GetLabelDetailByID(post.LabelID)
		if err != nil {
			zap.L().Error("mysql.GetLabelDetailByID() failed",
				zap.Int64("label_id", post.LabelID),
				zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:  user.Username,
			VoteNum:     voteData[idx],
			Post:        post,
			LabelDetail: label,
		}
		data = append(data, postDetail)
	}
	return
}

func AddCollection(p *models.Triger, userID string) (err error) {
	err = mysql.AddCollection(p, userID)
	if err != nil {
		zap.L().Error("mysql.AddCollection(p, userID) failed",
			zap.String("user_id", userID),
			zap.Error(err))
		return
	}

	err = mysql.AddCollectNum(p)
	if err != nil {
		zap.L().Error("mysql.AddCollectNum() failed",
			zap.String("user_id", userID),
			zap.Error(err))
		return
	}
	return nil
}

func DeleteCollection(p *models.Triger, userID string) (err error) {
	err = mysql.DeleteCollection(p, userID)
	if err != nil {
		zap.L().Error("mysql.AddCollection(p, userID) failed",
			zap.String("user_id", userID),
			zap.Error(err))
		return
	}

	err = mysql.DeleteCollectNum(p)
	if err != nil {
		zap.L().Error("mysql.AddCollectNum() failed",
			zap.String("user_id", userID),
			zap.Error(err))
		return
	}
	return nil
}

func GetCollection(p *models.ParamPostList, userID string) (data []*models.ApiPostDetail, err error) {
	ids, err := mysql.GetCollectionIDs(p, userID)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("mysql.GetCollectionIDs return 0 id")
		return
	}

	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}
	for idx, post := range posts {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID() failed",
				zap.String("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}

		label, err := mysql.GetLabelDetailByID(post.LabelID)
		if err != nil {
			zap.L().Error("mysql.GetLabelDetailByID() failed",
				zap.Int64("label_id", post.LabelID),
				zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:  user.Username,
			VoteNum:     voteData[idx],
			PicLink:     user.PicLink,
			Post:        post,
			LabelDetail: label,
		}
		data = append(data, postDetail)
	}
	return
}

func AddViewNum(p *models.Triger) (err error) {
	err = mysql.AddViewNum(p)
	if err != nil {
		zap.L().Error("mysql.AddCollectNum() failed", zap.Error(err))
		return
	}
	return nil
}

func AddFollow(p *models.Follow, userID string) (err error) {
	err = mysql.AddFollow(p, userID)
	if err != nil {
		zap.L().Error("mysql.AddFollow(p, userID) failed",
			zap.String("user_id", userID),
			zap.Error(err))
		return
	}
	return nil
}

func CancelFollow(p *models.Follow, userID string) (err error) {
	err = mysql.CancelFollow(p, userID)
	if err != nil {
		zap.L().Error("mysql.CancelFollow(p, userID) failed",
			zap.String("user_id", userID),
			zap.Error(err))
		return
	}
	return nil
}

func GetFollowUser(userID string) (data []*models.User, err error) {
	ids, err := mysql.GetFollowUser(userID)
	if err != nil {
		zap.L().Error("mysql.GetFollowUser(userID) failed",
			zap.String("user_id", userID),
			zap.Error(err))
		return nil, err
	}

	data, err = mysql.GetFollowUserByIDs(ids)
	if err != nil {
		zap.L().Error("mysql.GetFollowUserByIDs(ids) failed",
			zap.String("user_id", userID),
			zap.Error(err))
		return nil, err
	}

	return data, nil
}

func GetFollowedUser(userID string) (data []*models.User, err error) {
	ids, err := mysql.GetFollowedUser(userID)
	if err != nil {
		zap.L().Error("mysql.GetFollowedUser(userID) failed",
			zap.String("user_id", userID),
			zap.Error(err))
		return nil, err
	}

	data, err = mysql.GetFollowedUserByIDs(ids)
	if err != nil {
		zap.L().Error("mysql.GetFollowedUserByIDs(ids) failed",
			zap.String("user_id", userID),
			zap.Error(err))
		return nil, err
	}

	return data, nil
}

func GetFollowPost(p *models.ParamPostList, userID string) (data []*models.ApiPostDetail, err error) {
	ids, err := mysql.GetFollowUser(userID)
	if err != nil {
		zap.L().Error("mysql.GetFollowUser(userID) failed",
			zap.String("user_id", userID),
			zap.Error(err))
		return nil, err
	}

	IDs := make([]string, 0, len(ids))
	for _, v := range ids {
		IDs = append(IDs, v.FollowID)
	}

	data, err = mysql.GetFollowPostByIDs(p, IDs)
	if err != nil {
		zap.L().Error("mysql.GetFollowPostByIDs(p, ids) failed",
			zap.String("user_id", userID),
			zap.Error(err))
		return nil, err
	}

	return data, nil
}
