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

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	uid    int64
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	uid, _ := ctx.Value("uid").(json.Number).Int64()

	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		uid:    uid,
	}
}

func (l *DeleteLogic) Delete(req *types.ID) (err error) {
	role, _ := l.ctx.Value("role").(json.Number).Int64()
	if role != 0 && req.ID != uint(l.uid) {
		return errno.AccessDenied
	}

	return l.svcCtx.DB.Delete(&model.User{Model: gorm.Model{ID: req.ID}}).Error
}
