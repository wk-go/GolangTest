package main
import(
	"fmt"
	"time"
)
func main() {
	now := time.Now()
	fmt.Println("time.Now():", now)
	datesum := now.YearDay()%256
	fmt.Println("datesum:", datesum)
}
