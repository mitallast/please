package cmd

import (
	"bufio"
	"io"
	"log"
	"os/exec"
)

type Cmd struct {
	name   string
	cmd    *exec.Cmd
	stdout io.ReadCloser
	stderr io.ReadCloser
}

func LookPath(path string) bool {
	_, err := exec.LookPath(path)
	return err == nil
}

func Command(name string, arg ...string) *Cmd {
	return &Cmd{
		name: name,
		cmd:  exec.Command(name, arg...),
	}
}

func (c *Cmd) Start() error {
	var err error
	if c.stdout, err = c.cmd.StdoutPipe(); err != nil {
		return err
	}
	if c.stderr, err = c.cmd.StderrPipe(); err != nil {
		return err
	}
	if err = c.cmd.Start(); err != nil {
		return err
	}
	return nil
}

func (c *Cmd) Stdout() io.Reader {
	return c.stdout
}

func (c *Cmd) Scanner() *bufio.Scanner {
	return bufio.NewScanner(c.stdout)
}

func (c *Cmd) Wait() error {
	scanner := bufio.NewScanner(c.stderr)
	for scanner.Scan() {
		log.Printf("[%s] error: %s", c.name, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return c.cmd.Wait()
}
