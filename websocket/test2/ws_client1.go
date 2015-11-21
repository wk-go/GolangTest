//websocket client1
package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

type String string

func (self String) ServeHTTP(
w http.ResponseWriter,
r *http.Request) {
	fmt.Fprint(w, self)
}

type WS_handle struct {
}

func (self *WS_handle) ServeHTTP(
w http.ResponseWriter,
r *http.Request) {
	tpl_file := "./html/ws_client1.html"
	ws_page, err := ioutil.ReadFile(tpl_file)
	if err != nil {
		fmt.Fprint(w, "Error:There is no template file", tpl_file)
		return
	}
	fmt.Fprint(w, string(ws_page))
}

func main() {
	http.Handle("/", &WS_handle{})
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./html/assets"))))
	log.Fatal(http.ListenAndServe("localhost:4000", nil))

}