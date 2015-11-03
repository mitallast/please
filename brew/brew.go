package brew

import (
	"bufio"
	"log"
	"os/exec"
)

type BrewProvider struct {
}

func NewBrewProvider() *BrewProvider {
	return &BrewProvider{}
}

func (p *BrewProvider) Search(arg ...string) ([]string, error) {
	return brew("search", arg...)
}

func (p *BrewProvider) Install(arg ...string) ([]string, error) {
	return brew("install", arg...)
}

func brew(action string, arg ...string) ([]string, error) {
	cmd := exec.Command("brew", append([]string{action}, arg...)...)
	log.Printf("execute: %v", cmd.Args)
	lines := []string{}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return lines, err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return lines, err
	}
	if err := cmd.Start(); err != nil {
		return lines, err
	}
	//	stdout
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("[brew] %s", line)
		lines = append(lines, line)
	}
	if scanner.Err() != nil {
		return lines, nil
	}
	//	stderr
	scanner = bufio.NewScanner(stderr)
	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("[brew][stderr] %s", line)
	}
	if scanner.Err() != nil {
		return lines, nil
	}
	if err := cmd.Wait(); err != nil {
		log.Printf("exit status %d", cmd.ProcessState.Sys())
		return lines, err
	}
	return lines, nil
}