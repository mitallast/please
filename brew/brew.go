package brew

import (
	"bufio"
	"log"
	"os/exec"
	"strings"
)

type BrewProvider struct {
}

func Supports() bool {
	_, err := exec.LookPath("brew")
	return err == nil
}

func NewProvider() *BrewProvider {
	return &BrewProvider{}
}

func (p *BrewProvider) Search(arg ...string) ([]string, error) {
	cmd := exec.Command("brew", append([]string{"search"}, arg...)...)
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
		line = strings.Replace(line, " (installed)", "", -1)
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

func (p *BrewProvider) Contains(arg ...string) ([]string, error) {
	lines, err := p.Search(arg...)
	matches := []string{}
	if err != nil {
		return matches, err
	} else {
		for _, found := range lines {
			if found == arg[0] {
				log.Printf("contains: %s", found)
				matches = append(matches, found)
			}
		}
		return matches, nil
	}
}

func (p *BrewProvider) Install(arg ...string) ([]string, error) {
	cmd := exec.Command("brew", append([]string{"install"}, arg...)...)
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
