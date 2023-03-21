package user

import (
	"context"

	"backend/internal/pkg/errno"
	"backend/internal/pkg/res"
	"backend/internal/svc"
	"backend/internal/types"
	"backend/internal/utils/auth"
	"backend/internal/utils/bcrypt"
	"backend/model"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type ChangePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswordLogic {
	return &ChangePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswordLogic) ChangePassword(req *types.PasswordChange) (err error) {
	claims := *l.svcCtx.Claims
	if auth.IsRoot(req.ID.ID) {
		return errno.AccessDenied
	}
	if !claims.Has(res.UserUpdate) && (!claims.Has(res.DetailChange) || req.ID.ID != claims.UserID) {
		return errno.AccessDenied
	}

	user := model.User{Model: gorm.Model{ID: req.ID.ID}}
	err = l.svcCtx.DB.First(&user).Error
	if err != nil {
		return err
	}

	err = bcrypt.Verify(user.Password, req.Password)
	if err != nil {
		return err
	}

	return l.svcCtx.DB.Updates(&model.User{Model: gorm.Model{ID: req.ID.ID}, Password: req.NewPassword}).Error
}
