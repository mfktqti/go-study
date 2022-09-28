package main

import (
	"flag"
	"fmt"

	"go-study/mail_gozero/internal/config"
	"go-study/mail_gozero/internal/handler"
	"go-study/mail_gozero/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/mailgozero-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	// server := rest.MustNewServer(c.RestConf, rest.WithCors())
	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
