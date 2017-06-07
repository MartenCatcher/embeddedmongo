package embeded_mongo

import (
	"os"
	"os/exec"
	"log"
	"distribution/uuid"
)

type Process struct {
	Pid int
	Tmp string
	c   *exec.Cmd
}

func NewProcess(app string, dir string) (*Process, error) {
	id, err := moveToTmp(app, dir)
	if err != nil {
		return nil, err
	}

	tmp := dir + id

	c := exec.Command(tmp+"/"+app, "--logpath", tmp+"/mongo.log", "--dbpath", tmp+"/db")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err = c.Start(); err != nil {
		return nil, err
	}

	log.Printf("Started Mongo instance with pid=%+v", c.Process.Pid)

	return &Process{c: c, Tmp: tmp, Pid: c.Process.Pid}, nil
}

func moveToTmp(app string, dir string) (string, error) {
	id := uuid.Generate().String()

	tmp := dir + id
	if err := os.MkdirAll(tmp+"/db", 0755); err != nil {
		return "", err
	}
	if err := os.Rename(dir+"/"+app, tmp+"/"+app); err != nil {
		return "", err
	}

	return id, nil
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
