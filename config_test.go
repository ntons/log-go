package log

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestYAML2JSON(t *testing.T) {
	b1, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		t.Fatal("read config.yaml fail:", err)
		return
	}
	b2, err := ioutil.ReadFile("config.json")
	if err != nil {
		t.Fatal("read config.json fail:", err)
		return
	}
	if b1, err = YAML2JSON(b1); err != nil {
		t.Fatal("YAML2JSON fail:", err)
		return
	}
	b := bytes.NewBuffer(nil)
	if err = json.Compact(b, b1); err != nil {
		t.Fatal("json.Compact fail:", err)
		return
	}
	if b2 = b.Bytes(); !bytes.Equal(b1, b2) {
		t.Fatal("TestYAML2JSON fail: not equal")
		return
	}
}
