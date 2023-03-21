package public

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"backend/internal/svc"
	"backend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OneTimePasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOneTimePasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OneTimePasswordLogic {
	return &OneTimePasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OneTimePasswordLogic) OneTimePassword(req *types.OtpReq) (resp *types.OtpResp, err error) {
	i, err := rand.Int(rand.Reader, new(big.Int).SetInt64(1000000))
	if err != nil {
		return nil, err
	}

	var limit int
	if node, ok := l.svcCtx.OTPList[req.Mobile]; ok {
		limit = int(time.Duration(time.Minute - time.Since(node.Time)).Seconds())
	}
	if limit <= 0 {
		l.svcCtx.OTPList[req.Mobile] = &svc.OTPNode{
			Time: time.Now(),
			OTP:  i.String(),
		}
		limit = 60
	}

	// send otp to user
	fmt.Println("opt: ", i)

	return &types.OtpResp{Limit: limit}, nil
}
