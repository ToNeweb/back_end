package service

import (
	"context"
	"server04/config"
	"server04/ent"
)

type LikeOps struct {
	ctx    context.Context
	client *ent.Client
}

func NewLikesOps(ctx context.Context) *LikeOps {
	dbClient, _ := config.GetClients()
	return &LikeOps{
		ctx:    ctx,
		client: dbClient,
	}
}

func (r *LikeOps) PutLike(userId int, videoId int) (*ent.Likes, error) {
	return r.client.Likes.Create().
		AddUserIDs(userId).
		AddVideoIDs(videoId).
		Save(r.ctx)
}
