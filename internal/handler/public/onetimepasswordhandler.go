package public

import (
	"net/http"

	"backend/internal/logic/public"
	"backend/internal/svc"
	"backend/internal/types"

	"backend/internal/pkg/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func OneTimePasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OtpReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := public.NewOneTimePasswordLogic(r.Context(), svcCtx)
		resp, err := l.OneTimePassword(&req)
		response.Response(w, resp, err)
	}
}
