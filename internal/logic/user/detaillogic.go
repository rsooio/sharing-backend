package user

import (
	"context"
	"encoding/json"

	"backend/internal/pkg/errno"
	"backend/internal/svc"
	"backend/internal/types"
	"backend/model"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	uid    int64
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	uid, _ := ctx.Value("uid").(json.Number).Int64()

	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		uid:    uid,
	}
}

func (l *DetailLogic) Detail(req *types.ID) (resp *types.Detail, err error) {
	role, _ := l.ctx.Value("role").(json.Number).Int64()
	if role != 0 && req.ID != uint(l.uid) {
		return nil, errno.AccessDenied
	}

	user := model.User{Model: gorm.Model{ID: req.ID}}
	err = l.svcCtx.DB.Where(&user).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &types.Detail{
		ID: types.ID{ID: user.ID},
		Resource: types.Resource{
			Role:     int(user.Role),
			Mobile:   user.Mobile,
			Username: user.Username,
		},
	}, nil
}
