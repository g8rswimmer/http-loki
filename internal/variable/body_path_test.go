package variable

import (
	"sort"
	"testing"

	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestBodyPaths(t *testing.T) {
	type args struct {
		body     any
		currPath string
		paths    []model.BodyVariable
	}
	tests := []struct {
		name string
		args args
		want []model.BodyVariable
	}{
		{
			name: "success basic",
			args: args{
				body: map[string]any{
					"uuid_var":      "{{ uuid }}",
					"var_with_args": "{{ args:arg1|arg2 }}",
					"nested": map[string]any{
						"nest_var": "{{ nested }}",
					},
				},
				currPath: "",
				paths:    []model.BodyVariable{},
			},
			want: []model.BodyVariable{
				{
					Path: "uuid_var",
					Func: "uuid",
					Args: []string{},
				},
				{
					Path: "var_with_args",
					Func: "args",
					Args: []string{"arg1", "arg2"},
				},
				{
					Path: "nested.nest_var",
					Func: "nested",
					Args: []string{},
				},
			},
		},
		{
			name: "success prefix suffix",
			args: args{
				body: map[string]any{
					"uuid_var":      "prefix{{ uuid }}",
					"var_with_args": "{{ args:arg1|arg2 }}",
					"nested": map[string]any{
						"nest_var": "prefix  {{ nested }}  suffix ",
					},
				},
				currPath: "",
				paths:    []model.BodyVariable{},
			},
			want: []model.BodyVariable{
				{
					Path:   "uuid_var",
					Func:   "uuid",
					Args:   []string{},
					Prefix: "prefix",
				},
				{
					Path: "var_with_args",
					Func: "args",
					Args: []string{"arg1", "arg2"},
				},
				{
					Path:   "nested.nest_var",
					Func:   "nested",
					Args:   []string{},
					Prefix: "prefix  ",
					Suffix: "  suffix ",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BodyPaths(tt.args.body, tt.args.currPath, tt.args.paths)
			sort.Slice(got, func(i, j int) bool {
				return got[i].Path < got[j].Path
			})
			sort.Slice(tt.want, func(i, j int) bool {
				return tt.want[i].Path < tt.want[j].Path
			})
			assert.Equal(t, tt.want, got)
		})
	}
}
