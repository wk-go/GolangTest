package transtpl

import "testing"

type PeerTemplate struct {
	Count uint
	SANS  []string
}
type PeerUsers struct {
	Count int
}

func TestTransitionFileToFile(t *testing.T) {
	srcFile := "examples/peer.tpl.yaml"
	destFile := "examples/out/pear.yaml"
	data := map[string]interface{}{
		"Name":          "myName",
		"Domain":        "myDomain",
		"EnableNodeOUs": "true",
		"Template": &PeerTemplate{
			Count: 100,
			SANS:  []string{"localhost", "other_host"},
		},
		"Users": &PeerUsers{
			Count: 1000,
		},
	}
	err := TransitionFileToFile(srcFile, destFile, data)
	if err != nil {
		t.Error(err)
	}
}
