package public

import (
	"context"
	"errors"
	"time"

	"backend/internal/pkg/res"
	"backend/internal/svc"
	"backend/internal/types"
	"backend/internal/utils/auth"
	"backend/internal/utils/jwt"
	"backend/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"gorm.io/gorm"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	user, err := l.login(req)
	if err != nil {
		return nil, err
	}

	var roles []*model.UserRole
	err = l.svcCtx.DB.Where(&model.UserRole{UserID: user.ID}).First(&roles).Error
	if err != nil {
		return nil, err
	}

	var resources res.ResourceFlag = res.ResourceFlag(^uint64(0))
	if !auth.IsRoot(user.ID) {
		resources, err = mr.MapReduce(func(source chan<- uint) {
			for _, v := range roles {
				source <- v.RoleID
			}
		}, func(roleID uint, writer mr.Writer[res.ResourceFlag], cancel func(error)) {
			role := model.Role{Model: gorm.Model{ID: roleID}}
			err := l.svcCtx.DB.Find(&role).Error
			if err != nil {
				cancel(err)
			}
			writer.Write(role.Resources)
		}, func(pipe <-chan res.ResourceFlag, writer mr.Writer[res.ResourceFlag], cancel func(error)) {
			var resources res.ResourceFlag = 0
			for flags := range pipe {
				resources |= flags
			}
			writer.Write(resources)
		})
		if err != nil {
			return nil, err
		}
	}

	token, err := jwt.Sign(user.ID, resources, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		Token:     token,
		ExpireAt:  l.svcCtx.Config.Auth.AccessExpire,
		Resources: uint64(resources),
	}, nil
}

func (l *LoginLogic) login(req *types.LoginReq) (user *model.User, err error) {
	if req.Type == "sms" {
		err = l.verifyOTP(req.Mobile, req.Secret)
		if err != nil {
			return nil, err
		}
	}

	err = l.svcCtx.DB.Where(&model.User{Mobile: req.Mobile}).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		user = &model.User{Username: req.Mobile, Mobile: req.Mobile}
		err = l.svcCtx.DB.Create(&user).Error
	}
	if err != nil {
		return nil, err
	}

	if req.Type == "pwd" {
		return nil, errors.New("not implement")
	}

	return user, nil
}

func (l *LoginLogic) verifyOTP(mobile string, OTP string) (err error) {
	if otp, ok := l.svcCtx.OTPList[mobile]; !ok {
		err = errors.New("record not found")
	} else if time.Since(otp.Time) > time.Minute*5 {
		err = errors.New("record expired")
		delete(l.svcCtx.OTPList, mobile)
	} else if otp.OTP != OTP {
		err = errors.New("verification failed")
	} else {
		delete(l.svcCtx.OTPList, mobile)
	}
	return
}

// func (l *LoginLogic) getJwtToken(user *model.User, resources uint64) (string, error) {
// 	claims := make(jwt.MapClaims)
// 	iat := time.Now().Unix()
// 	claims["exp"] = iat + l.svcCtx.Config.Auth.AccessExpire
// 	claims["iat"] = iat
// 	claims["uid"] = user.ID
// 	claims["res"] = resources
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	token.Claims = claims
// 	return token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
// }

func ErrorCatcher0(err error, f func() error) error {
	if err != nil {
		return err
	}

	return f()
}

func ErrorCatcher1[T any](err error, f func() (T, error)) (ret T, e error) {
	if err != nil {
		return ret, err
	}

	return f()
}
