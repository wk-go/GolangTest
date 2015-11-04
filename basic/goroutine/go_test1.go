//goroutine test
package main
import(
	"fmt"
	"time"
)
func say(val string){
	for i :=0; i<5;i++{
		time.Sleep(time.Microsecond *100)
		fmt.Println(val)
	}
}
func main(){
	go say("world")
	say("hello")
}