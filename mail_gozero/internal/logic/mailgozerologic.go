package logic

import (
	"context"

	"go-study/mail_gozero/internal/svc"
	"go-study/mail_gozero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Mail_gozeroLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMail_gozeroLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Mail_gozeroLogic {
	return &Mail_gozeroLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Mail_gozeroLogic) Mail_gozero(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
