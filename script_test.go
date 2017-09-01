package ginit

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

var testScript = `
#!/bin/sh
set -e

for i in $(seq 1 100); do 
	(( $i % 5 == 0 )) && echo $i
	echo >&2 $i
done
`

func TestCall(t *testing.T) {
	dir, err := ioutil.TempDir("/tmp", "ginit")
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile(filepath.Join(dir, "test.sh"), []byte(testScript), 0777)
	if err != nil {
		t.Fatal(err)
	}
	var (
		stdout int
		stderr int
	)
	err = Call(ScriptArgs{
		Cmd:      "/bin/sh",
		Args:     []string{filepath.Join(dir, "test.sh")},
		OnStdout: func(str string) { stdout++ },
		OnStderr: func(str string) { stderr++ },
	})
	if err != nil {
		t.Fatal(err)
	}
	if stdout != 20 {
		t.Errorf("stdout produced %d lines, should be 20", stdout)
	}
	if stderr != 100 {
		t.Errorf("stderr produced %d lines, should be 100", stderr)
	}
}
