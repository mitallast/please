package npm

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strings"
	"encoding/json"
	"path/filepath"
)

type NpmProvider struct {
}

type NpmPackage struct {
	Name        string
	description string
}

func NewNpmProvider() *NpmProvider {
	return &NpmProvider{}
}

func (p *NpmProvider) Search(arg ...string) ([]string, error) {
	cache, err := p.getCache()
	matches := []string{}
	if err != nil {
		return matches, err
	}

	packagesFile := filepath.Join(cache, "registry.npmjs.org/-/all/.cache.json");
	file, err := os.Open(packagesFile)
	if err != nil {
		return matches, err
	}
	reader := bufio.NewReader(file)
	dec := json.NewDecoder(reader)
	// read open bracket
	if _, err = dec.Token(); err != nil {
		return matches, err
	}

	var npmPackage NpmPackage
	// while the array contains values
	for dec.More() {
		if _, err = dec.Token(); err != nil {
			return matches, err
		}
		// decode an array value (Message)
		err = dec.Decode(&npmPackage)
		// ignore unmarshal error
		if err == nil {
			for _, search := range arg {
				if search == npmPackage.Name {
					log.Printf("found: %s\n", npmPackage.Name)
					matches = append([]string{npmPackage.Name}, matches...)
				} else if strings.Contains(npmPackage.Name, search) {
					log.Printf("found: %s\n", npmPackage.Name)
					matches = append(matches, npmPackage.Name)
				}
			}
		}
	}

	// read close bracket
	if _, err = dec.Token(); err != nil {
		return matches, err
	}

	if err := file.Close(); err != nil {
		return matches, err
	}
	return matches, nil
}

func (p *NpmProvider) Contains(arg ...string) ([]string, error) {
	cache, err := p.getCache()
	matches := []string{}
	if err != nil {
		return matches, err
	}

	packagesFile := filepath.Join(cache, "registry.npmjs.org/-/all/.cache.json");
	log.Println("[npm] path to cache ", packagesFile)
	file, err := os.Open(packagesFile)
	if err != nil {
		return matches, err
	}
	reader := bufio.NewReader(file)
	dec := json.NewDecoder(reader)
	// read open bracket
	if _, err = dec.Token(); err != nil {
		return matches, err
	}

	var npmPackage NpmPackage
	// while the array contains values
	for dec.More() {
		if _, err = dec.Token(); err != nil {
			return matches, err
		}
		// decode an array value (Message)
		err = dec.Decode(&npmPackage)
		// ignore unmarshal error
		if err == nil {
			for _, search := range arg {
				if search == npmPackage.Name {
					log.Printf("found: %s\n", npmPackage.Name)
					matches = append(matches, npmPackage.Name)
				}
			}
		}
	}

	// read close bracket
	if _, err = dec.Token(); err != nil {
		return matches, err
	}

	if err := file.Close(); err != nil {
		return matches, err
	}
	return matches, nil
}

func (p *NpmProvider) Install(arg ...string) ([]string, error) {
	cmd := exec.Command("npm", append([]string{"install"}, arg...)...)
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
		log.Printf("[npm] %s", line)
		lines = append(lines, line)
	}
	if scanner.Err() != nil {
		return lines, nil
	}
	//	stderr
	scanner = bufio.NewScanner(stderr)
	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("[npm][stderr] %s", line)
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

func (p *NpmProvider) getCache() (string, error) {
	cmd := exec.Command("npm", "config", "get", "cache")
	log.Printf("execute: %v", cmd.Args)
	out, err := cmd.CombinedOutput()
	outStr := string(out)
	outStr = strings.Trim(outStr, "\n\r")
	log.Printf("[npm] cache [%s]\n", outStr)
	return outStr, err
}