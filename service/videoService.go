package service

import (
	"context"
	"log"
	"server04/config"
	"server04/ent"
	"server04/ent/usersec"
)

type VideosOps struct {
	ctx    context.Context
	client *ent.Client
}

func NewvideosOps(ctx context.Context) *VideosOps {
	dbClient, _ := config.GetClients()
	return &VideosOps{
		ctx:    ctx,
		client: dbClient,
	}
}

func (r *VideosOps) VideoCreate(newVideo ent.Videos, userID int) (*ent.Videos, error) {
	user, _ := r.client.UserSec.Query().Where(
		usersec.ID(userID),
	).Only(r.ctx)
	newCreatedUser, err := r.client.Videos.Create().
		SetDesc(newVideo.Desc).
		SetVideoLink(newVideo.VideoLink).
		SetThumb(newVideo.Thumb).
		SetLikeNum(0).
		SetCommentNum(0).
		SetUser(user).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return newCreatedUser, nil
}

func (r *VideosOps) VideoGetByID(id int) (*ent.Videos, error) {

	user, err := r.client.Videos.Get(r.ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

/*
func (r *UserOps) UserGetByEmail(email string) (*ent.UserSec, error) {

		user, err := r.client.UserSec.GetByEmail(r.ctx, email)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
*/
func (r *VideosOps) VideoGetAll() (int, error) {

	users, _ := r.client.Videos.Query().Select().Count(r.ctx)

	return users, nil
}

func (r *VideosOps) VideoGetBatch(minLast int, numberRequested int) ([]*ent.Videos, error) {

	videos, _ := r.client.Videos.GetBatch(r.ctx, minLast, numberRequested)

	return videos, nil
}

func (r *VideosOps) VideoGetSearch(searchTitelString string) ([]*ent.Videos, error) {

	videos, _ := r.client.Videos.GetByVideoTitle(r.ctx, searchTitelString)

	return videos, nil
}

func (r *VideosOps) AddVideoLikeIfVideoDoesNotHaveLikeFromUser(userId int, videoId int) bool {
	video, _ := r.client.Videos.Get(r.ctx, videoId)
	yes, err := r.client.Videos.QueryLikeId(video).QueryUser().Where(usersec.ID(userId)).Count(r.ctx)
	log.Println(yes, err)
	if yes > 0 {
		return true
	}
	r.client.Videos.Update().AddLikeNum(int64(video.LikeNum) + 1).Save(r.ctx)
	return false
}
