package validate

import (
	"encoding/json"
	"testing"

	"github.com/g8rswimmer/http-loki/internal/model"
)

func TestRegEx(t *testing.T) {
	type args struct {
		req any
		bv  model.BodyVariable
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
				bv: model.BodyVariable{
					Path: "reg_it",
					Args: []string{"p([a-z]+)ch"},
				},
			},
			wantErr: false,
		},
		{
			name: "success prefix",
			args: args{
				req: map[string]any{
					"reg_it": "nope peach",
				},
				bv: model.BodyVariable{
					Path:   "reg_it",
					Args:   []string{"p([a-z]+)ch"},
					Prefix: "nope ",
				},
			},
			wantErr: false,
		},
		{
			name: "success suffix",
			args: args{
				req: map[string]any{
					"reg_it": "peach after",
				},
				bv: model.BodyVariable{
					Path:   "reg_it",
					Args:   []string{"p([a-z]+)ch"},
					Suffix: " after",
				},
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				req: map[string]any{
					"reg_it": "nope",
				},
				bv: model.BodyVariable{
					Path: "reg_it",
					Args: []string{"p([a-z]+)ch"},
				},
			},
			wantErr: true,
		},
		{
			name: "no args",
			args: args{
				req: map[string]any{
					"reg_it": "peach",
				},
				bv: model.BodyVariable{
					Path: "reg_it",
					Args: []string{},
				},
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
			if err := RegEx(string(req), tt.args.bv); (err != nil) != tt.wantErr {
				t.Errorf("RegEx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
