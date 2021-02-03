package redhat

import (
	"testing"
)

func TestLoad(t *testing.T) {
	list, err := Load()
	if err != nil {
		t.Fatal(err)
	}

	if len(list) <= 20000 {
		t.Fatal("too small")
	}
	
	ok := false
	for _, v := range list {
		if v.CVE == "CVE-2016-10200" {
			ok = true
			t.Log(v)
			if v.Source != "redhat" || v.Data == nil || len(v.Packages) != 4 {
				ok = false
			}
			break
		}
	}
	if !ok {
		t.Fatal("bug CVE-2016-10200 parse failed")
	}
}
