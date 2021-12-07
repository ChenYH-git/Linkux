package mysql

import (
	"Linkux/dao/redis"
	"Linkux/models"
	"errors"
	"strconv"
	"strings"

	"go.uber.org/zap"

	"github.com/jmoiron/sqlx"
)

func InsertUser(user *models.User) (err error) {
	sqlStr := `select count(user_id) from user where user_id = ?`
	var count int
	if err := db.Get(&count, sqlStr, user.UserID); err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	sqlStr = `insert into user(user_id, username, pic_link, contribution) values(?,?,?,?)`

	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.PicLink, user.Contribution)
	return
}

func GetUserByID(uid string) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select username, pic_link from user where user_id = ?`
	err = db.Get(user, sqlStr, uid)
	if err != nil {
		return nil, err
	}

	var count int
	sqlStr = `select count(user_id) from vuser where user_id = ?`
	err = db.Get(&count, sqlStr, uid)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		user.Qualified = false
	} else {
		user.Qualified = true
	}
	return
}

func GetFollowUserByIDs(ids []*models.Follow) (user []*models.User, err error) {
	sqlStr := `select user_id, username, pic_link from user where user_id = ?`

	user = make([]*models.User, 0, len(ids))
	for _, v := range ids {
		u := new(models.User)
		err = db.Get(u, sqlStr, v.FollowID)
		user = append(user, u)
	}

	sqlStr = `select count(user_id) from vuser where user_id = ?`

	for i, _ := range user {
		var count int
		err = db.Get(&count, sqlStr, user[i].UserID)
		if err != nil {
			return nil, err
		}
		if count == 0 {
			user[i].Qualified = false
			continue
		}
		user[i].Qualified = true
	}
	return
}

func GetFollowedUserByIDs(ids []*models.Follow) (user []*models.User, err error) {
	sqlStr := `select user_id, username, pic_link from user where user_id = ?`

	user = make([]*models.User, 0, len(ids))
	for _, v := range ids {
		u := new(models.User)
		err = db.Get(u, sqlStr, v.FollowedID)
		user = append(user, u)
	}

	sqlStr = `select count(user_id) from vuser where user_id = ?`

	for i, _ := range user {
		var count int
		err = db.Get(&count, sqlStr, user[i].UserID)
		if err != nil {
			return nil, err
		}
		if count == 0 {
			user[i].Qualified = false
			continue
		}
		user[i].Qualified = true
	}
	return
}

func GetPostListOfMy(ids []string, userID string) (postList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, label_id, create_time, status
	from post
	where post_id in (?)
	and author_id = ?
	order by FIND_IN_SET(post_id, ?)
	`

	query, args, err := sqlx.In(sqlStr, ids, userID, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)

	err = db.Select(&postList, query, args...)
	if err != nil {
		return nil, err
	}

	sqlStr = `select count(post_id) from vpost where post_id = ?`

	for i, _ := range postList {
		var count int
		err = db.Get(&count, sqlStr, postList[i].PostID)
		if err != nil {
			return nil, err
		}
		if count == 0 {
			postList[i].Qualified = false
			continue
		}
		postList[i].Qualified = true
	}
	return

}

func AddCollection(p *models.Trigger, userID string) (err error) {
	sqlStr := `select author_id from post where post_id = ?`
	var id string
	if err := db.Get(&id, sqlStr, p.PostID); err != nil {
		return err
	}
	if id == userID {
		return errors.New("无法收藏自己的帖子")
	}

	sqlStr = `select count(*) from collection where post_id = ? and user_id = ?`
	var count int
	if err := db.Get(&count, sqlStr, p.PostID, userID); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("重复收藏")
	}

	sqlStr = `insert into collection(user_id, post_id) values(?, ?)`
	_, err = db.Exec(sqlStr, userID, p.PostID)
	return
}

func CheckCollect(pid, uid string) (flag bool, err error) {
	sqlStr := `select count(*) from collection where post_id = ? and user_id = ?`
	var count int
	if err := db.Get(&count, sqlStr, pid, uid); err != nil {
		return false, err
	}
	if count > 0 {
		flag = true
	} else {
		flag = false
	}
	return
}

func AddCollectNum(p *models.Trigger) (err error) {
	sqlStr := `update post set collect_num = collect_num + 1 where post_id = ?`
	_, err = db.Exec(sqlStr, p.PostID)
	return
}

func DeleteCollection(p *models.Trigger, userID string) (err error) {
	sqlStr := `select count(*) from collection where post_id = ? and user_id = ?`
	var count int
	if err := db.Get(&count, sqlStr, p.PostID, userID); err != nil {
		return err
	}
	if count == 0 {
		return errors.New("无此帖子")
	}
	sqlStr = `delete from collection where user_id = ? and post_id = ?`
	_, err = db.Exec(sqlStr, userID, p.PostID)
	return
}

func DeleteCollectNum(p *models.Trigger) (err error) {
	sqlStr := `update post set collect_num = collect_num - 1 where post_id = ?`
	_, err = db.Exec(sqlStr, p.PostID)
	return
}

func GetCollectionIDs(p *models.ParamPostList, userID string) (ids []string, err error) {
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	sqlStr := `select post_id
	from collection
	where user_id = ?
	limit ?, ?`
	var postList []*models.Post
	err = db.Select(&postList, sqlStr, userID, start, end)
	ids = make([]string, 0, len(postList))
	for _, v := range postList {
		ids = append(ids, strconv.FormatInt(v.PostID, 10))
	}
	return
}

func AddViewNum(p *models.Trigger) (err error) {
	sqlStr := `update post set viewd_num = viewd_num + 1 where post_id = ?`
	_, err = db.Exec(sqlStr, p.PostID)
	return
}

func AddFollow(p *models.Follow, userID string) (err error) {
	sqlStr := `select count(*) from follow where follow_id = ? and user_id = ?`
	var count int
	if err := db.Get(&count, sqlStr, p.FollowID, userID); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("重复关注")
	}

	sqlStr = `insert into follow(follow_id, user_id) values(?, ?)`
	_, err = db.Exec(sqlStr, p.FollowID, userID)
	if err != nil {
		return err
	}

	sqlStr = `insert into followed(followed_id, user_id) values(?, ?)`
	_, err = db.Exec(sqlStr, userID, p.FollowID)
	if err != nil {
		return err
	}

	return nil
}

func CancelFollow(p *models.Follow, userID string) (err error) {
	sqlStr := `select count(*) from follow where follow_id = ? and user_id = ?`
	var count int
	if err := db.Get(&count, sqlStr, p.FollowID, userID); err != nil {
		return err
	}
	if count == 0 {
		return errors.New("取消错误")
	}

	sqlStr = `delete from follow where follow_id = ? and user_id = ?`
	_, err = db.Exec(sqlStr, p.FollowID, userID)
	if err != nil {
		return err
	}

	sqlStr = `delete from followed where followed_id = ? and user_id = ?`
	_, err = db.Exec(sqlStr, userID, p.FollowID)
	if err != nil {
		return err
	}
	return
}

func GetFollowUser(userID string) (data []*models.Follow, err error) {
	sqlStr := `select follow_id from follow where user_id = ?`
	err = db.Select(&data, sqlStr, userID)
	return
}

func GetFollowedUser(userID string) (data []*models.Follow, err error) {
	sqlStr := `select followed_id from followed where user_id = ?`
	err = db.Select(&data, sqlStr, userID)
	return
}

func GetFollowPostByIDs(p *models.ParamPostList, IDs []string) (data []*models.ApiPostDetail, err error) {
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1

	sqlStr := `select post_id, title, content, author_id, label_id, collect_num, viewd_num, create_time, status
	from post
	where author_id in (?)
	and status = 1
	limit ?,?
	`

	query, args, err := sqlx.In(sqlStr, IDs, start, end)
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)

	err = db.Select(&data, query, args...)
	if err != nil {
		return nil, err
	}

	sqlStr = `select count(post_id) from vpost where post_id = ?`

	for i, _ := range data {
		var count int
		err = db.Get(&count, sqlStr, data[i].PostID)
		if err != nil {
			return nil, err
		}
		if count == 0 {
			data[i].Qualified = false
			continue
		}
		data[i].Qualified = true
	}
	for idx, post := range data {
		user, err := GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID() failed",
				zap.String("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}

		label, err := GetLabelDetailByID(post.LabelID)
		if err != nil {
			zap.L().Error("mysql.GetLabelDetailByID() failed",
				zap.Int64("label_id", post.LabelID),
				zap.Error(err))
			continue
		}
		voteData, err := redis.GetPostVoteDataSingle(strconv.FormatInt(post.PostID, 10))
		if err != nil {
			return nil, err
		}
		data[idx].VoteNum = voteData
		data[idx].AuthorName = user.Username
		data[idx].PicLink = user.PicLink
		data[idx].AuthorQualified = user.Qualified
		data[idx].LabelDetail = label
	}
	return
}
