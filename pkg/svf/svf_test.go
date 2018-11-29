package svf

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		path string
		want *ShellVariablesFile
	}{
		{
			path: "/etc/os-release",
			want: &ShellVariablesFile{path: "/etc/os-release"},
		},
	}

	for _, tt := range tests {
		got := New(tt.path)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("tt.path: %v, tt.want: %v, got: %v", tt.path, tt.want, got)
		}
	}
}

func TestReadFile(t *testing.T) {
	tempDir, _ := ioutil.TempDir("", "")
	defer os.RemoveAll(tempDir)
	path := filepath.Join(tempDir, "test")

	tests := []struct {
		isFilePresent bool
		content       []byte
		want          *ShellVariablesFile
		err           bool
	}{
		{
			isFilePresent: false,
			content:       []byte(""),
			want:          nil,
			err:           true,
		},
		{
			isFilePresent: true,
			content:       []byte("FOO"),
			want:          nil,
			err:           true,
		},
		{
			isFilePresent: true,
			content:       []byte("FOO=BAR"),
			want: &ShellVariablesFile{
				path:  path,
				items: map[string]string{"FOO": "BAR"},
			},
			err: false,
		},
	}

	for _, tt := range tests {
		if tt.isFilePresent {
			ioutil.WriteFile(path, tt.content, 0644)
		}

		got, err := ReadFile(path)

		if tt.err && err == nil {
			t.Errorf("tt.isFilePresent: %v, tt.content: %s, tt.want: %v, tt.err: %v, got: %v, err: %v",
				tt.isFilePresent, tt.content, tt.want, tt.err, got, err)
		}

		if tt.err && err == nil {
			t.Errorf("tt.isFilePresent: %v, tt.content: %s, tt.want: %v, tt.err: %v, got: %v, err: %v",
				tt.isFilePresent, tt.content, tt.want, tt.err, got, err)
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("tt.isFilePresent: %v, tt.content: %s, tt.want: %v, tt.err: %v, got: %v, err: %v",
				tt.isFilePresent, tt.content, tt.want, tt.err, got, err)
		}
	}
}

func TestSetItems(t *testing.T) {
	tests := []struct {
		content string
		want    map[string]string
		err     bool
	}{
		{
			content: "",
			want:    map[string]string{},
			err:     false,
		},
		{
			content: "#",
			want:    map[string]string{},
			err:     false,
		},
		{
			content: " #",
			want:    map[string]string{},
			err:     false,
		},
		{
			content: "FOO",
			want:    map[string]string{},
			err:     true,
		},
		{
			content: "FOO=BAR\nFOO=BAR",
			want:    map[string]string{"FOO": "BAR"},
			err:     true,
		},
		{
			content: "FOO=BAR",
			want:    map[string]string{"FOO": "BAR"},
			err:     false,
		},
	}

	for _, tt := range tests {
		s := &ShellVariablesFile{}
		err := s.setItems(tt.content)

		if tt.err && err == nil {
			t.Errorf("tt.content: %v, tt.want: %v, s.items: %v, tt.err: %v, err: %v", tt.content, tt.want, s.items, tt.err, err)
		}

		if !tt.err && err != nil {
			t.Errorf("tt.content: %v, tt.want: %v, s.items: %v, tt.err: %v, err: %v", tt.content, tt.want, s.items, tt.err, err)
		}

		if !reflect.DeepEqual(s.items, tt.want) {
			t.Errorf("tt.content: %v, tt.want: %v, s.items: %v, tt.err: %v, err: %v", tt.content, tt.want, s.items, tt.err, err)
		}
	}
}

func TestGetRawValue(t *testing.T) {
	tests := []struct {
		items map[string]string
		key   string
		want  string
		err   bool
	}{
		{
			items: map[string]string{"foo": `"bar"`},
			key:   "foo",
			want:  `"bar"`,
			err:   false,
		},
		{
			items: map[string]string{"foo": `"bar"`},
			key:   "baz",
			want:  "",
			err:   true,
		},
	}

	for _, tt := range tests {
		s := &ShellVariablesFile{items: tt.items}
		got, err := s.GetRawValue(tt.key)

		if tt.err && err == nil {
			t.Errorf("items: %v, key: %v, tt.want: %v, tt.err: %v, got: %v, err: %v", tt.items, tt.key, tt.want, tt.err, got, err)
		}

		if !tt.err && err != nil {
			t.Errorf("items: %v, key: %v, tt.want: %v, tt.err: %v, got: %v, err: %v", tt.items, tt.key, tt.want, tt.err, got, err)
		}

		if got != tt.want {
			t.Errorf("items: %v, key: %v, tt.want: %v, tt.err: %v, got: %v, err: %v", tt.items, tt.key, tt.want, tt.err, got, err)
		}
	}
}

func TestGetValue(t *testing.T) {
	tests := []struct {
		items map[string]string
		key   string
		want  string
		err   bool
	}{
		{
			items: map[string]string{"foo": `"bar"`},
			key:   "foo",
			want:  "bar",
			err:   false,
		},
		{
			items: map[string]string{"foo": `'bar'`},
			key:   "foo",
			want:  "bar",
			err:   false,
		},
		{
			items: map[string]string{"foo": "bar"},
			key:   "foo",
			want:  "bar",
			err:   false,
		},
		{
			items: map[string]string{"foo": `"bar"`},
			key:   "baz",
			want:  "",
			err:   true,
		},
	}

	for _, tt := range tests {
		s := &ShellVariablesFile{items: tt.items}
		got, err := s.GetValue(tt.key)

		if tt.err && err == nil {
			t.Errorf("items: %v, key: %v, tt.want: %v, tt.err: %v, got: %v, err: %v", tt.items, tt.key, tt.want, tt.err, got, err)
		}

		if !tt.err && err != nil {
			t.Errorf("items: %v, key: %v, tt.want: %v, tt.err: %v, got: %v, err: %v", tt.items, tt.key, tt.want, tt.err, got, err)
		}

		if got != tt.want {
			t.Errorf("items: %v, key: %v, tt.want: %v, tt.err: %v, got: %v, err: %v", tt.items, tt.key, tt.want, tt.err, got, err)
		}
	}
}

func TestIsValidKeys(t *testing.T) {
	tests := []struct {
		items map[string]string
		keys  []string
		err   bool
	}{
		{
			items: map[string]string{"1": "A", "2": "B"},
			keys:  []string{"1"},
			err:   false,
		},
		{
			items: map[string]string{"1": "A", "2": "B"},
			keys:  []string{"3"},
			err:   true,
		},
	}

	for _, tt := range tests {
		s := &ShellVariablesFile{}
		s.items = tt.items
		err := s.IsValidKeys(tt.keys)

		if tt.err && err == nil {
			t.Errorf("tt.items: %v, tt.keys: %v, tt.err: %v, err: %v", tt.items, tt.keys, tt.err, err)
		}

		if !tt.err && err != nil {
			t.Errorf("tt.items: %v, tt.keys: %v, tt.err: %v, err: %v", tt.items, tt.keys, tt.err, err)
		}
	}
}

func TestKeys(t *testing.T) {
	tests := []struct {
		items map[string]string
		want  []string
	}{
		{
			items: map[string]string{"2": "B", "1": "A"},
			want:  []string{"1", "2"},
		},
	}

	for _, tt := range tests {
		s := &ShellVariablesFile{}
		s.items = tt.items
		got := s.Keys()

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("tt.items: %v, tt.want: %v, got: %v", tt.items, tt.want, got)
		}
	}
}

func TestValues(t *testing.T) {
	tests := []struct {
		keys []string
		want []string
		err  bool
	}{
		{
			keys: []string{"4", "1"},
			want: []string{},
			err:  true,
		},
		{
			keys: []string{"3", "1"},
			want: []string{"A", "C"},
			err:  false,
		},
	}

	for _, tt := range tests {
		s := &ShellVariablesFile{}
		s.items = map[string]string{"3": "C", "2": "B", "1": "A"}
		got, err := s.Values(tt.keys)

		if tt.err && err == nil {
			t.Errorf("tt.keys: %v, tt.want: %v, got: %v, tt.err: %v, err: %v", tt.keys, tt.want, got, tt.err, err)
		}

		if !tt.err && err != nil {
			t.Errorf("tt.keys: %v, tt.want: %v, got: %v, tt.err: %v, err: %v", tt.keys, tt.want, got, tt.err, err)
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("tt.keys: %v, tt.want: %v, got: %v, tt.err: %v, err: %v", tt.keys, tt.want, got, tt.err, err)
		}
	}
}
