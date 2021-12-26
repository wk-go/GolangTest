package conf_handle

import "regexp"

func getStrings(_re, s string) [][]string {
	re := regexp.MustCompile(_re)
	if re == nil {
		return [][]string{}
	}
	return re.FindAllStringSubmatch(s, -1)
}
