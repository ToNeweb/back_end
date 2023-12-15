package service

import (
	"context"
	"server04/config"
	"server04/ent"
)

type UserOps struct {
	ctx    context.Context
	client *ent.Client
}

func NewUserOps(ctx context.Context) *UserOps {
	return &UserOps{
		ctx:    ctx,
		client: config.GetClient(),
	}
}

func (r *UserOps) UserCreate(newUser ent.UserSec) (*ent.UserSec, error) {

	newCreatedUser, err := r.client.UserSec.Create().
		SetEmail(newUser.Email).
		SetPassword(newUser.Password).
		Save(r.ctx)

	if err != nil {
		return nil, err
	}

	return newCreatedUser, nil
}

func (r *UserOps) UserGetByID(id int) (*ent.UserSec, error) {

	user, err := r.client.UserSec.Get(r.ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserOps) UserGetByEmail(email string) (*ent.UserSec, error) {

	user, err := r.client.UserSec.GetByEmail(r.ctx, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserOps) UserGetALL() (int, error) {

	users, _ := r.client.UserSec.Query().Select().Count(r.ctx)

	return users, nil
}
