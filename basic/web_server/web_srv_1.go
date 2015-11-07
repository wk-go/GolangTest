//web server test 1
package main

import(
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request)  {
	fmt.Fprint(w, "Hello！！！!")
}

func main(){
	var h Hello
	if err := http.ListenAndServe("localhost:4000", h); err != nil{
		log.Fatal(err)
	}
}