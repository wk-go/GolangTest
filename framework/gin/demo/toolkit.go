package main

import (
    "strings"
    "net/url"
    "log"
    "fmt"
)

// strtr()
//
// If the parameter length is 1, type is: map[string]string
// Strtr("baab", map[string]string{"ab": "01"}) will return "ba01"
// If the parameter length is 2, type is: string, string
// Strtr("baab", "ab", "01") will return "1001", a => 0; b => 1.
func Strtr(haystack string, params ...interface{}) string {
    ac := len(params)
    if ac == 1 {
        pairs := params[0].(map[string]string)
        length := len(pairs)
        if length == 0 {
            return haystack
        }
        oldnew := make([]string, length*2)
        for o, n := range pairs {
            if o == "" {
                return haystack
            }
            oldnew = append(oldnew, o, n)
        }
        return strings.NewReplacer(oldnew...).Replace(haystack)
    } else if ac == 2 {
        from := params[0].(string)
        to := params[1].(string)
        trlen, lt := len(from), len(to)
        if trlen > lt {
            trlen = lt
        }

        if trlen == 0 {
            return haystack
        } else {
            str := make([]uint8, len(haystack))
            var xlat [256]uint8
            var i int
            var j uint8
            if trlen == 1 {
                for i = 0; i < len(haystack); i++ {
                    if haystack[i] == from[0] {
                        str[i] = to[0]
                    } else {
                        str[i] = haystack[i]
                    }
                }
                return string(str)
            } else {
                for {
                    xlat[j] = j
                    if j++; j == 0 {
                        break
                    }
                }
                for i = 0; i < trlen; i++ {
                    xlat[from[i]] = to[i]
                }
                for i = 0; i < len(haystack); i++ {
                    str[i] = xlat[haystack[i]]
                }
                return string(str)
            }
        }
    }

    return haystack
}

func ChangeUrlParams(urlStr string, values ...interface{})string{
    u, err := url.Parse(urlStr)
    if err != nil {
        log.Fatal(err)
        return urlStr
    }
    if len(values) == 0 || len(values)%2 != 0{
        return urlStr
    }

    q := u.Query()
    key := ""
    for k, v := range values {
        if k%2 == 0 {
            key = fmt.Sprint(v)
        } else {
            q.Set(key, fmt.Sprint(v))
        }
    }
    u.RawQuery = q.Encode()
    return u.String()
}