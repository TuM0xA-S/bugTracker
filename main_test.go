package main

import (
	"bytes"
	"testing"

	"github.com/TuM0xA-S/bugTracker/load"
)

func TestMain(t *testing.T) {
	host := parseConfig().Host
	have, err := load.Fetch("http://" + host + "/api/update")
	if err != nil {
		t.Fatal("troubles when trying to update:", err)
	}
	want := []byte("\"update done\"\n")
	if !bytes.Equal(want, have) {
		t.Fatal("returned data not valid")
	}
	have, err = load.Fetch("http://" + host + "/api/cve/2018-10906")
	if err != nil {
		t.Fatal("troubles when trying to update:")
	}
	if !bytes.Contains(have, []byte("debian")) ||
		!bytes.Contains(have, []byte("ubuntu")) ||
		!bytes.Contains(have, []byte("redhat")) {
		t.Fatal("returned data not valid")
	}
}
