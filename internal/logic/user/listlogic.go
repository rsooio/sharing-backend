package user

import (
	"context"
	"encoding/json"

	"backend/internal/pkg/errno"
	"backend/internal/svc"
	"backend/internal/types"
	"backend/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	uid    int64
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	uid, _ := ctx.Value("uid").(json.Number).Int64()

	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		uid:    uid,
	}
}

func (l *ListLogic) List(req *types.Paginator) (resp *types.DetailList, err error) {
	role, _ := l.ctx.Value("role").(json.Number).Int64()
	if role != 0 {
		return nil, errno.AccessDenied
	}

	var ret types.DetailList

	err = mr.Finish(func() (err error) {
		var list []*model.User
		offset := (req.PageNumber + 1) * req.PageSize
		err = l.svcCtx.DB.Offset(offset).Limit(req.PageSize).Find(&list).Error
		if err != nil {
			return
		}

		ret.DetailList = make([]types.Detail, 0, len(list))
		for i, v := range list {
			ret.DetailList[i] = types.Detail{
				ID: types.ID{ID: v.ID},
				Resource: types.Resource{
					Role:     int(v.Role),
					Mobile:   v.Mobile,
					Username: v.Username,
				},
			}
		}
		return
	}, func() error {
		return l.svcCtx.DB.Where(&model.User{}).Count(&ret.Total).Error
	})

	return &ret, err
}
