package body

import (
	"encoding/json"
	"testing"

	"github.com/g8rswimmer/http-loki/internal/model"
)

func TestIgnore(t *testing.T) {
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
			name: "success prefix",
			args: args{
				req: map[string]any{
					"ignore_this": "prefix something",
				},
				bv: model.BodyVariable{
					Path: "ignore_this",
					VariableParams: model.VariableParams{
						Prefix: "prefix ",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "success sufix",
			args: args{
				req: map[string]any{
					"ignore_this": "prefix something suffix",
				},
				bv: model.BodyVariable{
					Path: "ignore_this",
					VariableParams: model.VariableParams{
						Suffix: "suffix",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := json.Marshal(tt.args.req)
			if err != nil {
				t.Errorf("request encoding error %v", err)
				return
			}

			if err := Ignore(string(req), tt.args.bv); (err != nil) != tt.wantErr {
				t.Errorf("Ignore() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
