package variable

import (
	"encoding/json"
	"testing"

	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestReplace(t *testing.T) {
	type args struct {
		req  any
		resp any
		vars []model.BodyVariable
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]any
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				req: map[string]any{
					"other":    3,
					"new_path": "this is what it should say",
					"others":   true,
				},
				resp: map[string]any{
					"replace": "nothing",
					"hello":   42.0,
				},
				vars: []model.BodyVariable{
					{
						Path: "replace",
						Func: "path",
						VariableParams: model.VariableParams{
							Args: []string{"new_path"},
						},
					},
				},
			},
			want: map[string]any{
				"replace": "this is what it should say",
				"hello":   42.0,
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
			resp, err := json.Marshal(tt.args.resp)
			if err != nil {
				t.Errorf("response encoding error %v", err)
				return
			}

			newResp, err := Replace(string(req), string(resp), tt.args.vars)
			switch {
			case (err != nil) != tt.wantErr:
				t.Errorf("Replace() error = %v, wantErr %v", err, tt.wantErr)
				return
			case err != nil:
			default:
				var got map[string]any
				if err := json.Unmarshal([]byte(newResp), &got); err != nil {
					t.Errorf("new response encoding %v", err)
					return
				}
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
