package main
import(
    "os"
    "fmt"
)
func main() {
    err := os.MkdirAll("./dir1/dir2/dir3", 0777)
    if err != nil{
        fmt.Println("os.MkdirAll err:", err)
    }
    err = os.Mkdir("./dirX", 0777)
    if err != nil{
        fmt.Println("os.Mkdir err:", err)
    }
    err = os.Mkdir("./dirX1/dirX2/dirX3", 0777)
    if err != nil{
        fmt.Println("os.Mkdir 2 err:", err)
    }
}
