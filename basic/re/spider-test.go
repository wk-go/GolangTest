package main

import (
    "regexp"
    "fmt"
)

func main(){

    spiderPattern := map[string]*regexp.Regexp{
        "spider": regexp.MustCompile("compatible; (.*?)/(.*?);"),
        "bot": regexp.MustCompile("compatible; (.*?)/(.*?);"),
        "360Spider": regexp.MustCompile(" (360Spider)"),
    }

    spiderDemos:=[]string{
        "Mozilla/5.0 (Linux;u;Android 4.2.2;zh-cn;) AppleWebKit/534.46 (KHTML,like Gecko) Version/5.1 Mobile Safari/10600.6.3 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)",
        "Mozilla/5.0 (Linux;u;Android 4.2.2;zh-cn;) AppleWebKit/534.46 (KHTML,like Gecko) Version/5.1 Mobile Safari/10600.6.3 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)",
        "Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)",
        "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36; 360Spider",
        "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
        "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.96 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
        "Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots)",
    }

    for i, uaStr := range spiderDemos{
        for key,pattern := range spiderPattern {
            if pattern.MatchString(uaStr) {
                result := pattern.FindAllStringSubmatch(uaStr, -1)
                fmt.Printf("%02d(%s): %q\n", i+1, key, result)
                break
            }
            fmt.Printf("%02d(%s) not matched\n", i+1, key)
        }
    }
}