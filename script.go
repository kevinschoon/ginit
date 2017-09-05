package ginit

import (
	"bufio"
	"os"
	"os/exec"
	"sync"
)

type ScriptArgs struct {
	Cmd  string
	Args []string
	// OnStdout is passed each line
	// the script writes to stdout
	OnStdout func(string)
	// OnStderr is passed each line
	// the script writes to stderr
	OnStderr func(string)
}

// Call executes the script parameters copying the
// existing environment into the command. Call blocks
// until the command is finished and stdout/stderr
// have been synchronized with OnStdout/OnSterr functions.
func Call(args ScriptArgs) error {
	cmd := exec.Command(args.Cmd, args.Args...)
	cmd.Env = os.Environ()
	var wg sync.WaitGroup
	if args.OnStdout != nil {
		wg.Add(1)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return err
		}
		go func() {
			defer wg.Done()
			scanner := bufio.NewScanner(stdout)
			for scanner.Scan() {
				args.OnStdout(scanner.Text())
			}
		}()
	}
	if args.OnStderr != nil {
		wg.Add(1)
		stderr, err := cmd.StderrPipe()
		if err != nil {
			return err
		}
		go func() {
			defer wg.Done()
			scanner := bufio.NewScanner(stderr)
			for scanner.Scan() {
				args.OnStderr(scanner.Text())
			}
		}()
	}
	err := cmd.Start()
	if err != nil {
		return err
	}
	// Ensure any stdout/stderr function
	// calls are synchronized before exiting.
	wg.Wait()
	return cmd.Wait()
}
