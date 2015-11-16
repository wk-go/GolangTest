//websocket client1
package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

func (self String) ServeHTTP(
w http.ResponseWriter,
r *http.Request) {
	fmt.Fprint(w, self)
}

type Struct struct {
	page string
}

func (self *Struct) ServeHTTP(
w http.ResponseWriter,
r *http.Request) {
	fmt.Fprint(w, self.page)
}

func main() {

	http.Handle("/", &Struct{page:ws_page})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
var ws_page = `<html>
<head></head>
<body>
<script type="text/javascript">
var sock = null;
var wsuri = "ws://127.0.0.1:1234";

window.onload = function() {

console.log("onload");

sock = new WebSocket(wsuri);

sock.onopen = function() {
console.log("connected to " + wsuri);
}

sock.onclose = function(e) {
console.log("connection closed (" + e.code + ")");
}

sock.onmessage = function(e) {
console.log("message received: " + e.data);
}
};

function send() {
var msg = document.getElementById('message').value;
sock.send(msg);
};
</script>
<h1>WebSocket Echo Test</h1>
<form>
<p>
Message: <input id="message" type="text" value="Hello, world!">
</p>
</form>
<button onclick="send();">Send Message</button>
</body>
</html>`