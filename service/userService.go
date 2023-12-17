package service

import (
	"context"
	"server04/config"
	"server04/ent"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type UserOps struct {
	ctx         context.Context
	client      *ent.Client
	redisClient *redis.Client
}

func NewUserOps(ctx context.Context) *UserOps {
	dbClient, redisClientGotten := config.GetClients()
	return &UserOps{
		ctx:         ctx,
		client:      dbClient,
		redisClient: redisClientGotten,
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

func (r *UserOps) UserCreateWithValidationEmail(newUserEmail string) (*ent.UserSec, error) {

	newCreatedUser, err := r.client.UserSec.Create(). /// check without setting password, does it create security problem?
								SetEmail(newUserEmail).
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

func (r *UserOps) UserAddValidateCode(email string, validationCode int) {
	err := r.redisClient.Set(r.ctx, email, validationCode, 3600000000000).Err()
	if err != nil {
		panic(err)
	}
}
func (r *UserOps) UserGetValidateCode(email string) int { /// should be deleted from redis?
	value, err := r.redisClient.Get(r.ctx, email).Result()
	if err != nil {
		panic(err)
	}
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return valueInt
}
