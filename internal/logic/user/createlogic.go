package user

import (
	"context"

	"backend/internal/pkg/errno"
	"backend/internal/pkg/res"
	"backend/internal/svc"
	"backend/internal/types"
	"backend/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.Resource) (resp *types.ID, err error) {
	claims := *l.svcCtx.Claims
	if !claims.Has(res.UserCreate) {
		return nil, errno.AccessDenied
	}

	// password, err := bcrypt.Encrypt(req.)
	user := model.User{
		Username: req.Username,
		Mobile:   req.Mobile,
	}
	err = l.svcCtx.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &types.ID{ID: user.ID}, nil
}
