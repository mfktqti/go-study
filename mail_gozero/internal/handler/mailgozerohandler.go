package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-study/mail_gozero/internal/logic"
	"go-study/mail_gozero/internal/svc"
	"go-study/mail_gozero/internal/types"
)

func Mail_gozeroHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMail_gozeroLogic(r.Context(), svcCtx)
		resp, err := l.Mail_gozero(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
