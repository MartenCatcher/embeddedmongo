package embeded_mongo

import (
	"os"
	"os/exec"
	"log"
)

type Process struct {
	Pid int
	Tmp string
	c   *exec.Cmd
}

func NewProcess(app string, dir string) (*Process, error) {
	if err := CreateDir(dir+"/db"); err != nil {
		return nil, err
	}

	c := exec.Command(dir+"/"+app, "--logpath", dir+"/mongo.log", "--dbpath", dir+"/db")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Start(); err != nil {
		return nil, err
	}

	log.Printf("Started Mongo instance with pid=%+v", c.Process.Pid)

	return &Process{c: c, Tmp: dir, Pid: c.Process.Pid}, nil
}

func (p *Process) Stop() error {
	var err error
	defer func() {
		if err := os.RemoveAll(p.Tmp); err != nil {
			log.Printf("Can't remove tmp dir: %v", err)
		}
	}()

	if err = p.c.Process.Kill(); err != nil {
		log.Printf("Process kill error [pid='%v', err=%v]", p.Pid, err)
		return err
	}

	if err = p.c.Wait(); err != nil {
		log.Printf("Process completion waiting error [pid='%v', err=%v]", p.Pid, err)
		return err
	}

	log.Printf("Process [pid='%v'] stopped", p.Pid)
	return nil
}
