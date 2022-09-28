package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	MailConfig struct {
		From           string // 从该邮件发出
		To             string // 收邮件
		Code           string // 发邮件的安全码
		MailServerAddr string
		MailServerHost string
		Cors           bool
	}
}
