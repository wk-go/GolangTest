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
    "strconv"
    "github.com/bitly/go-simplejson"
    "encoding/json"
    "os"
    "io"
)

func main() {
    b := NewWxBot()
    b.GetUuid()
    b.GenQrCode()
    b.Wait4Login()
    b.Login()
    b.init()
    b.status_notify()
    b.get_contact()

    fmt.Printf(":::WxBot:::%+v\n",b )
}

/********** helper func *****/
func strInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}
/********** helper func ****/

const (
    UNKONWN = "unkonwn"
    SUCCESS = "200"
    SCANED = "201"
    TIMEOUT = "408"
)

type WxUser struct {
    Uin              int
    UserName         string
    NickName         string
    HeadImgUrl       string
    ContactFlag      int
    MemberCount      int
    MemberList       []map[string]interface{}
    RemarkName       string
    HideInputBarFlag int
    Sex              int
    Signature        string
    VerifyFlag       int
    OwnerUin         int
    PYInitial        string
    PYQuanPin        string
    RemarkPYInitial  string
    RemarkPYQuanPin  string
    StarFriend       int
    AppAccountFlag   int
    Statues          int
    AttrStatus       int
    Province         string
    City             string
    Alias            string
    SnsFlag          int
    UniFriend        int
    DisplayName      string
    ChatRoomId       int
    KeyWord          string
    EncryChatRoomId  string
    IsOwner          int
}

type WxMe struct {
    Uin               int
    UserName          string
    NickName          string
    HeadImgUrl        string
    RemarkName        string
    PYInitial         string
    PYQuanPin         string
    RemarkPYInitial   string
    RemarkPYQuanPin   string
    HideInputBarFlag  int
    StarFriend        int
    Sex               int
    Signature         string
    AppAccountFlag    int
    VerifyFlag        int
    ContactFlag       int
    WebWxPluginSwitch int
    HeadImgFlag       int
    SnsFlag           int
}
type WxSyncKey struct{
    Count int
    List []WxSyncKeyItem
}
type WxSyncKeyItem struct {
    Key int
    Val int
}
type WxBot struct {
    base_uri     string
    Client       *http.Client
    uuid         string
    LoginQr      []byte
    redirect_uri string
    base_host    string
    base_request map[string]string
    skey string
    uin string
    sid string
    pass_ticket string
    device_id string
    sync_key *WxSyncKey
    my_account *WxMe
    sync_key_str string
    is_big_contact bool
    temp_pwd string
    DEBUG bool

    member_list  []*WxUser
    contact_list []*WxUser
    public_list []*WxUser
    special_list []*WxUser
    group_list []*WxUser
}
func NewWxBot() *WxBot{
    wxBot := &WxBot{}
    wxBot.device_id = "e" + strconv.Itoa(10000000000000000+rand.Intn(9999999999999999))[2:17]
    wxBot.temp_pwd = "./tmp"
    wxBot.is_big_contact = false
    wxBot.DEBUG = true
    wxBot.Client = &http.Client{}
    return wxBot
}

func (this *WxBot) Get(urlStr string) ([]byte, error) {
    return this.Do("GET",urlStr, nil,"")
}
func (this *WxBot) Post(urlStr, bodyType, body string)([]byte, error){
    return this.Do("POST",urlStr, strings.NewReader(body),bodyType)
}
func (this *WxBot) PostForm(url string, data url.Values) ([]byte, error) {
    return this.Post(url, "application/x-www-form-urlencoded", data.Encode())
}
func (this *WxBot) Do(method, urlStr string, body io.Reader, bodyType string) ([]byte,error){
    req, err := http.NewRequest(method, urlStr, body)
    if err != nil {
        return nil, err
    }
    if method == "POST"{
        if len(bodyType) == 0{
            bodyType = "application/x-www-form-urlencoded"
        }
        req.Header.Set("Content-Type", bodyType)
    }


    req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux i686; U;) Gecko/20070322 Kazehakase/0.4.5")

    resp, err := this.Client.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil,err
    }
    return data,err
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
    name :=""
    for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
        switch token := t.(type) {
        // 处理元素开始（标签）
        case xml.StartElement:
            name = token.Name.Local
            /*fmt.Printf("Token name: %s\n", name)
            for _, attr := range token.Attr {
                attrName := attr.Name.Local
                attrValue := attr.Value
                fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
            }*/
        // 处理元素结束（标签）
        case xml.EndElement:
            //fmt.Printf("Token of '%s' end\n", token.Name.Local)
            name = ""
        // 处理字符数据（这里就是元素的文本）
        case xml.CharData:
            content := string([]byte(token))
            switch name {
            case "skey":
                this.skey =content
            case "wxsid":
                this.sid = content
            case "wxuin":
                this.uin = content
            case "pass_ticket":
                this.pass_ticket = content
            }
            //fmt.Printf("This is the content: %v\n", content)
        default:
        // ...
        }
    }
    if len(this.skey) == 0 || len(this.sid) == 0 || len(this.uin) == 0 || len(this.pass_ticket) == 0{
        return false
    }
    this.base_request = map[string]string{
        "Uin": this.uin,
        "Sid": this.sid,
        "Skey": this.skey,
        "DeviceID": this.device_id,
    }
    return true
}

func (this *WxBot) init() bool{
    urlStr := this.base_uri + fmt.Sprintf("/webwxinit?r=%d&lang=en_US&pass_ticket=%s" , time.Now().Unix(), this.pass_ticket)
    j := simplejson.New()
    j.Set("BaseRequest",this.base_request)
    data,_ := j.MarshalJSON()
    fmt.Println(":::data:::", string(data))
    r,_ := this.Client.Post(urlStr, "raw", strings.NewReader(string(data)))
    defer r.Body.Close()
    body,_ := ioutil.ReadAll(r.Body)
    //fmt.Println(":::body:::",string(body))
    j,_ = simplejson.NewJson(body)
    sync_key,_ := j.Get("SyncKey").Map()
    //fmt.Println(":::sync_key:::", sync_key)
    b,_ :=json.Marshal(sync_key)
    this.sync_key = &WxSyncKey{}
    json.Unmarshal(b,this.sync_key)
    //fmt.Printf(":::this.sync_key:::%+v\n",this.sync_key)

    my_account  := j.Get("User").Interface()
    fmt.Println(":::my_account:::", my_account)
    b,_ = json.Marshal(my_account)
    //fmt.Println(":::b:::", string(b))
    this.my_account = &WxMe{}
    json.Unmarshal(b,this.my_account)
    //fmt.Printf(":::this.my_account:::%+v\n", this.my_account)

    conj := "";
    this.sync_key_str = ""
    for _,v := range this.sync_key.List{
            this.sync_key_str += conj + strconv.Itoa(v.Key) + "_" + strconv.Itoa(v.Val)
            if len(conj) == 0{
                conj ="|"
            }
    }
    ret,_ :=j.GetPath("BaseResponse/Ret").Int()
    return ret == 0
}

func (this *WxBot) status_notify() bool {
    urlStr := this.base_uri + fmt.Sprintf("/webwxstatusnotify?lang=zh_CN&pass_ticket=%s" , this.pass_ticket)
    //this.base_request["Uin"] = int(this.base_request["Uin"])
    paramsJson := simplejson.New()
    paramsJson.Set("BaseRequest",this.base_request)
    paramsJson.Set("Code",3)
    paramsJson.Set("FromUserName", this.my_account.UserName)
    paramsJson.Set("ToUserName", this.my_account.UserName)
    paramsJson.Set("ClientMsgId", int(time.Now().Unix()))
    params,_ := paramsJson.MarshalJSON()
    r,_ := this.Client.Post(urlStr,"raw", strings.NewReader(string(params)))
    defer r.Body.Close()
    body,_ := ioutil.ReadAll(r.Body)
    //fmt.Println(":::status_notify body:::",string(body))
    respJson,_ := simplejson.NewJson(body)
    ret,_ := respJson.GetPath("BaseResponse/Ret").Int()
    return ret == 0
}
func (this *WxBot) get_contact() bool{
    //获取当前账户的所有相关账号(包括联系人、公众号、群聊、特殊账号)
    fmt.Println(":::get_contact::: start")
    if this.is_big_contact {
        return false
    }
    urlStr := this.base_uri + fmt.Sprintf("/webwxgetcontact?pass_ticket=%s&skey=%s&r=%d" , this.pass_ticket, this.skey, int(time.Now().Unix()))
    fmt.Println(":::get_contact::: createUrl:",urlStr)

    //如果通讯录联系人过多，这里会直接获取失败
    fmt.Println(":::get_contact::: Geting")
    body,err := this.Post(urlStr,"raw","{}")
    fmt.Println(":::get_contact::: Get")
    if err != nil{
        fmt.Println(":::get_contact error:::", err)
        this.is_big_contact = true
        return false
    }
    fmt.Printf(":::body:::%+v\n",string(body))
    if this.DEBUG {
        if f,err := os.OpenFile(this.temp_pwd+"/contacts.json",os.O_RDWR,os.ModePerm); err == nil{
            defer f.Close()
            f.Write(body)
        }else{
            fmt.Println(":::file err:::", err)
        }
    }

    j,_ := simplejson.NewJson(body)
    member_list := j.Get("MemberList").Interface()
    fmt.Printf(":::this.member_list:::%+v\n",this.member_list)
    b,_ := json.Marshal(member_list)
    this.member_list = []*WxUser{}
    json.Unmarshal(b, this.member_list)
    fmt.Printf(":::this.member_list:::%+v\n",this.member_list)

    /*    special_users := []string{"newsapp", "fmessage", "filehelper", "weibo", "qqmail",
            "fmessage", "tmessage", "qmessage", "qqsync", "floatbottle",
            "lbsapp", "shakeapp", "medianote", "qqfriend", "readerapp",
            "blogapp", "facebookapp", "masssendapp", "meishiapp",
            "feedsapp", "voip", "blogappweixin", "weixin", "brandsessionholder",
            "weixinreminder", "wxid_novlwrv3lqwv11", "gh_22b87fa7cb3c",
            "officialaccounts", "notification_messages", "wxid_novlwrv3lqwv11",
            "gh_22b87fa7cb3c", "wxitil", "userexperience_alarm", "notification_messages"}

        this.contact_list = []*WxUser{}
        this.public_list = []*WxUser{}
        this.special_list = []*WxUser{}
        this.group_list = []*WxUser{}

        for contact := range this.member_list {
            switch contact {
            case contact.VerifyFlag & 8 != 0:  // 公众号
                this.public_list = append(this.public_list, contact)
            *//*this.account_info["normal_member"][contact["UserName"]] =
            {
                "type": "public", "info": contact,
            }*//*
            case strInSlice(contact.UserName, special_users):// 特殊账户
                this.special_list = append(this.special_list, contact)
            *//*this.account_info["normal_member"][contact["UserName"]] =
            {
                "type": "special", "info": contact
            }*//*
            case contact["UserName"].find("@@") != -1:  // 群聊
                this.group_list = append(this.group_list, contact)
            *//*this.account_info["normal_member"][contact["UserName"]] =
            {
                "type": "group", "info": contact
            }*//*
            case contact["UserName"] == this.my_account.UserName:  // 自己
            *//*this.account_info["normal_member"][contact["UserName"]] =
            {
                "type": "self", "info": contact
            }*//*
            default:
                this.contact_list = append(this.contact_list, contact)
            //this.account_info["normal_member"][contact["UserName"]] = {"type": "contact", "info": contact}
            }
        }*/

/*    this.batch_get_group_members()

    for group in this.group_members:
        for member in this.group_members[group]:
            if member["UserName"] not in this.account_info:
                this.account_info["group_member"][member["UserName"]] = \
                {"type": "group_member", "info": member, "group": group}

    if this.DEBUG{
    with open(os.path.join(this.temp_pwd, "contact_list.json"), "w") as f:
    f.write(json.dumps(this.contact_list))
    with open(os.path.join(this.temp_pwd, "special_list.json"), "w") as f:
    f.write(json.dumps(this.special_list))
    with open(os.path.join(this.temp_pwd, "group_list.json"), "w") as f:
    f.write(json.dumps(this.group_list))
    with open(os.path.join(this.temp_pwd, "public_list.json"), "w") as f:
    f.write(json.dumps(this.public_list))
    with open(os.path.join(this.temp_pwd, "member_list.json"), "w") as f:
    f.write(json.dumps(this.member_list))
    with open(os.path.join(this.temp_pwd, "group_users.json"), "w") as f:
    f.write(json.dumps(this.group_members))
    with open(os.path.join(this.temp_pwd, "account_info.json"), "w") as f:
    f.write(json.dumps(this.account_info))
}*/
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

