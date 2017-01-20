package main

import (
    "net/http"
    "time"
    "math/rand"
    "fmt"
    "net/url"
    "io/ioutil"
    "regexp"
    "github.com/skip2/go-qrcode"
    "strings"
    "encoding/xml"
)

func main() {
    b := &WxBot{}
    b.GetUuid()
    b.GenQrCode()
    b.Wait4Login()
    b.Login()
}

const (
    UNKONWN = "unkonwn"
    SUCCESS = "200"
    SCANED = "201"
    TIMEOUT = "408"
)

type WxBot struct {
    base_uri     string
    Client       http.Client
    uuid         string
    LoginQr      []byte
    redirect_uri string
    base_host    string
    base_request map[string]string
}
func (this *WxBot) GetUuid() bool{
    urlStr := "https://login.weixin.qq.com/jslogin"
    params := url.Values{
        "appid": []string{"wx782c26e4c19acffb"},
        "fun": []string{"new"},
        "lang": []string{"zh_CN"},
        "_": []string{string(int(time.Now().Unix()) * 1000 + rand.Int())},
    }
    resp, err :=this.Client.PostForm(urlStr,params)
    if err != nil {
        // handle error
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err !=nil{

    }
    fmt.Println(":::resp.Body:::",string(body))

    regx := "window.QRLogin.code = ([0-9]+); window.QRLogin.uuid = \"(.+?)\""
    re, err := regexp.Compile(regx)
    submatch := re.FindSubmatch(body)

    for i,v := range submatch {
        fmt.Println(":::submatch:::", i,string(v))
    }

    if len(submatch) == 3 {
        this.uuid = string(submatch[2])
        return string(submatch[1]) == "200"
    }
    return false
}

func (this *WxBot) GenQrCode(){
    urlStr := "https://login.weixin.qq.com/l/"+this.uuid
    /*var png []byte
    png, err := qrcode.Encode(urlStr, qrcode.Medium, 256)
    if err != nil{

    }
    this.LoginQr = png*/
    err := qrcode.WriteFile(urlStr, qrcode.Medium, 256, "qr.png")
    if err != nil{

    }
}
func (this *WxBot) Wait4Login() string{
    LOGIN_TEMPLATE := "https://login.weixin.qq.com/cgi-bin/mmwebwx-bin/login?tip=%d&uuid=%s&_=%v"
    tip :=1

    try_later_secs := 1*time.Second
    MAX_RETRY_TIMES := 10

    code := UNKONWN

    retry_time := MAX_RETRY_TIMES
    for ;retry_time > 0; {
        urlStr := fmt.Sprintf(LOGIN_TEMPLATE, tip, this.uuid, time.Now().Unix())
        fmt.Printf(":::url::: %s \n",urlStr)
        code, data := this.do_request(urlStr)
        switch {
            case code == SCANED :
                fmt.Printf("[INFO] Please confirm to login.\n")
                tip = 0
            case code == SUCCESS :
                // 确认登录成功
                re, err := regexp.Compile("window.redirect_uri=\"(.+?)\";")
                if err != nil {

                }
                fmt.Println(":::body:::",string(data))
                param := re.FindSubmatch(data)
                fmt.Println(":::param:::",param)

                RedirectUri := string(param[1]) + "&fun=new"
                this.redirect_uri = RedirectUri
                this.base_uri = RedirectUri[:strings.LastIndex(RedirectUri, "/")]
                temp_host := this.base_uri[8:]
                this.base_host = temp_host[:strings.LastIndex(temp_host, "/")]
                fmt.Printf("[Success] WeChat login:%v\n", this)
                return code
            case code == TIMEOUT :
                fmt.Printf("[ERROR] WeChat login timeout.retry in % s secs later...\n", try_later_secs)
                tip = 1  //重置
                retry_time -= 1
                time.Sleep(try_later_secs)
            default:
                fmt.Printf("[ERROR] WeChat login exception return_code = %s.retry in %s secs later...\n", code, try_later_secs)
                tip = 1
                retry_time -= 1
                time.Sleep(try_later_secs)
        }
    }
    return code
}

func (this *WxBot) Login() bool{
    if len(this.redirect_uri) < 4 {
        fmt.Println("[ERROR] Login failed due to network problem, please try again.")
        return false
    }

    r,err := this.Client.Get(this.redirect_uri)
    if err != nil{
        fmt.Println(":::err::",err)
        return false
    }
    defer r.Body.Close()
    body,_ := ioutil.ReadAll(r.Body)
    data := string(body)
    decoder := xml.NewDecoder(strings.NewReader(data))
    for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
        switch token := t.(type) {
        // 处理元素开始（标签）
        case xml.StartElement:
            name := token.Name.Local
            fmt.Printf("Token name: %s\n", name)
            for _, attr := range token.Attr {
                attrName := attr.Name.Local
                attrValue := attr.Value
                fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
            }
        // 处理元素结束（标签）
        case xml.EndElement:
            fmt.Printf("Token of '%s' end\n", token.Name.Local)
        // 处理字符数据（这里就是元素的文本）
        case xml.CharData:
            content := string([]byte(token))
            fmt.Printf("This is the content: %v\n", content)
        default:
        // ...
        }
    }
    /*
    doc = xml.dom.minidom.parseString(data)
    root = doc.documentElement

    for node := range root.childNodes{

    if node.nodeName == "skey":
    this.skey = node.childNodes[0].data
    elif
    node.nodeName == "
    wxsid
    ":
    this.sid = node.childNodes[0].data
    elif
    node.nodeName == "
    wxuin
    ":
    this.uin = node.childNodes[0].data
    elif
    node.nodeName == "
    pass_ticket
    ":
    this.pass_ticket = node.childNodes[0].data
}

    if "" in(this.skey, this.sid, this.uin, this.pass_ticket):
    return False

    this.base_request = {
        "
    Uin": this.uin,
    "Sid": this.sid,
    "Skey": this.skey,
    "DeviceID": this.device_id,
}
*/
return true
}

func (this *WxBot) do_request(url string) (string, []byte) {
    r, err := this.Client.Get(url)
    if err != nil {

    }
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {

    }
    fmt.Println(":::body:::",string(body))
    re, err := regexp.Compile("window.code=([0-9]+?);")
    if err != nil {

    }
    param := re.FindSubmatch(body)

    fmt.Println(":::param:::", param)
    code := string(param[1])

    return code, body
}

