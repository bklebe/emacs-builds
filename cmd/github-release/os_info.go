package main

import (
	"os/exec"
	"strings"
)

type OSInfo struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Arch    string `yaml:"arch"`
}

func NewOSInfo() (*OSInfo, error) {
	version, err := exec.Command("sw_vers", "-productVersion").CombinedOutput()
	if err != nil {
		return nil, err
	}

	arch, err := exec.Command("uname", "-m").CombinedOutput()
	if err != nil {
		return nil, err
	}

	return &OSInfo{
		Name:    "macOS",
		Version: strings.TrimSpace(string(version)),
		Arch:    strings.TrimSpace(string(arch)),
	}, nil
}

func (s *OSInfo) ShortVersion() string {
	parts := strings.Split(s.Version, ".")
	max := len(parts)
	if max > 2 {
		max = 2
	}

	return strings.Join(parts[0:max], ".")
}
