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
