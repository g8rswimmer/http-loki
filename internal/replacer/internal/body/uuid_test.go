package body

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/go-playground/validator/v10"
)

func TestUUID(t *testing.T) {
	type args struct {
		resp any
		bv   model.BodyVariable
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				resp: map[string]any{
					"id": 42,
				},
				bv: model.BodyVariable{
					Path: "id",
				},
			},
			want:    "id",
			wantErr: false,
		},
		{
			name: "success with prefix",
			args: args{
				resp: map[string]any{
					"id": 42,
				},
				bv: model.BodyVariable{
					Path: "id",
					VariableParams: model.VariableParams{
						Prefix: "my prefix",
					},
				},
			},
			want:    "id",
			wantErr: false,
		},
		{
			name: "success with suffic",
			args: args{
				resp: map[string]any{
					"id": 42,
				},
				bv: model.BodyVariable{
					Path: "id",
					VariableParams: model.VariableParams{
						Suffix: "my suffix",
					},
				},
			},
			want:    "id",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := json.Marshal(tt.args.resp)
			if err != nil {
				t.Errorf("response encoding error %v", err)
				return
			}
			newResp, err := UUID(nil, string(resp), tt.args.bv)
			if (err != nil) != tt.wantErr {
				t.Errorf("UUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			var got map[string]any
			if err := json.Unmarshal([]byte(newResp), &got); err != nil {
				t.Errorf("new response encoding %v", err)
				return
			}
			v, ok := got[tt.want].(string)
			if !ok {
				t.Errorf("new response uuid assert %T", got[tt.want])
				return
			}
			switch {
			case len(tt.args.bv.Prefix) == 0:
			case !strings.HasPrefix(v, tt.args.bv.Prefix):
				t.Errorf("new response value does not have prefix %s %s", tt.args.bv.Prefix, v)
				return
			default:
				v = strings.TrimPrefix(v, tt.args.bv.Prefix)
			}
			switch {
			case len(tt.args.bv.Suffix) == 0:
			case !strings.HasSuffix(v, tt.args.bv.Suffix):
				t.Errorf("new response value does not have suffix %s %s", tt.args.bv.Suffix, v)
				return
			default:
				v = strings.TrimSuffix(v, tt.args.bv.Suffix)
			}

			if err := validator.New().Var(v, "uuid4"); err != nil {
				t.Errorf("new response uuid validation error %s", v)
			}
		})
	}
}
