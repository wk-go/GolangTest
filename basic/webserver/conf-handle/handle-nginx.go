package conf_handle

import (
	"fmt"
	"reflect"
	"regexp"
)

type NginxConfig struct {
	Server []*NginxVhost `json:"server"`
}
type NginxVhost struct {
	Listen     string `json:"listen" regexp:"[^#]listen +?(.*?);"`
	ServerName string `json:"server_name" regexp:"[^#]server_name +?(.*?);"`
	AccessLog  string `json:"access_log" regexp:"[^#]access_log +?(.*?);"`
	Root       string `json:"root" regexp:"[^#]root +?(.*?);"`
}

func ParseNginxVhost(s string) (vhosts []*NginxVhost, err error) {
	var config NginxConfig
	regServer := regexp.MustCompile("server\\s*?\\{")
	_groupString := regServer.Split(s, -1)[1:]

	config = NginxConfig{Server: make([]*NginxVhost, len(_groupString))}

	for i, str := range _groupString {
		_vhost := NginxVhost{}
		_v := reflect.ValueOf(&_vhost).Elem()
		_t := _v.Type()
		for _i := 0; _i < _v.NumField(); _i++ {
			_field := _t.Field(_i)
			_regexp := _field.Tag.Get("regexp")
			strSlice := getStrings(_regexp, str)
			_v.FieldByName(_field.Name).Set(reflect.ValueOf(strSlice[0][1]))
		}
		config.Server[i] = &_vhost
	}
	if err != nil {
		fmt.Printf("Failed to load configuration: %s\n", err)
	}
	fmt.Printf("Configuration is %#v\n", config)
	vhosts = config.Server
	return
}
