package validate

import (
	"encoding/json"
	"testing"
)

func TestIntRange(t *testing.T) {
	type args struct {
		req  any
		path string
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				req: map[string]any{
					"id":     "hi",
					"number": 6,
				},
				path: "number",
				args: []string{"-10", "10"},
			},
			wantErr: false,
		},
		{
			name: "range fail",
			args: args{
				req: map[string]any{
					"id":     "hi",
					"number": 100,
				},
				path: "number",
				args: []string{"-10", "10"},
			},
			wantErr: true,
		},
		{
			name: "arg length fail",
			args: args{
				req: map[string]any{
					"id":     "hi",
					"number": 6,
				},
				path: "number",
				args: []string{"-10"},
			},
			wantErr: true,
		},
		{
			name: "value type fail",
			args: args{
				req: map[string]any{
					"id":     "hi",
					"number": "6",
				},
				path: "number",
				args: []string{"-10", "10"},
			},
			wantErr: true,
		},
		{
			name: "arg NaN less",
			args: args{
				req: map[string]any{
					"id":     "hi",
					"number": 6,
				},
				path: "number",
				args: []string{"hi", "10"},
			},
			wantErr: true,
		},
		{
			name: "arg NaN great",
			args: args{
				req: map[string]any{
					"id":     "hi",
					"number": 6,
				},
				path: "number",
				args: []string{"-10", "bye"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := json.Marshal(tt.args.req)
			if err != nil {
				t.Errorf("request encoding error %v", err)
				return
			}
			if err := IntRange(string(req), tt.args.path, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("IntRange() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
