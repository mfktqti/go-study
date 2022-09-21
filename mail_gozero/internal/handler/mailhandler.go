package handler

import (
	"net/http"

	"go-study/mail_gozero/internal/logic"
	"go-study/mail_gozero/internal/svc"
	"go-study/mail_gozero/internal/types"
	"go-study/mail_gozero/internal/util"

	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RecMailReqeust
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		if err := util.Validate.Struct(req); err != nil {
			httpx.SetErrorHandler(func(err error) (int, interface{}) {
				errs := err.(validator.ValidationErrors)
				return http.StatusBadRequest, util.RemoveTopStruct(errs.Translate(util.Trans))
			})
			httpx.Error(w, err)
			return
		}

		l := logic.NewMailLogic(r.Context(), svcCtx)
		resp, err := l.Mail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
