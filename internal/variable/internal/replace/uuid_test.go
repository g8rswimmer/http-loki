package replace

import (
	"encoding/json"
	"testing"

	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/go-playground/validator/v10"
)

func TestUUID(t *testing.T) {
	type args struct {
		in0  string
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := json.Marshal(tt.args.resp)
			if err != nil {
				t.Errorf("response encoding error %v", err)
				return
			}
			newResp, err := UUID(tt.args.in0, string(resp), tt.args.bv)
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
			if err := validator.New().Var(v, "uuid4"); err != nil {
				t.Errorf("new response uuid validation error %s", v)
			}
		})
	}
}
