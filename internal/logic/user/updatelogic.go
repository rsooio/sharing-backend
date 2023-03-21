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

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	uid    int64
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	uid, _ := ctx.Value("uid").(json.Number).Int64()

	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		uid:    uid,
	}
}

func (l *UpdateLogic) Update(req *types.Detail) (err error) {
	role, _ := l.ctx.Value("role").(json.Number).Int64()

	if role != 0 && int64(req.ID.ID) != l.uid {
		return errno.AccessDenied
	}

	return l.svcCtx.DB.Updates(&model.User{
		Model:    gorm.Model{ID: req.ID.ID},
		Role:     model.RoleType(req.Role),
		Mobile:   req.Mobile,
		Username: req.Username,
	}).Error
}
