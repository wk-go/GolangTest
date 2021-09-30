package conf_handle

import (
	"errors"
	"regexp"
	"strings"
)

var (
	apacheVhostMap = map[string]string{
		"Port":            "",
		"ServerAdmin":     "ServerAdmin *?(.*?)\n",
		"php_admin_value": "php_admin_value *?(.*?)\n",
		"DocumentRoot":    "DocumentRoot *?(.*?)\n",
		"ServerName":      "ServerName *?(.*?)\n",
		"ErrorLog":        "ErrorLog *?(.*?)\n",
		"CustomLog":       "CustomLog *?(.*?)\n",
		"Directory":       "<Directory *?(.*?)>",
	}
)

type ApacheVhost struct {
	Port          string
	ServerAdmin   string
	PhpAdminValue map[string]string
	DocumentRoot  string
	ServerName    string
	ErrorLog      string
	CustomLog     string
	Directory     string
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
		for key, _reg := range apacheVhostMap {
			if _reg == "" {
				continue
			}
			_regData := apacheGetStrings(_reg, result[2])
			if len(_regData) == 0 {
				continue
			}
			switch key {
			case "ServerAdmin":
				vhosts[i].ServerAdmin = _regData[0][1]
			case "php_admin_value":
				vhosts[i].PhpAdminValue = make(map[string]string, len(_regData))
				for _, _val := range _regData {
					_s := strings.Split(_val[1], " ")
					println(_val[1], _s)
					_x := _s
					vhosts[i].PhpAdminValue[_x[0]] = _x[1]
				}
			case "DocumentRoot":
				vhosts[i].DocumentRoot = _regData[0][1]
			case "ServerName":
				vhosts[i].ServerName = _regData[0][1]
			case "ErrorLog":
				vhosts[i].ErrorLog = _regData[0][1]
			case "CustomLog":
				vhosts[i].CustomLog = _regData[0][1]
			case "Directory":
				vhosts[i].Directory = _regData[0][1]
			}
		}
	}

	return
}
func apacheGetStrings(_re, s string) [][]string {
	re := regexp.MustCompile(_re)
	if re == nil {
		return [][]string{}
	}
	return re.FindAllStringSubmatch(s, -1)
}
