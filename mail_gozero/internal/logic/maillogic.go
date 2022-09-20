package logic

import (
	"context"

	"go-study/mail_gozero/internal/svc"
	"go-study/mail_gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailLogic {
	return &MailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailLogic) Mail(req *types.RecMailReqeust) (resp *types.RecMailResponse, err error) {
	l.Logger.Infov(req)
	return
}
