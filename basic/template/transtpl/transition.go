package transtpl

import (
	"bytes"
	"io/ioutil"
	"text/template"
)

func TransitionFileToFile(src, dest string, data map[string]interface{}) (err error) {
	result, err := TransitionFileToBytes(src, data)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(dest, result, 0755)
	return
}

func TransitionFileToBytes(src string, data map[string]interface{}) (result []byte, err error) {
	srcContent, err := ioutil.ReadFile(src)
	if err != nil {
		return
	}
	result, err = Transition(src, string(srcContent), data)
	if err != nil {
		return
	}
	return
}

func Transition(name, text string, data map[string]interface{}) (result []byte, err error) {
	tmpl, err := template.New(name).Parse(text)
	if err != nil {
		return
	}
	var buf = bytes.NewBuffer(make([]byte, 0))
	err = tmpl.Execute(buf, data)
	if err != nil {
		return
	}
	return buf.Bytes(), nil
}
