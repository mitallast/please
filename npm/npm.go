package npm

import (
	"bufio"
	"encoding/json"
	"github.com/mitallast/please/cmd"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type NpmProvider struct {
}

type NpmPackage struct {
	Name string
}

func Supports() bool {
	return cmd.LookPath("npm")
}

func NewProvider() *NpmProvider {
	return &NpmProvider{}
}

func (p *NpmProvider) Search(arg ...string) ([]string, error) {
	cache, err := p.getCache()
	if err != nil {
		return nil, err
	}

	packagesFile := filepath.Join(cache, "registry.npmjs.org/-/all/.cache.json")
	file, err := os.Open(packagesFile)

	if err != nil {
		if os.IsNotExist(err) {
			return searchNpm(arg...)
		} else {
			return nil, err
		}
	} else {
		defer file.Close()
		return searchJson(file, arg...)
	}
}

func searchJson(file *os.File, arg ...string) ([]string, error) {
	matches := []string{}
	reader := bufio.NewReader(file)
	dec := json.NewDecoder(reader)
	// read open bracket
	if _, err := dec.Token(); err != nil {
		return matches, err
	}

	var npmPackage NpmPackage
	// while the array contains values
	for dec.More() {
		if _, err := dec.Token(); err != nil {
			return matches, err
		}
		// decode an array value (Message)
		err := dec.Decode(&npmPackage)
		// ignore unmarshal error
		if err == nil {
			for _, search := range arg {
				if search == npmPackage.Name {
					log.Printf("[npm] %s\n", npmPackage.Name)
					matches = append([]string{npmPackage.Name}, matches...)
				} else if strings.Contains(npmPackage.Name, search) {
					log.Printf("[npm] %s\n", npmPackage.Name)
					matches = append(matches, npmPackage.Name)
				}
			}
		}
	}

	// read close bracket
	if _, err := dec.Token(); err != nil {
		return matches, err
	}

	return matches, nil
}

func searchNpm(arg ...string) ([]string, error) {
	cmd := cmd.Command("npm", append([]string{"--no-color", "search"}, arg...)...)
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	defer cmd.Wait()
	matches := []string{}
	scanner := cmd.Scanner()
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "NAME") {
			if index := strings.Index(line, " "); index > 0 {
				match := line[0:index]
				log.Printf("[npm] %s", match)
				matches = append(matches, match)
			}
		}
	}
	return matches, nil
}

func (p *NpmProvider) Contains(arg ...string) ([]string, error) {
	lines, err := p.Search(arg...)
	if err != nil {
		return nil, err
	} else {
		matches := []string{}
		for _, found := range lines {
			if found == arg[0] {
				log.Printf("[npm] contains: %s", found)
				matches = append(matches, found)
			}
		}
		return matches, nil
	}
}

func (p *NpmProvider) Install(arg ...string) error {
	cmd := cmd.Command("npm", append([]string{"--no-color", "install"}, arg...)...)
	if err := cmd.Start(); err != nil {
		return err
	}
	defer cmd.Wait()
	scanner := cmd.Scanner()
	for scanner.Scan() {
		log.Printf("[npm] %s", scanner.Text())
		if err := scanner.Err(); err != nil {
			return err
		}
	}
	return nil
}

func (p *NpmProvider) getCache() (string, error) {
	cmd := cmd.Command("npm", "--no-color", "config", "get", "cache")
	if err := cmd.Start(); err != nil {
		return "", err
	}
	defer cmd.Wait()
	scanner := cmd.Scanner()
	if scanner.Scan() {
		path := strings.Trim(scanner.Text(), "\n\r")
		log.Printf("[npm] %s", path)
		return path, nil
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", os.ErrNotExist
}
