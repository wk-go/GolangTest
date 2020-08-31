package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"os"
	"time"
)

func main() {
	//增加选项，允许chrome窗口显示出来
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	//创建chrome窗口
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()
	url := "https://www.baidu.com"
	if len(os.Args) > 1 {
		url = os.Args[1]
	}

	var document string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(5*time.Second),
		chromedp.OuterHTML("html", &document),
	); err != nil {
		panic(err)
	}
	fmt.Println(url)
	fmt.Println("Body:")
	fmt.Println(document)
	f, err := os.OpenFile("./retrieveHtmlUrl.html", os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write([]byte(document))
}
