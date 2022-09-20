package main

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func main() {

	e := email.NewEmail()
	e.From = "周黎 <514588891@qq.com>"
	e.To = []string{"mfktqti_zhou@163.com"}
	e.Bcc = []string{"514588891@qq.com"}
	e.Cc = []string{"514588891@qq.com"}
	e.Subject = "测试golang email"
	e.Text = []byte("本文 测试 内容！")
	// e.HTML = []byte("<h1>html 邮件内容!</h1>")
	err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", "514588891@qq.com", "oenxsptxcafj", "smtp.qq.com"))
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("\"发送邮件成功\": %v\n", "发送邮件成功")
	}
}
