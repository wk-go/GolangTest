package conf_handle

import (
	"errors"
	"io/ioutil"
	"strings"
	"testing"
)

func TestParseNginxVhost(t *testing.T) {
	dataBytes, err := ioutil.ReadFile("nginx/apivcloud.domain.com.conf")
	if err != nil {
		t.Error(err)
	}

	vhosts, err := ParseNginxVhost(string(dataBytes))
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(vhosts[0].Listen, "80") {
		t.Error(errors.New("port error"))
	}
	if !strings.Contains(vhosts[0].Root, "frontend") {
		t.Error(errors.New("root error"))
	}
}
