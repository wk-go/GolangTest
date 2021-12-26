package conf_handle

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestHandleApacheVhost(t *testing.T) {
	dataBytes, err := ioutil.ReadFile("apache/video.domain.com.conf")
	if err != nil {
		t.Error(err)
	}

	vhosts, err := ParseApacheVhost(string(dataBytes))
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(vhosts[0].Port, "*:") {
		fmt.Printf("%+v, %t\n", vhosts[0].Port, vhosts[0].Port == "*:88")
		t.Error(errors.New("port error"))
	}
	if value, ok := vhosts[0].PhpAdminValue["open_basedir"]; ok && !strings.Contains(value, "var/tmp") {
		t.Error(errors.New("open_basedir error"))
	}
}
