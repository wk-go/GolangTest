package main

/*
测试示例来源:https://www.cnblogs.com/pu369/p/12330074.html
*/
import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	cdpruntime "github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
	"github.com/chromedp/chromedp/kb"
)

func main() {
	log.Printf("example:")
	cmd := "ExampleListenTarget_acceptAlert"
	args := os.Args
	if len(args) >= 2 {
		cmd = os.Args[1]
	}
	switch cmd {
	case "ExampleListenTarget_acceptAlert":
		ExampleListenTarget_acceptAlert()
	case "ByJSPath":
		ByJSPath()
	case "Emulate":
		Emulate()
	case "ConsoleLog":
		ConsoleLog()
	case "ManyTabs":
		ManyTabs()
	case "Title":
		Title()
	case "WaitNewTarget":
		WaitNewTarget()
	case "DocumentDump":
		DocumentDump()
	case "RetrieveHTML":
		RetrieveHTML()
	case "Download1":
		Download1()
	case "Download2":
		Download2()
	case "Testiframe":
		Testiframe()
	case "Click":
		Click()
	case "Cookie":
		Cookie()
	case "Evaluate":
		Evaluate()
	case "Headers":
		Headers()
	case "Keys":
		Keys()
	case "Logic":
		Logic()
	case "Remote":
		Remote()
	}
}

//监听并自动关闭弹出的alert对话框。其中也包括了ExecAllocator的用法
func ExampleListenTarget_acceptAlert() {
	//内置http测试服务器，用于在网页上显示alert按钮
	ts := httptest.NewServer(writeHTML(`
<input id='alert' type='button' value='alert' onclick='alert("alert text");'/>please等5秒后，自动点击Alert,并自动关闭alert对话框。
    `))
	defer ts.Close()
	//fmt.Println(ts.URL)
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

	//chromedp监听网页上弹出alert对话框
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if ev, ok := ev.(*page.EventJavascriptDialogOpening); ok {
			fmt.Println("closing alert:", ev.Message)
			go func() {
				//自动关闭alert对话框
				if err := chromedp.Run(ctx,
					//注释掉下一行可以更清楚地看到效果
					page.HandleJavaScriptDialog(true),
				); err != nil {
					panic(err)
				}
			}()
		}
	})

	if err := chromedp.Run(ctx,
		chromedp.Navigate(ts.URL),
		chromedp.Sleep(5*time.Second),
		//自动点击页面上的alert按钮，弹出alert对话框
		chromedp.Click("#alert", chromedp.ByID),
	); err != nil {
		panic(err)
	}

	// Output:
	// closing alert: alert text
}

//向内置http服务器页面输出内容
func writeHTML(content string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//在这里设置utf-8,避免乱码
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(w, strings.TrimSpace(content))
	})
}

//用JSPath获取网页元素
func ByJSPath() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ts := httptest.NewServer(writeHTML(`
<body>
    <div id="content">cool content</div>
</body>
    `))
	defer ts.Close()

	var ids []cdp.NodeID
	var html string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(ts.URL),
		//可以在Chrome devTools中，在想要选择的元素上，用鼠标右键>Copy>Copy JS path 得到JSPath
		chromedp.NodeIDs(`document.querySelector("body > div#content")`, &ids, chromedp.ByJSPath),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			html, err = dom.GetOuterHTML().WithNodeID(ids[0]).Do(ctx)
			return err
		}),
	); err != nil {
		panic(err)
	}

	fmt.Println("Outer HTML:")
	fmt.Println(html)
}

//模拟不同的设备访问网站。其中也包括了ExecAllocator的用法
func Emulate() {
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

	var buf []byte
	if err := chromedp.Run(ctx,
		chromedp.Emulate(device.IPhone7),
		chromedp.Navigate(`https://baidu.com/`),
		//chromedp.WaitVisible(`#kw`, chromedp.ByID),
		chromedp.Sleep(3*time.Second),
		chromedp.SendKeys(`input[name=word]`, "what's my user agent?\n"),
		//chromedp.WaitVisible(`#lg`, chromedp.ByID),
		chromedp.Sleep(3*time.Second),
		chromedp.CaptureScreenshot(&buf),
	); err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile("baidu-iphone7.png", buf, 0644); err != nil {
		panic(err)
	}
}

//获取chrome控制台的内容
func ConsoleLog() {
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
	ts := httptest.NewServer(writeHTML(`
<body>
<script>
    console.log("hello js world")
    console.warn("warn hello js world")    
    var p = document.createElement("div");
    p.setAttribute("id", "jsFinished");
    document.body.appendChild(p);
</script>
</body>
    `))
	defer ts.Close()

	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *cdpruntime.EventConsoleAPICalled:
			fmt.Printf("console.%s call:\n", ev.Type)
			for _, arg := range ev.Args {
				fmt.Printf("%s - %s\n", arg.Type, arg.Value)
			}

		}

	})

	if err := chromedp.Run(ctx,
		chromedp.Navigate(ts.URL),
		chromedp.Sleep(15*time.Second), //在打开的chrome窗口上有提示：chrome正受到自动测试软件的控制
		chromedp.WaitVisible("#jsFinished", chromedp.ByID),
		chromedp.Sleep(15*time.Second),
	); err != nil {
		panic(err)
	}
}

//控制多个chrome tab页
func ManyTabs() {
	// new browser, first tab
	ctx1, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// ensure the first tab is created
	if err := chromedp.Run(ctx1); err != nil {
		panic(err)
	}

	// same browser, second tab
	ctx2, _ := chromedp.NewContext(ctx1)

	// ensure the second tab is created
	if err := chromedp.Run(ctx2); err != nil {
		panic(err)
	}

	c1 := chromedp.FromContext(ctx1)
	c2 := chromedp.FromContext(ctx2)

	fmt.Printf("Same browser: %t\n", c1.Browser == c2.Browser)
	fmt.Printf("Same tab: %t\n", c1.Target == c2.Target)
}
func Title() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ts := httptest.NewServer(writeHTML(`
<head>
    <title>fancy website title</title>
</head>
<body>
    <div id="content"></div>
</body>
    `))
	defer ts.Close()

	var title string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(ts.URL),
		chromedp.Title(&title),
	); err != nil {
		panic(err)
	}
	fmt.Println(title)
}

//等待chrome打开新的tab页，
func WaitNewTarget() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	mux := http.NewServeMux()
	mux.Handle("/first", writeHTML(`
<input id='newtab' type='button' value='open' onclick='window.open("/second", "_blank");'/>
    `))
	mux.Handle("/second", writeHTML(``))
	ts := httptest.NewServer(mux)
	defer ts.Close()

	// Grab the first spawned tab that isn't blank.
	ch := chromedp.WaitNewTarget(ctx, func(info *target.Info) bool {
		return info.URL != ""
	})
	if err := chromedp.Run(ctx,
		chromedp.Navigate(ts.URL+"/first"),
		chromedp.Click("#newtab", chromedp.ByID),
	); err != nil {
		panic(err)
	}
	newCtx, cancel := chromedp.NewContext(ctx, chromedp.WithTargetID(<-ch))
	defer cancel()

	var urlstr string
	if err := chromedp.Run(newCtx, chromedp.Navigate("https://www.baidu.com"), chromedp.Location(&urlstr)); err != nil {
		panic(err)
	}
	fmt.Println("new tab's path:", strings.TrimPrefix(urlstr, ts.URL))
}

//向页面注入javascript并执行，获取执行后的页面DOM结构。
//核心是使用runtime.Evaluate
func DocumentDump() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ts := httptest.NewServer(writeHTML(`<!doctype html>
<html>
<body>
  <div id="content">the content</div>
<iframe src="https://www.baidu.com"></iframe>
</body>
</html>`))
	defer ts.Close()

	const expr = `(function(d, id, v) {
        var b = d.querySelector('body');
        var el = d.createElement('div');
        el.id = id;
        el.innerText = v;
        b.insertBefore(el, b.childNodes[0]);
    })(document, %q, %q);`

	var nodes []*cdp.Node
	if err := chromedp.Run(ctx,
		chromedp.Navigate(ts.URL),
		chromedp.Nodes(`document`, &nodes, chromedp.ByJSPath),
		chromedp.WaitVisible(`#content`),
		chromedp.ActionFunc(func(ctx context.Context) error {
			s := fmt.Sprintf(expr, "thing", "a new thing!")
			_, exp, err := cdpruntime.Evaluate(s).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.WaitVisible(`#thing`),
	); err != nil {
		panic(err)
	}

	fmt.Println("Document tree:")
	fmt.Print(nodes[0].Dump("  ", "  ", false))
}

//获取网页html源码
//核心是使用OuterHTML
func RetrieveHTML() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ts := httptest.NewServer(writeHTML(`
<body>
<p id="content" onclick="changeText()">Original content.</p>
<iframe src="https://www.baidu.com"></iframe>
<script>
function changeText() {
    document.getElementById("content").textContent = "New content!"
}
</script>
</body>
    `))
	defer ts.Close()

	var outerBefore, outerAfter string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(ts.URL),
		chromedp.OuterHTML("#content", &outerBefore),
		chromedp.Click("#content", chromedp.ByID),
		chromedp.OuterHTML("#content", &outerAfter),
	); err != nil {
		panic(err)
	}
	fmt.Println("OuterHTML before clicking:")
	fmt.Println(outerBefore)
	fmt.Println("OuterHTML after clicking:")
	fmt.Println(outerAfter)
}

//download1的测试
func Download1() {
	//开启静态文件服务器
	http.Handle("/", http.FileServer(http.Dir("E:/goapp/src/chromedpnewExample/")))
	http.Handle("/index", writeHTML(`<!doctype html>
<html>
<body>  
 <a  id="down" href="123.zip">download</a>
</body>
</html>`))
	go http.ListenAndServe(":9090", nil)
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
	//下载目录
	wd, _ := os.Getwd()
	wd += "\\d"
	fmt.Println(wd)
	if err := chromedp.Run(ctx,
		page.SetDownloadBehavior(page.SetDownloadBehaviorBehaviorAllow).WithDownloadPath(wd),
		chromedp.Navigate("http://www.mersenne.org/ftp_root/gimps/p95v287.MacOSX.noGUI.tar.gz"),
		chromedp.Sleep(3*time.Second), //这里不知道如何等待下载结束？如果chrome设置成headless为true的时候下载不正常?
	); err != nil {
		panic(err)
	}
	fmt.Println("down load ok")

}

//download2的测试
//123.zip为任意压缩包，存放于主程序目录下；并在主程序所在目录中新建子目录d
func Download2() {
	//开启静态文件服务器
	http.Handle("/", http.FileServer(http.Dir("E:/goapp/src/chromedpnewExample/")))
	http.Handle("/index", writeHTML(`<!doctype html>
<html>
<body>  
 <a  id="down" href="123.zip">download</a>
</body>
</html>`))
	go http.ListenAndServe(":9090", nil)
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
	//下载目录
	wd, _ := os.Getwd()
	wd += "\\d"
	fmt.Println(wd)
	if err := chromedp.Run(ctx,
		page.SetDownloadBehavior(page.SetDownloadBehaviorBehaviorAllow).WithDownloadPath(wd),
		chromedp.Navigate("http://localhost:9090/index"),
		chromedp.WaitVisible("#down"),
		chromedp.Click("#down", chromedp.ByID),
		chromedp.Sleep(3*time.Second), //这里不知道如何等待下载结束？如果chrome设置成headless为true的时候下载不正常?
	); err != nil {
		panic(err)
	}
	fmt.Println("down load ok")

}

//iframe的测试

/*123.html
<html>
<body>
<a id="clickme" href="https://www.google.com>clickme!</a>
</body>
</htm>
*/
func Testiframe() {
	//开启静态文件服务器
	wd, _ := os.Getwd()
	http.Handle("/", http.FileServer(http.Dir(wd)))
	http.Handle("/index", writeHTML(`<!doctype html>
<html>
<body>
  <div id="content" name="content name">the content</div> 
<iframe src="123.html"></iframe>
</body>
</html>`))
	go http.ListenAndServe(":9090", nil)
	framefile123 := `
<html><body>
<a id="clickme" href="https://www.google.com">clickme!</a>
</body></html>
    `
	f123 := []byte(framefile123)
	if err := ioutil.WriteFile("./123.html", f123, 0666); err != nil {
		panic(err)
	}

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
	var outer, url, frameurl string
	var ok bool
	if err := chromedp.Run(ctx,
		chromedp.Navigate("http://localhost:9090/index"),
		chromedp.Sleep(2*time.Second),
		chromedp.WaitVisible("iframe", chromedp.ByQuery),
		chromedp.AttributeValue("iframe", "src", &outer, &ok),
		chromedp.Location(&url),
	); err != nil {
		panic(err)
	}
	fmt.Println("Outer:", outer, " url:", url)
	frameurl = url[0:strings.LastIndex(url, "/")] + "/" + outer
	fmt.Println("frameurl", frameurl)
	if err := chromedp.Run(ctx,
		chromedp.Navigate(frameurl),
		chromedp.Sleep(2*time.Second),
		chromedp.WaitVisible("#clickme"),
		chromedp.Click("#clickme", chromedp.ByID),
		chromedp.Sleep(2*time.Second),
	); err != nil {
		panic(err)
	}
}

//以下例子来自https://github.com/chromedp/examples
//单击元素,获取textarea的value
func Click() {
	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
	var example string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://golang.org/pkg/time/`),
		// wait for footer element is visible (ie, page is loaded)
		chromedp.WaitVisible(`body > footer`),
		// find and click "Expand All" link
		chromedp.Click(`#pkg-examples > div`, chromedp.NodeVisible),
		// retrieve the value of the textarea
		chromedp.Value(`#example_After .play .input textarea`, &example),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Go's time.After example:\n%s", example)
}

//设置Cookie
func Cookie() {
	var (
		flagPort = flag.Int("port", 8544, "port")
	)
	flag.Parse()

	// start cookie server
	go cookieServer(fmt.Sprintf(":%d", *flagPort))

	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res string
	err := chromedp.Run(ctx, setcookies(
		fmt.Sprintf("http://localhost:%d", *flagPort), &res,
		"cookie1", "value1",
		"cookie2", "value2",
	))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("main: chrome received cookies: %s", res)
}

// cookieServer creates a simple HTTP server that logs any passed cookies.
func cookieServer(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		cookies := req.Cookies()
		for i, cookie := range cookies {
			log.Printf("server:  from %s, server received cookie %d: %v", req.RemoteAddr, i, cookie)
		}
		buf, err := json.MarshalIndent(req.Cookies(), "", "  ")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(res, indexHTML, string(buf))
	})
	return http.ListenAndServe(addr, mux)
}

// setcookies returns a task to navigate to a host with the passed cookies set
// on the network request.
func setcookies(host string, res *string, cookies ...string) chromedp.Tasks {
	if len(cookies)%2 != 0 {
		panic("length of cookies must be divisible by 2")
	}
	return chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			// create cookie expiration
			expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
			// add cookies to chrome
			for i := 0; i < len(cookies); i += 2 {
				success, err := network.SetCookie(cookies[i], cookies[i+1]).
					WithExpires(&expr).
					WithDomain("localhost").
					WithHTTPOnly(true).
					Do(ctx)
				if err != nil {
					return err
				}
				if !success {
					return fmt.Errorf("could not set cookie %q to %q", cookies[i], cookies[i+1])
				}
			}
			return nil
		}),
		// navigate to site
		chromedp.Navigate(host),
		// read the returned values
		chromedp.Text(`#result`, res, chromedp.ByID, chromedp.NodeVisible),
		// read network values
		chromedp.ActionFunc(func(ctx context.Context) error {
			cookies, err := network.GetAllCookies().Do(ctx)
			if err != nil {
				return err
			}

			for i, cookie := range cookies {
				log.Printf("chromedp:  chrome cookie %d: %+v", i, cookie)
			}

			return nil
		}),
	}
}

const (
	indexHTML = `<!doctype html>
<html>
<body>
  <div id="result">%s</div>
</body>
</html>`
)

//执行javascript,显示window对象的keys
func Evaluate() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res []string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.baidu.com/`),
		chromedp.WaitVisible(`#wrapper`, chromedp.ByID),
		chromedp.Evaluate(`Object.keys(window);`, &res),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("window object keys: %v", res)
}

func Headers() {
	var (
		flagPort = flag.Int("port", 8544, "port")
	)
	flag.Parse()

	// run server
	go headerServer(fmt.Sprintf(":%d", *flagPort))

	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var res string
	err := chromedp.Run(ctx, setheaders(
		fmt.Sprintf("http://localhost:%d", *flagPort),
		map[string]interface{}{
			"X-Header": "my request header",
		},
		&res,
	))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("received headers: %s", res)
}

// headerServer is a simple HTTP server that displays the passed headers in the html.
func headerServer(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		buf, err := json.MarshalIndent(req.Header, "", "  ")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(res, indexHTML_headers, string(buf))
	})
	return http.ListenAndServe(addr, mux)
}

// setheaders returns a task list that sets the passed headers.
func setheaders(host string, headers map[string]interface{}, res *string) chromedp.Tasks {
	headersJson, _ := json.Marshal(headers)
	return chromedp.Tasks{
		network.Enable(),
		network.SetExtraHTTPHeaders(network.Headers(headersJson)),
		chromedp.Navigate(host),
		chromedp.Text(`#result`, res, chromedp.ByID, chromedp.NodeVisible),
	}
}

const indexHTML_headers = `<!doctype html>
<html>
<body>
  <div id="result">%s</div>
</body>
</html>`

func Keys() {
	var (
		flagPort = flag.Int("port", 8544, "port")
	)
	flag.Parse()

	// run server
	go testServer(fmt.Sprintf(":%d", *flagPort))

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
	// create context
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// run task list
	var val1, val2, val3, val4 string
	err := chromedp.Run(ctx, sendkeys(fmt.Sprintf("http://localhost:%d", *flagPort), &val1, &val2, &val3, &val4))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("#input1 value: %s", val1)
	log.Printf("#textarea1 value: %s", val2)
	log.Printf("#input2 value: %s", val3)
	log.Printf("#select1 value: %s", val4)
}

// sendkeys sends
func sendkeys(host string, val1, val2, val3, val4 *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(host),
		chromedp.WaitVisible(`#input1`, chromedp.ByID),
		chromedp.WaitVisible(`#textarea1`, chromedp.ByID),
		chromedp.Sleep(3 * time.Second),
		chromedp.SendKeys(`#textarea1`, kb.End+"\b\b\n\naoeu\n\ntest1\n\nblah2\n\n\t\t\t\b\bother box!\t\ntest4", chromedp.ByID),
		chromedp.Sleep(3 * time.Second),
		chromedp.Value(`#input1`, val1, chromedp.ByID),
		chromedp.Value(`#textarea1`, val2, chromedp.ByID),
		chromedp.SetValue(`#input2`, "test3", chromedp.ByID),
		chromedp.Sleep(3 * time.Second),
		chromedp.Value(`#input2`, val3, chromedp.ByID),
		chromedp.SendKeys(`#select1`, kb.ArrowDown+kb.ArrowDown, chromedp.ByID),
		chromedp.Sleep(3 * time.Second),
		chromedp.Value(`#select1`, val4, chromedp.ByID),
	}
}

// testServer is a simple HTTP server that displays the passed headers in the html.
func testServer(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(res, indexHTML_keys)
	})
	return http.ListenAndServe(addr, mux)
}

const indexHTML_keys = `<!doctype html>
<html>
<head>
  <title>example</title>
</head>
<body>
  <div id="box1" style="display:none">
    <div id="box2">
      <p>box2</p>
    </div>
  </div>
  <div id="box3">
    <h2>box3</h3>
    <p id="box4">
      box4 text
      <input id="input1" value="some value"><br><br>
      <textarea id="textarea1" style="width:500px;height:400px">textarea</textarea><br><br>
      <input id="input2" type="submit" value="Next">
      <select id="select1">
        <option value="one">1</option>
        <option value="two">2</option>
        <option value="three">3</option>
        <option value="four">4</option>
      </select>
    </p>
  </div>
</body>
</html>`

//复杂逻辑，获取页面标题和链接
func Logic() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// list awesome go projects for the "Selenium and browser control tools."
	res, err := listAwesomeGoProjects(ctx, "Selenium and browser control tools.")
	if err != nil {
		log.Fatalf("could not list awesome go projects: %v", err)
	}

	// output the values
	for k, v := range res {
		log.Printf("project %s (%s): '%s'", k, v.URL, v.Description)
	}
}

// projectDesc contains a url, description for a project.
type projectDesc struct {
	URL, Description string
}

// listAwesomeGoProjects is the highest level logic for browsing to the
// awesome-go page, finding the specified section sect, and retrieving the
// associated projects from the page.
func listAwesomeGoProjects(ctx context.Context, sect string) (map[string]projectDesc, error) {
	// force max timeout of 15 seconds for retrieving and processing the data
	var cancel func()
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	sel := fmt.Sprintf(`case "p[text": p[text()[contains(., '%s')]]`, sect)

	// navigate
	if err := chromedp.Run(ctx, chromedp.Navigate(`https://github.com/avelino/awesome-go`)); err != nil {
		return nil, fmt.Errorf("could not navigate to github: %v", err)
	}

	// wait visible
	if err := chromedp.Run(ctx, chromedp.WaitVisible(sel)); err != nil {
		return nil, fmt.Errorf("could not get section: %v", err)
	}

	sib := sel + `/following-sibling::ul/li`

	// get project link text
	var projects []*cdp.Node
	if err := chromedp.Run(ctx, chromedp.Nodes(sib+`/child::a/text()`, &projects)); err != nil {
		return nil, fmt.Errorf("could not get projects: %v", err)
	}

	// get links and description text
	var linksAndDescriptions []*cdp.Node
	if err := chromedp.Run(ctx, chromedp.Nodes(sib+`/child::node()`, &linksAndDescriptions)); err != nil {
		return nil, fmt.Errorf("could not get links and descriptions: %v", err)
	}

	// check length
	if 2*len(projects) != len(linksAndDescriptions) {
		return nil, fmt.Errorf("projects and links and descriptions lengths do not match (2*%d != %d)", len(projects), len(linksAndDescriptions))
	}

	// process data
	res := make(map[string]projectDesc)
	for i := 0; i < len(projects); i++ {
		res[projects[i].NodeValue] = projectDesc{
			URL:         linksAndDescriptions[2*i].AttributeValue("href"),
			Description: strings.TrimPrefix(strings.TrimSpace(linksAndDescriptions[2*i+1].NodeValue), "- "),
		}
	}

	return res, nil
}

//这个没测试成功
func Remote() {
	var devToolWsUrl string
	flag.StringVar(&devToolWsUrl, "devtools-ws-url", "", "DevTools Websocket URL")
	flag.Parse()

	actxt, cancelActxt := chromedp.NewRemoteAllocator(context.Background(), devToolWsUrl)
	defer cancelActxt()

	ctxt, cancelCtxt := chromedp.NewContext(actxt) // create new tab
	defer cancelCtxt()                             // close tab afterwards

	var body string
	if err := chromedp.Run(ctxt,
		chromedp.Navigate("https://baidu.com"),
		chromedp.WaitVisible("#logo_homepage_link"),
		chromedp.OuterHTML("html", &body),
	); err != nil {
		log.Fatalf("Failed getting body of pleasegiveurl: %v", err)
	}

	log.Println("Body of  pleasegiveurl starts with:")
	log.Println(body[0:100])
}
