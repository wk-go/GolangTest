package main
import(
	"fmt"
	"time"
)
func main() {
	now := time.Now()
	fmt.Println("-------local----------")
	fmt.Println("time.Location():", now.Location())
	fmt.Println("time.Now():", now)
	datesum := now.YearDay()%256
	fmt.Println("datesum:", datesum)
	fmt.Println("-------utc----------")
	nowUTC := now.UTC()
	fmt.Println("time.Now():", nowUTC.Location())
	fmt.Println("time.Now():", nowUTC)
	datesum = now.YearDay()%256
	fmt.Println("datesum:", datesum)
}
