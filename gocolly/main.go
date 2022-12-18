package main

import (
	"fmt"
	"log"
	"time"

	colly "github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
)

func main() {
	goColly()
	//goColly1()
}

func goColly() {

	// create a new collector
	c := colly.NewCollector(
		//colly.Async(true),
		colly.Debugger(&debug.LogDebugger{}),
	)

	c.OnError(func(r *colly.Response, err error) {
		log.Println("OnError", r.StatusCode, err)
	})

	c.OnResponse(func(r *colly.Response) {
		//WriteToLog(string(r.Body))
		r.Save(time.Now().Format("20060102150405") + ".html")
		log.Println("response received", r.StatusCode)

		c.OnResponse(func(r *colly.Response) {
			r.Save("rewards" + time.Now().Format("20060102150405") + ".json")
		})

		//c.Visit(fmt.Sprintf("https://secure.accor.com/web/user/v2/user?t=%d&appId=all.accor&lang=en", time.Now().UnixMicro()))

	})

	c.OnHTML("body", func(h *colly.HTMLElement) {
		//fmt.Printf("h: %#v\n", h)
		h.Response.Save("body" + time.Now().Format("20060102150405") + ".html")
	})

	c.OnScraped(func(r *colly.Response) {
		//r.Save(time.Now().Format("scraped20060102150405") + ".txt")

		log.Println("response received", r.StatusCode)
	})

	// start scraping
	//
	c.Visit("https://all.accor.com/usa/index.en.shtml")

	err := c.Post("https://login.accor.com/as/fkYDf/resume/as/authorization.ping?persistent=yes",
		map[string]string{
			"pf.username":             "adeshaye@hotmail.com",
			"pf.pass":                 "pm1maki",
			"rememberUsername":        "on",
			"pf.ok":                   "clicked",
			"pf.cancel":               "",
			"pf.alternateAuthnSystem": "",
			"pf.adapterId":            "AuthenticationPagePersistent",
			"pf.rememberUsername":     "",
		})
	if err != nil {
		log.Fatal(err)
	}

	c.Wait()
}

func goColly1() {

	mUrl := "http://www.ifeng.com/"
	//colly的主体是Collector对象，管理网络通信和负责在作业运行时执行附加的回掉函数
	c := colly.NewCollector(
		// 开启本机debug
		colly.Debugger(&debug.LogDebugger{}),
	)
	//发送请求之前的执行函数
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("这里是发送之前执行的函数")
	})
	//发送请求错误被回调
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Print(err)
	})

	//响应请求之后被回调
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response body length：", len(r.Body))
	})
	//response之后会调用该函数，分析页面数据
	c.OnHTML("div#newsList h1 a", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	//在OnHTML之后被调用
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})
	//这里是执行访问url
	c.Visit(mUrl)

}
