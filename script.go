package ginit

import (
	"fmt"
	"os"
	"os/exec"
)

type ScriptOptions struct {
	Path string
	Args []string
	Env  []string
}

func RunScript(opts ScriptOptions) error {
	cmd := exec.Command(opts.Path, opts.Args...)
	if len(opts.Env) == 0 {
		cmd.Env = os.Environ()
	} else {
		cmd.Env = opts.Env
	}
	raw, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(raw))
	return nil
}
