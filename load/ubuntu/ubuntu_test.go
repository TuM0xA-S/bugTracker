package ubuntu

import (
	"testing"
)

func TestLoad(t *testing.T) {
	bugs, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	need := []string {"CVE-2012-5534", "CVE-2021-3345", "CVE-2014-9970", "CVE-2017-10227"}
	cnt := 0
	for _, b := range bugs {
		if find(need, b.CVE) {
			cnt++
		}
	}
	if cnt < len(need) {
		t.Fatalf("values was not parsed")
	}
}

func find(a []string, key string) bool {
	for _, s := range a {
		if s == key {
			return true
		}
	}
	return false
}
