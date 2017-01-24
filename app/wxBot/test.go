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
    "net/http/cookiejar"
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
    b.proc_msg()

    fmt.Printf(":::WxBot:::%+v\n",b )
    fmt.Printf("%+v",b.Client.Jar)
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
type WxBaseResponse struct {
    Ret int
    ErrMsg string
}
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
type WxGroup struct{
    Uin int
    UserName string
    NickName string
    HeadImgUrl string
    ContactFlag int
    MemberCount int
    MemberList []*WxGroupMember
    RemarkName string
    HideInputBarFlag int
    Sex int
    Signature string
    VerifyFlag int
    OwnerUin int
    PYInitial string
    PYQuanPin string
    RemarkPYInitial string
    RemarkPYQuanPin string
    StarFriend int
    AppAccountFlag int
    Statues int
    AttrStatus int
    Province string
    City string
    Alias string
    SnsFlag int
    UniFriend int
    DisplayName string
    ChatRoomId int
    KeyWord string
    EncryChatRoomId string
    IsOwner int
}

type WxGroupMember struct{
    Uin int
    UserName string
    NickName string
    AttrStatus int
    PYInitial string
    PYQuanPin string
    RemarkPYInitial string
    RemarkPYQuanPin string
    MemberStatus int
    DisplayName string
    KeyWord string
}
type WxSyncResponse struct {
    BaseResponse *WxBaseResponse
    AddMsgCount int
    AddMsgList []*WxMsg
    ModContactCount int
    ModContactList []interface{}
    DelContactCount int
    DelContactList []interface{}
    ModChatRoomMemberCount int
    ModChatRoomMemberList []interface{}
    Profile interface{}
    ContinueFlag int
    SyncKey *WxSyncKey
    SKey string
    SyncCheckKey *WxSyncKey
}
type WxMsg struct {
    MsgId string
    FromUserName string
    ToUserName string
    MsgType int
    Content string
    Status int
    ImgStatus int
    CreateTime int
    VoiceLength int
    PlayLength int
    FileName string
    FileSize string
    MediaId string
    Url string
    AppMsgType int
    StatusNotifyCode int
    StatusNotifyUserName string
RecommendInfo *WxMsgRecommendInfo
ForwardFlag int
AppInfo *WxMsgAppInfo
HasProductId int
Ticket string
ImgHeight int
ImgWidth int
SubMsgType int
NewMsgId int
OriContent string
}
type WxMsgRecommendInfo struct {
    UserName string
    NickName string
    QQNum int
    Province string
    City string
    Content string
    Signature string
    Alias string
    Scene int
    VerifyFlag int
    AttrStatus int
    Sex int
    Ticket string
    OpCode int
}

type WxMsgAppInfo struct {
    AppID string
    Type int
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

    sync_host string

    member_list  []*WxUser
    contact_list []*WxUser
    public_list []*WxUser
    special_list []*WxUser
    group_list []*WxUser

    group_members map[string][]*WxGroupMember
    encry_chat_room_id_list map[string]string
}
func NewWxBot() *WxBot{
    wxBot := &WxBot{}
    wxBot.device_id = "e" + strconv.Itoa(10000000000000000+rand.Intn(9999999999999999))[2:17]
    wxBot.temp_pwd = "./tmp"
    wxBot.is_big_contact = false
    wxBot.DEBUG = true
    cookieJar, _ := cookiejar.New(nil)
    wxBot.Client = &http.Client{
        Jar: cookieJar,
    }
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
    body, err :=this.PostForm(urlStr,params)
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
                this.base_host = temp_host[:strings.Index(temp_host, "/")]
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

    body,err := this.Get(this.redirect_uri)
    if err != nil{
        fmt.Println(":::err::",err)
        return false
    }
    data := string(body)
    fmt.Println("::::::login data:::::", data)
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
    body,_ := this.Post(urlStr, "raw", string(data))
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
    body,_ := this.Post(urlStr,"raw", string(params))
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
    fmt.Printf(":::member_list:::%+v\n",member_list)
    b,_ := json.Marshal(member_list)
    fmt.Println("::::b:::::",string(b))
    this.member_list = []*WxUser{}
    json.Unmarshal(b, &this.member_list)
    fmt.Printf(":::this.member_list:::%+v\n",this.member_list)/*
    fmt.Printf(":::this.member_list0:::%+v\n",this.member_list[0])
    fmt.Printf(":::this.member_list1:::%+v\n",this.member_list[1])
    fmt.Printf(":::this.member_list2:::%+v\n",this.member_list[2])
    fmt.Printf(":::this.member_list3:::%+v\n",this.member_list[3])*/

        special_users := []string{"newsapp", "fmessage", "filehelper", "weibo", "qqmail",
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

        for _,contact := range this.member_list {
            fmt.Printf(":::::contact:::::::%+v\n\n",contact)
            switch {
            case (contact.VerifyFlag & 8) != 0:  // 公众号
                this.public_list = append(this.public_list, contact)
            //this.account_info["normal_member"][contact["UserName"]] ={"type": "public", "info": contact,}
            case strInSlice(contact.UserName, special_users):// 特殊账户
                this.special_list = append(this.special_list, contact)
            //this.account_info["normal_member"][contact["UserName"]] ={"type": "special", "info": contact}
            case strings.Index(contact.UserName,"@@") != -1:  // 群聊
                this.group_list = append(this.group_list, contact)
            //this.account_info["normal_member"][contact["UserName"]] ={"type": "group", "info": contact}
            case this.my_account !=nil && contact.UserName == this.my_account.UserName:  // 自己
            //this.account_info["normal_member"][contact["UserName"]] ={"type": "self", "info": contact}
            default:
                this.contact_list = append(this.contact_list, contact)
            //this.account_info["normal_member"][contact["UserName"]] = {"type": "contact", "info": contact}
            }
        }

    this.batch_get_group_members()

    /*    for group in this.group_members:
            for member in this.group_members[group]:
                if member["UserName"] not in this.account_info:
                    this.account_info["group_member"][member["UserName"]] = {"type": "group_member", "info": member, "group": group}

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

func (this *WxBot) batch_get_group_members(){
    //批量获取所有群聊成员信息
    urlStr := this.base_uri + fmt.Sprintf("/webwxbatchgetcontact?type=ex&r=%s&pass_ticket=%s", int(time.Now().Unix()), this.pass_ticket)
    j := simplejson.New()
    j.Set("BaseRequest",this.base_request)
    j.Set("Count", len(this.group_list))
    list := []map[string]string{}
    for _,v :=range this.group_list{
        list = append(list, map[string]string{"UserName": v.UserName,"EncryChatRoomId":"",})
    }
    j.Set("List",list)
    params,_ := j.MarshalJSON()
    body,err := this.Post(urlStr,"raw", string(params))
    if err != nil{

    }
    //fmt.Printf(":::body:::%+v\n",string(body))
    j,_ = simplejson.NewJson(body)
    groupList := []*WxGroup{}
    gList := j.Get("ContactList").Interface()
    b,_ := json.Marshal(gList)
    json.Unmarshal(b,&groupList)
    fmt.Printf("::::groupList::::%+v\n\n", groupList[0])
    this.group_members = map[string][]*WxGroupMember{}
    this.encry_chat_room_id_list = map[string]string{}
    for _,group := range groupList{
        gid := group.UserName
        members := group.MemberList
        this. group_members[gid] = members
        this.encry_chat_room_id_list[gid] = group.EncryChatRoomId
    }
}

func (this *WxBot) test_sync_check() bool{
    for _, host1 := range []string{"webpush.", "webpush2."}{
        this.sync_host = host1+this.base_host
        retcode, _ := this.sync_check()
        if retcode == "0"{
            return true
        }
    }
    return false
}

func (this *WxBot) sync_check() (string,string){
    params := url.Values{
        "r": []string{strconv.Itoa(int(time.Now().Unix()))},
        "sid": []string{this.sid},
        "uin": []string{this.uin},
        "skey": []string{this.skey},
        "deviceid": []string{this.device_id},
        "synckey": []string{this.sync_key_str},
        "_": []string{strconv.Itoa(int(time.Now().Unix()))},
    }
    urlStr := "https://" + this.sync_host + "/cgi-bin/mmwebwx-bin/synccheck?" + params.Encode()
    fmt.Println("::::urlStr::::", urlStr)

    body,err := this.Get(urlStr)
    if len(body) > 500{
        fmt.Println("::::body::::",string(body[:500]))
    }else{
        fmt.Println("::::body::::",string(body))
    }
    if err != nil {
        fmt.Println("::::snyc_check::::",err)
        return "-1", "-1"
    }
    re,_ := regexp.Compile("window.synccheck=\\{retcode:\"([0-9]+?)\",selector:\"([0-9]+?)\"\\}")
    subMatch := re.FindSubmatch(body)
    if len(subMatch) < 3{
        return "-1", "-1"
    }
    retcode := subMatch[1]
    selector := subMatch[2]
    return string(retcode), string(selector)
}

func (this *WxBot) proc_msg(){
    this.test_sync_check()
    durationTime := int64(800*time.Millisecond)
    for {
        check_time := time.Now().UnixNano()

        retcode, selector := this.sync_check()
        // print "[DEBUG] sync_check:", retcode, selector
        switch {
        case retcode == "1100":  // 从微信客户端上登出
            break
        case retcode == "1101":  // 从其它设备上登了网页微信
            break
        case retcode == "0":
            switch {
            case selector == "2":  // 有新消息
                r := this.sync()
                if r != nil{
                    this.handle_msg(r)
                }
            case selector == "3":  // 未知
                r := this.sync()
                if r != nil{
                    this.handle_msg(r)}
            case selector == "4":  // 通讯录更新
                r := this.sync()
                if r != nil{
                    this.get_contact()}
            case selector == "6":  // 可能是红包
                r := this.sync()
                if r != nil{
                    this.handle_msg(r)}
            case selector == "7":  // 在手机上操作了微信
                r := this.sync()
                if r != nil{
                    this.handle_msg(r)}
            case selector == "0":  // 无事件

            default:
                fmt.Println("[DEBUG] sync_check:", retcode, selector)
                r := this.sync()
                if r !=nil{
                    this.handle_msg(r)
                }
            }
        default:
            fmt.Println("[DEBUG] sync_check:", retcode, selector)
            time.Sleep(10 * time.Second)
        }
        this.schedule()
        //except:
        //    print "[ERROR] Except in proc_msg"
        //    print format_exc()
        check_time = time.Now().UnixNano() - check_time
        if check_time < durationTime{
            time.Sleep(1*time.Second - time.Duration(check_time))
        }
    }
}

func (this *WxBot) sync()*WxSyncResponse{
    urlStr := this.base_uri + fmt.Sprintf("/webwxsync?sid=%s&skey=%s&lang=en_US&pass_ticket=%s", this.sid, this.skey, this.pass_ticket)
    params := map[string]interface{}{
        "BaseRequest": this.base_request,
        "SyncKey": this.sync_key,
        "rr": int(time.Now().Unix()),
    }
    b,_ := json.Marshal(params)
    j,_ := simplejson.NewJson(b)
    paramData,_ := j.MarshalJSON()
    body,err := this.Post(urlStr, "raw", string(paramData))
    if err != nil{
    }
    fmt.Println("::::sync::::",string(body))
    dic := &WxSyncResponse{}
    err = json.Unmarshal(body, dic)
    if err != nil {
        fmt.Println("::::sync err::::" , err)
        return nil
    }
    this.sync_key = dic.SyncKey

    conj := "";
    this.sync_key_str = ""
    for _,v := range this.sync_key.List{
        this.sync_key_str += conj + strconv.Itoa(v.Key) + "_" + strconv.Itoa(v.Val)
        if len(conj) == 0{
            conj ="|"
        }
    }

    return dic
}
/**
处理原始微信消息的内部函数
        msg_type_id:
            0 -> Init
            1 -> Self
            2 -> FileHelper
            3 -> Group
            4 -> Contact
            5 -> Public
            6 -> Special
            99 -> Unknown
 */
func (this *WxBot) handle_msg(r *WxSyncResponse){

    fmt.Printf(":::::::handle_msg dic:::::::::%+v\n",r)
    fmt.Printf(":::::::handle_msg msg:::::::::%+v\n",r.AddMsgList[0])
    /*msg_type_id := 0
    for _,msg := range r.AddMsgList{
        user := map[string]string{"id": msg.FromUserName, "name": "unknown"}
        switch{
        case msg["MsgType"] == 51 && msg["StatusNotifyCode"] == 4:  // init message
            msg_type_id = 0
            user["name"] = "system"
            //会获取所有联系人的username 和 wxid，但是会收到3次这个消息，只取第一次
            *//*if this.is_big_contact && len(this.full_user_name_list) == 0{
            this.full_user_name_list = msg["StatusNotifyUserName"].split(",")
            this.wxid_list = re.search(r"username&gt;(.*?)&lt;/username", msg["Content"]).group(1).split(",")
            with open(os.path.join(this.temp_pwd,"UserName.txt"), "w") as f:
            f.write(msg["StatusNotifyUserName"])
            with open(os.path.join(this.temp_pwd,"wxid.txt"), "w") as f:
            f.write(json.dumps(this.wxid_list))
            fmt.Println("[INFO] Contact list is too big. Now start to fetch member list .")
            this.get_big_contact()
        }*//*

        case msg["MsgType"] == 37:  // friend request
            msg_type_id = 37
        // content = msg["Content"]
        // username = content[content.index("fromusername="): content.index("encryptusername")]
        // username = username[username.index(""") + 1: username.rindex(""")]
        // print u"[Friend Request]"
        // print u"       Nickname：" + msg["RecommendInfo"]["NickName"]
        // print u"       附加消息："+msg["RecommendInfo"]["Content"]
        // // print u"Ticket："+msg["RecommendInfo"]["Ticket"] // Ticket添加好友时要用
        // print u"       微信号："+username //未设置微信号的 腾讯会自动生成一段微信ID 但是无法通过搜索 搜索到此人
        case msg["FromUserName"] == this.my_account["UserName"]:  // Self
            msg_type_id = 1
            user["name"] = "this.
        case msg["ToUserName"] == "filehelper":  // File Helper
            msg_type_id = 2
            user["name"] = "file_helper"
        case msg["FromUserName"][:2] == "@@":  // Group
            msg_type_id = 3
            user["name"] = this.get_contact_prefer_name(this.get_contact_name(user["id"]))
        case this.is_contact(msg.FromUserName):  // Contact
            msg_type_id = 4
            user["name"] = this.get_contact_prefer_name(this.get_contact_name(user["id"]))
        case this.is_public(msg.FromUserName):  // Public
            msg_type_id = 5
            user["name"] = this.get_contact_prefer_name(this.get_contact_name(user["id"]))
        case this.is_special(msg.FromUserName):  // Special
            msg_type_id = 6
            user["name"] = this.get_contact_prefer_name(this.get_contact_name(user["id"]))
        default:
            msg_type_id = 99
            user["name"] = "unknown"
        }
        if len(user["name"])==0{
            user["name"] = "unknown"
        }
        user["name"] = html.UnescapeString(user["name"])

        if this.DEBUG && msg_type_id != 0{
            fmt.Printf("[MSG] %s:\n" , user["name"])
        }
        content := this.extract_msg_content(msg_type_id, msg)
        message := map[string]interface{}{"msg_type_id": msg_type_id,
            "msg_id": msg["MsgId"],
            "content": content,
            "to_user_id": msg["ToUserName"],
            "user": user}
        this.handle_msg_all(message)
    }*/
}

func (this *WxBot) handle_msg_all(msg map[string]interface{}){

}
func (this *WxBot) schedule(){

}

func (this *WxBot) extract_msg_content(msg_type_id int, msg *WxMsg) string{
    return ""
}

func (this *WxBot) get_contact_name(name string) string{

    return ""
}

func (this *WxBot) get_contact_prefer_name(uid string){

}

func (this *WxBot) is_contact(uid string)bool{
    return true
}

func (this *WxBot) is_public(uid string)bool{
    return true

}

func (this *WxBot) is_special(uid string)bool{
    return true

}



func (this *WxBot) do_request(url string) (string, []byte) {
    body, err := this.Get(url)
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

