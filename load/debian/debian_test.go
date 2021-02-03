package debian

import (
	"testing"
)

func TestLoad(t *testing.T) {
	bugs, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	ok := false
	for _, v := range bugs {
		if v.CVE == "CVE-2018-10906" {
			ok = true
			if v.Data == nil || v.Source != "debian" || len(v.Packages) != 2 {
				ok = false
			}
			break
		}
	}
	if !ok {
		t.Fatalf("bug CVE-2018-10906 parsed incorrectly")
	}
}
