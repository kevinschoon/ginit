package ginit

import (
	"testing"
)

func TestKeyFS(t *testing.T) {
	kfs := KeyFS{Base: "/sys"}
	results, err := kfs.Find("/class/net", "address", true)
	if err != nil {
		t.Fatal(err)
	}
	if len(results) == 0 {
		t.Error("no results")
	}
}
