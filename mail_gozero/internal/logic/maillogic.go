package logic

import (
	"context"
	"fmt"
	"net/smtp"

	"go-study/mail_gozero/internal/svc"
	"go-study/mail_gozero/internal/types"

	"github.com/jordan-wright/email"
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
	// l.Logger.Infov(req)
	mailConfig := l.svcCtx.Config.MailConfig
	e := email.NewEmail()
	e.From = mailConfig.From
	e.To = []string{mailConfig.To}
	e.Subject = fmt.Sprintf("来自用户%s %s的留言,phone:%s", req.Last, req.First, req.Phone)
	content := fmt.Sprintf("mail:%s,company name:%s,detail:%s", req.Email, req.CompanyName, req.Detail)
	e.Text = []byte(content)
	err = e.Send(mailConfig.MailServerAddr, smtp.PlainAuth("", mailConfig.From, mailConfig.Code, mailConfig.MailServerHost))
	return
}
