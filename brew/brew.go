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
	cmd := exec.Command("brew", append([]string{"search"}, arg...)...)
	log.Printf("execute: %v", cmd.Args)
	lines := []string{}
	reader, err := cmd.StdoutPipe()
	if err != nil {
		return lines, err
	}
	if err := cmd.Start(); err != nil {
		return lines, err
	}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		log.Print(line)
		lines = append(lines, line)
	}
	if scanner.Err() != nil {
		return lines, nil
	}
	if err := cmd.Wait(); err != nil {
		return lines, err
	}
	return lines, nil
}

func (p *BrewProvider) Install(arg ...string) ([]byte, error) {
	cmd := exec.Command("brew", append([]string{"search"}, arg...)...)
	return cmd.CombinedOutput()
}
