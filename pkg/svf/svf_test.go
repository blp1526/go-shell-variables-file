package svf

import (
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
