package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/xuri/excelize/v2"
)

func main() {
	// 禁用chrome headless
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	f, err := excelize.OpenFile("config.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	var status, points string
	username := "adeshaye@hotmail.com"
	pass := "pm1maki"
	tasks := []chromedp.Action{
		chromedp.Navigate(`https://all.accor.com/usa/index.en.shtml`),
		chromedp.DoubleClick(`#onetrust-close-btn-container > button`, chromedp.NodeVisible),
		chromedp.Sleep(2 * time.Second), chromedp.WaitVisible(`#link-navigation-primaryHeader > div > div.link-navigation__connectZone > div > div > div > div`),
		chromedp.Click(`#link-navigation-primaryHeader > div > div.link-navigation__connectZone > div > div > div > div`, chromedp.NodeVisible),
		chromedp.Sleep(2 * time.Second),
		chromedp.Click(`#list-items > div:nth-child(5) > div > a`, chromedp.NodeVisible),
		chromedp.Sleep(2 * time.Second),
		chromedp.WaitVisible(`#primary_button`),
		chromedp.SetValue("#username-id", username),
		chromedp.SetValue("#password-id", pass),
		chromedp.Sleep(1 * time.Second),
		chromedp.Click(`#primary_button`, chromedp.NodeVisible),
		chromedp.Sleep(3 * time.Second),
		chromedp.WaitVisible(`#link-navigation-primaryHeader > div > div.link-navigation__connectZone > div > div > div > div`),
		chromedp.Sleep(5 * time.Second),
		chromedp.Click(`#link-navigation-primaryHeader > div > div.link-navigation__connectZone > div > div > div > div`, chromedp.NodeVisible),
		chromedp.Sleep(1 * time.Second),
		chromedp.InnerHTML(`#list-items > div:nth-child(3) > div > div > a > div.item__wrapper.item__wrapper--text__wrapper > span.value`, &status),
		chromedp.InnerHTML(`#list-items > div:nth-child(4) > div > div > a > div.item__wrapper.item__wrapper--text__wrapper > span.value`, &points),
		chromedp.Sleep(time.Second),
		chromedp.Click(`#logout-button`, chromedp.NodeVisible),
		chromedp.Sleep(3 * time.Second),
	}
	for i := 0; i < len(rows); i++ {
		cells := rows[i]
		username := cells[0]
		pass := cells[1]
		fmt.Printf("username: %v\n", username)
		fmt.Printf("pass: %v\n", pass)
		if i == 1 {
			tasks = append(tasks[:1], tasks[2:]...)
		}

		if i == 0 {
			tasks[8] = chromedp.SetValue("#username-id", (username))
			tasks[9] = chromedp.SetValue("#password-id", pass)
		} else {
			fmt.Printf("tasks[7]: %#v\n", tasks[7])
			tasks[7] = chromedp.SetValue("#username-id", username)
			tasks[8] = chromedp.SetValue("#password-id", pass)
		}

		err := chromedp.Run(ctx, tasks...)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		} else {
			log.Printf("Go's time.After example:\n%s,%s,%s,%s", username, pass, status, points)
		}

	}

}
