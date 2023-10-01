package validate

import (
	"encoding/json"
	"testing"
)

func TestRegEx(t *testing.T) {
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
					"reg_it": "peach",
				},
				path: "reg_it",
				args: []string{"p([a-z]+)ch"},
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				req: map[string]any{
					"reg_it": "nope",
				},
				path: "reg_it",
				args: []string{"p([a-z]+)ch"},
			},
			wantErr: true,
		},
		{
			name: "no args",
			args: args{
				req: map[string]any{
					"reg_it": "peach",
				},
				path: "reg_it",
				args: []string{},
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
			if err := RegEx(string(req), tt.args.path, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("RegEx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
