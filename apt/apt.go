package apt

import (
	"github.com/mitallast/please/cmd"
	"log"
	"strings"
)

type AptProvider struct {
}

func Supports() bool {
	return cmd.LookPath("apt-get")
}

func NewProvider() *AptProvider {
	return &AptProvider{}
}

func (p *AptProvider) Search(arg ...string) ([]string, error) {
	cmd := cmd.Command("apt-cache", append([]string{"search"}, arg...)...)
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	defer cmd.Wait()
	matches := []string{}
	scanner := cmd.Scanner()
	for scanner.Scan() {
		line := scanner.Text()
		if index := strings.Index(line, " "); index > 0 {
			match := line[0:index]
			log.Printf("[apt] %s", match)
			matches = append(matches, match)
		}
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

func (p *AptProvider) Install(arg ...string) error {
	cmd := cmd.Command("sudo", append([]string{"apt-get", "install", "-y"}, arg...)...)
	if err := cmd.Start(); err != nil {
		return err
	}
	defer cmd.Wait()
	scanner := cmd.Scanner()
	for scanner.Scan() {
		log.Printf("[apt] %s", scanner.Text())
		if err := scanner.Err(); err != nil {
			return err
		}
	}
	return nil
}
