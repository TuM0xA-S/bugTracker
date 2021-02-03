package load
import (
	"bytes"
	"testing"
)
func TestFetch(t *testing.T) {
	data, err := Fetch("https://security-tracker.debian.org/tracker/data/json")
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(data, []byte("CVE-2018-10906")) {
		t.Fatal("fetch data is broken")
	}
	if !bytes.Contains(data, []byte("status")) {
		t.Fatal("fetch data is broken")
	}
	if !bytes.Contains(data, []byte("buster")) {
		t.Fatal("fetch data is broken")
	}
}
