package validate

import (
	"encoding/json"
	"testing"

	"github.com/g8rswimmer/http-loki/internal/model"
)

func TestIntRange(t *testing.T) {
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
					"id":     "hi",
					"number": 6,
				},
				bv: model.BodyVariable{
					Path: "number",
					Args: []string{"-10", "10"},
				},
			},
			wantErr: false,
		},
		{
			name: "success array",
			args: args{
				req: map[string]any{
					"id":      "hi",
					"numbers": []int{6, -2},
				},
				bv: model.BodyVariable{
					Path: "numbers",
					Args: []string{"-10", "10"},
				},
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
				bv: model.BodyVariable{
					Path: "number",
					Args: []string{"-10", "10"},
				},
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
				bv: model.BodyVariable{
					Path: "number",
					Args: []string{"-10"},
				},
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
				bv: model.BodyVariable{
					Path: "number",
					Args: []string{"-10", "10"},
				},
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
				bv: model.BodyVariable{
					Path: "number",
					Args: []string{"hi", "10"},
				},
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
				bv: model.BodyVariable{
					Path: "number",
					Args: []string{"-10", "bye"},
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
			if err := IntRange(string(req), tt.args.bv); (err != nil) != tt.wantErr {
				t.Errorf("IntRange() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
