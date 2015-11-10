package apt

import (
	"bufio"
	"log"
	"os/exec"
	"strings"
)

type AptProvider struct {
}

func Supports() bool {
	_, err := exec.LookPath("apt-get")
	return err == nil
}

func NewProvider() *AptProvider {
	return &AptProvider{}
}

func (p *AptProvider) Search(arg ...string) ([]string, error) {
	cmd := exec.Command("apt-cache", append([]string{"search"}, arg...)...)
	log.Printf("execute: %v", cmd.Args)
	matches := []string{}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return matches, err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return matches, err
	}
	if err := cmd.Start(); err != nil {
		return matches, err
	}
	//	stdout
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		if index := strings.Index(line, " "); index > 0 {
			match := line[0:index]
			log.Printf("[apt] %s", match)
			matches = append(matches, match)
		}
	}
	if scanner.Err() != nil {
		return matches, nil
	}
	//	stderr
	scanner = bufio.NewScanner(stderr)
	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("[apt] stderr: %s", line)
	}
	if scanner.Err() != nil {
		return matches, nil
	}
	if err := cmd.Wait(); err != nil {
		log.Printf("exit status %d", cmd.ProcessState.Sys())
		return matches, err
	}
	return matches, nil
}

func (p *AptProvider) Contains(arg ...string) ([]string, error) {
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

func (p *AptProvider) Install(arg ...string) ([]string, error) {
	cmd := exec.Command("apt-get", append([]string{"install", "-y"}, arg...)...)
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
		log.Printf("[apt] %s", line)
		lines = append(lines, line)
	}
	if scanner.Err() != nil {
		return lines, nil
	}
	//	stderr
	scanner = bufio.NewScanner(stderr)
	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("[apt][stderr] %s", line)
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
