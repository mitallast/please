package brew

import (
	"github.com/mitallast/please/cmd"
	"log"
	"strings"
)

type BrewProvider struct {
}

func Supports() bool {
	return cmd.LookPath("brew")
}

func NewProvider() *BrewProvider {
	return &BrewProvider{}
}

func (p *BrewProvider) Search(arg ...string) ([]string, error) {
	cmd := cmd.Command("brew", append([]string{"search"}, arg...)...)
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	defer cmd.Wait()
	matches := []string{}
	scanner := cmd.Scanner()
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, " (installed)", "", -1)
		log.Printf("[brew] %s", line)
		matches = append(matches, line)
	}
	return matches, nil
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

func (p *BrewProvider) Install(arg ...string) error {
	cmd := cmd.Command("brew", append([]string{"install"}, arg...)...)
	if err := cmd.Start(); err != nil {
		return err
	}
	defer cmd.Wait()
	scanner := cmd.Scanner()
	for scanner.Scan() {
		log.Printf("[brew] %s", scanner.Text())
		if err := scanner.Err(); err != nil {
			return err
		}
	}
	return nil
}
