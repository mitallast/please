package npm

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type NpmProvider struct {
}

type NpmPackage struct {
	Name        string
	description string
}

func Supports() bool {
	_, err := exec.LookPath("npm")
	return err == nil
}

func NewProvider() *NpmProvider {
	return &NpmProvider{}
}

func (p *NpmProvider) Search(arg ...string) ([]string, error) {
	cache, err := p.getCache()
	if err != nil {
		return []string{}, err
	}

	packagesFile := filepath.Join(cache, "registry.npmjs.org/-/all/.cache.json")
	file, err := os.Open(packagesFile)

	if err != nil {
		if os.IsNotExist(err) {
			return searchNpm(arg...)
		} else {
			return []string{}, err
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
	if _, err := dec.Token(); err != nil {
		return matches, err
	}

	return matches, nil
}

func searchNpm(arg ...string) ([]string, error) {
	cmd := exec.Command("npm", append([]string{"--no-color", "search"}, arg...)...)
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
		if !strings.HasPrefix(line, "NAME") {
			if index := strings.Index(line, " "); index > 0 {
				match := line[0:index]
				log.Printf("[npm] %s", match)
				matches = append(matches, match)
			}
		}
	}
	if scanner.Err() != nil {
		return matches, nil
	}
	//	stderr
	scanner = bufio.NewScanner(stderr)
	for scanner.Scan() {
		line := scanner.Text()
		log.Printf("[npm] stderr: %s", line)
	}
	if scanner.Err() != nil {
		return matches, nil
	}
	if err := cmd.Wait(); err != nil {
		log.Printf("[npm] exit status %d", cmd.ProcessState.Sys())
		return matches, err
	}
	return matches, nil
}

func (p *NpmProvider) Contains(arg ...string) ([]string, error) {
	lines, err := p.Search(arg...)
	matches := []string{}
	if err != nil {
		return matches, err
	} else {
		for _, found := range lines {
			if found == arg[0] {
				log.Printf("[npm] contains: %s", found)
				matches = append(matches, found)
			}
		}
		return matches, nil
	}
}

func (p *NpmProvider) Install(arg ...string) ([]string, error) {
	cmd := exec.Command("npm", append([]string{"--no-color", "install"}, arg...)...)
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
	cmd := exec.Command("npm", "--no-color", "config", "get", "cache")
	log.Printf("execute: %v", cmd.Args)
	out, err := cmd.CombinedOutput()
	outStr := string(out)
	outStr = strings.Trim(outStr, "\n\r")
	log.Printf("[npm] cache [%s]\n", outStr)
	return outStr, err
}
