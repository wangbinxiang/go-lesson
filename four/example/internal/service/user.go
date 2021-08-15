package service

import (
	"context"
	"strconv"

	"github.com/pkg/errors"
	v1 "github.com/wangbinxiang/go-lesson-four/api/blog/v1"
	"github.com/wangbinxiang/go-lesson-four/internal/biz"
	"google.golang.org/grpc/metadata"
)

type UserService struct {
	v1.UserServerServer
	usecase biz.UserUsecase
}

var MetaDataErr = errors.New("get metadata err")

func (u *UserService) GetUserInfo(ctx context.Context, req *v1.GetUserInfoRequest) (*v1.GetUserInfoResponse, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, MetaDataErr
	}

	data := meta.Get("id")
	if len(data) != 1 {
		return nil, errors.Wrapf(MetaDataErr, "id lens error, metadata: %v", data)
	}

	id, err := strconv.Atoi(data[0])
	if err != nil {
		return nil, errors.Wrapf(MetaDataErr, "id not num, id: %v", id)
	}

	user, err := u.usecase.GetUserInfo(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := &v1.GetUserInfoResponse{
		Id:       int32(user.Id),
		Username: user.Name,
		City:     user.City,
	}

	return resp, nil
}

func NewUserService(uc biz.UserUsecase) *UserService {
	return &UserService{usecase: uc}
}
