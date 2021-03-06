package svf

import (
	"fmt"
	"io/ioutil"
	"sort"
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
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	s := New(path)
	err = s.setItems(string(b))
	if err != nil {
		return nil, err
	}

	return s, nil
}

// setItems sets items.
func (s *ShellVariablesFile) setItems(content string) error {
	s.items = map[string]string{}

	lines := strings.Split(content, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "#") {
			continue
		}

		if !strings.Contains(line, "=") {
			return fmt.Errorf("line: \"%s\" has no \"=\" separator", line)
		}

		kv := strings.SplitN(line, "=", 2)
		key := kv[0]
		value := kv[1]

		_, hasKey := s.items[key]
		if hasKey {
			return fmt.Errorf("key: \"%s\" is duplicated", key)
		}

		s.items[key] = value
	}

	return nil
}

// GetRawValue gets a value by a key.
func (s *ShellVariablesFile) GetRawValue(key string) (string, error) {
	value, ok := s.items[key]
	if !ok {
		return "", fmt.Errorf("key: '%s' is not present", key)
	}

	return value, nil
}

// GetValue gets a quote trimmed value by a key.
func (s *ShellVariablesFile) GetValue(key string) (string, error) {
	value, err := s.GetRawValue(key)
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

// IsValidKeys validates all given keys.
func (s *ShellVariablesFile) IsValidKeys(keys []string) error {
	errors := []string{}

	for _, key := range keys {
		_, invalid := s.GetRawValue(key)
		if invalid != nil {
			errors = append(errors, invalid.Error())
		}
	}

	if len(errors) != 0 {
		return fmt.Errorf("%s", strings.Join(errors, ", "))
	}

	return nil
}

// Keys returns items' keys.
func (s *ShellVariablesFile) Keys() []string {
	keys := []string{}

	for key := range s.items {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	return keys
}

// Values returns items' values.
func (s *ShellVariablesFile) Values(keys []string) (values []string, err error) {
	values = []string{}

	err = s.IsValidKeys(keys)
	if err != nil {
		return values, err
	}

	for _, key := range keys {
		value, _ := s.GetValue(key)
		values = append(values, value)
	}

	sort.Strings(values)
	return values, nil
}
