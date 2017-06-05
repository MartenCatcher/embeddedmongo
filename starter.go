package embeded_mongo

import (
	"os"
	"os/exec"
	"log"
	"syscall"
)

type Process struct {
	Pid int
	c *exec.Cmd
}

func NewProcess(command string, args ...string) (*Process, error) {
	c := exec.Command(command, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Start()

	log.Printf("Started Mongo instance with pid=%+v", c.Process.Pid)

	return &Process{c: c, Pid: c.Process.Pid}, nil
}

func (p *Process) Stop() error {
	var err error
	if err = p.c.Process.Signal(syscall.SIGTERM); err != nil {
		return err
	}

	if err = p.c.Wait(); err != nil {
		return err
	}

	log.Printf("Process [pid='%s'] stopped", p.Pid)
	return nil
}