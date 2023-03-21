package user

import (
	"net/http"

	"backend/internal/logic/user"
	"backend/internal/svc"
	"backend/internal/types"

	"backend/internal/pkg/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Detail
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUpdateLogic(r.Context(), svcCtx)
		err := l.Update(&req)
		response.Response(w, nil, err)
	}
}
