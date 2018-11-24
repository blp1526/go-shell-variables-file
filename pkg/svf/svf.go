package svf

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// ShellVariablesFile represents file for shell variables.
//
// e.g. /etc/os-release, /etc/lsb-release
type ShellVariablesFile struct {
	path  string
	items map[string]string
}

// New initialize *ShellVariablesFile by path.
func New(path string) *ShellVariablesFile {
	s := &ShellVariablesFile{
		path: path,
	}

	return s
}

// ReadFile initialize *ShellVariablesFile by path and set items.
func ReadFile(path string) (*ShellVariablesFile, error) {
	s := New(path)
	s.items = map[string]string{}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "#") {
			continue
		}

		if !strings.Contains(line, "=") {
			return nil, fmt.Errorf("line: '%s' has no '=' separator", line)
		}

		kv := strings.SplitN(line, "=", 2)
		key := kv[0]
		value := kv[1]

		s.items[key] = value
	}

	return s, nil
}

// GetRaw gets value by name.
func (s *ShellVariablesFile) GetRaw(name string) (string, error) {
	value, ok := s.items[name]
	if !ok {
		return "", fmt.Errorf("name: '%s' is not present", name)
	}

	return value, nil
}

// Get gets quote trimmed value by name.
func (s *ShellVariablesFile) Get(name string) (string, error) {
	value, err := s.GetRaw(name)
	if err != nil {
		return "", err
	}

	doubleQuote := `"`
	if strings.HasPrefix(value, doubleQuote) && strings.HasSuffix(value, doubleQuote) {
		value = strings.TrimPrefix(value, doubleQuote)
		value = strings.TrimSuffix(value, doubleQuote)
		return value, nil
	}

	singleQuote := `'`
	if strings.HasPrefix(value, singleQuote) && strings.HasSuffix(value, singleQuote) {
		value = strings.TrimPrefix(value, singleQuote)
		value = strings.TrimSuffix(value, singleQuote)
		return value, nil
	}

	return value, nil
}

// Keys returns items keys.
func (s ShellVariablesFile) Keys() []string {
	keys := []string{}

	for key := range s.items {
		keys = append(keys, key)
	}

	return keys
}
