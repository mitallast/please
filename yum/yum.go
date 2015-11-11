package yum

import (
	"github.com/mitallast/please/cmd"
	"log"
	"strings"
)

type YumProvider struct {
}

func Supports() bool {
	return cmd.LookPath("yum")
}

func NewProvider() *YumProvider {
	return &YumProvider{}
}

func (p *YumProvider) Search(arg ...string) ([]string, error) {
	cmd := cmd.Command("yum", append([]string{"search", "--quiet"}, arg...)...)
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	defer cmd.Wait()
	matches := []string{}
	scanner := cmd.Scanner()
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "=") || strings.HasPrefix(line, " ") {
			continue
		}
		if index := strings.Index(line, " "); index > 0 {
			match := line[0:index]
			log.Printf("[yum] %s", match)
			matches = append(matches, match)
		}
	}
	return matches, nil
}

func (p *YumProvider) Contains(arg ...string) ([]string, error) {
	cmd := cmd.Command("yum", append([]string{"list", "--quiet"}, arg...)...)
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	defer cmd.Wait()
	matches := []string{}
	scanner := cmd.Scanner()
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "=") || strings.HasPrefix(line, " ") {
			continue
		}
		if strings.HasPrefix(line, "Installed ") || strings.HasPrefix(line, "Available ") {
			continue
		}
		if index := strings.Index(line, " "); index > 0 {
			match := line[0:index]
			log.Printf("[yum] %s", match)
			matches = append(matches, match)
		}
	}
	return matches, nil
}

func (p *YumProvider) Install(arg ...string) error {
	cmd := cmd.Command("sudo", append([]string{"yum", "install", "-y"}, arg...)...)
	if err := cmd.Start(); err != nil {
		return err
	}
	defer cmd.Wait()
	scanner := cmd.Scanner()
	for scanner.Scan() {
		log.Printf("[yum] %s", scanner.Text())
		if err := scanner.Err(); err != nil {
			return err
		}
	}
	return nil
}
