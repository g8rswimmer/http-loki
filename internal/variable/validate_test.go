package variable

import (
	"encoding/json"
	"testing"

	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	type args struct {
		req  any
		vars []model.BodyVariable
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				req: map[string]any{
					"id":     "hi",
					"number": 6,
				},
				vars: []model.BodyVariable{
					{
						Path: "number",
						Args: []string{"-10", "10"},
						Func: "intRange",
					},
				},
			},
			want: map[string]any{
				"id":     "hi",
				"number": validationValue,
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

			newResp, err := Validate(string(req), tt.args.vars)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
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
