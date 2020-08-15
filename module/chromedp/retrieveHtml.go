package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"net/http"
	"net/http/httptest"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ts := httptest.NewServer(writeHTML(`
<body>
<p id="content" onclick="changeText()">Original content.</p>
<script>
function changeText() {
	document.getElementById("content").textContent = "New content!"
}
</script>
</body>
	`))
	defer ts.Close()

	var outerBefore, outerAfter, document string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(ts.URL),
		chromedp.OuterHTML("#content", &outerBefore),
		chromedp.Click("#content", chromedp.ByID),
		chromedp.OuterHTML("#content", &outerAfter),
		chromedp.OuterHTML("html", &document),
	); err != nil {
		panic(err)
	}
	fmt.Println(ts.URL)
	fmt.Println("OuterHTML before clicking:")
	fmt.Println(outerBefore)
	fmt.Println("OuterHTML after clicking:")
	fmt.Println(outerAfter)

	fmt.Println("Body:")
	fmt.Println(document)
}

func writeHTML(s string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, s)
	})
}
