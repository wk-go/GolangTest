package conf_handle

import (
	"errors"
	"reflect"
	"regexp"
	"strings"
)

type ApacheVhost struct {
	Port          string            `json:"port" regexp:""`
	ServerAdmin   string            `json:"server_admin" regexp:"ServerAdmin *?(.*?)\n"`
	PhpAdminValue map[string]string `json:"php_admin_value" regexp:"php_admin_value *?(.*?)\n"`
	DocumentRoot  string            `json:"document_root" regexp:"DocumentRoot *?(.*?)\n"`
	ServerName    string            `json:"server_name" regexp:"ServerName *?(.*?)\n"`
	ErrorLog      string            `json:"error_log" regexp:"ErrorLog *?(.*?)\n"`
	CustomLog     string            `json:"custom_log" regexp:"CustomLog *?(.*?)\n"`
	Directory     string            `json:"directory" regexp:"<Directory *?(.*?)>"`
}

func ParseApacheVhost(s string) (vhosts []*ApacheVhost, err error) {
	reVhost := regexp.MustCompile("<VirtualHost[ ]+(.*?)>((?s).*?)</VirtualHost>")
	s = strings.Replace(s, "\r", "", -1)
	resultVhost := reVhost.FindAllStringSubmatch(s, -1)
	if len(resultVhost) == 0 {
		err = errors.New("no VirtualHost section")
		return
	}

	vhosts = make([]*ApacheVhost, len(resultVhost))
	for i, result := range resultVhost {
		vhosts[i] = &ApacheVhost{}
		vhosts[i].Port = result[1]
		_v := reflect.ValueOf(vhosts[i]).Elem()
		_t := _v.Type()
		for _i := 0; _i < _v.NumField(); _i++ {
			_field := _t.Field(_i)
			if _field.Name == "Port" {
				continue
			}

			_regexp := _field.Tag.Get("regexp")
			if _regexp == "" {
				continue
			}

			strSlice := getStrings(_regexp, result[2])
			if _field.Name == "PhpAdminValue" {
				vhosts[i].PhpAdminValue = make(map[string]string, len(strSlice))
				for _, _val := range strSlice {
					_s := strings.Split(_val[1], " ")
					println(_val[1], _s)
					_x := _s
					vhosts[i].PhpAdminValue[_x[0]] = _x[1]
				}
				continue
			}
			_v.FieldByName(_field.Name).Set(reflect.ValueOf(strSlice[0][1]))
		}
	}

	return
}
