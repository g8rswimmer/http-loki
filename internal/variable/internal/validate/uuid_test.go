package validate

import (
	"encoding/json"
	"testing"

	"github.com/g8rswimmer/http-loki/internal/model"
)

func TestUUID(t *testing.T) {
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
					"id": "b2b7fa03-7972-4910-a13e-60b9d63c8dcf",
				},
				bv: model.BodyVariable{
					Path: "id",
				},
			},
			wantErr: false,
		},
		{
			name: "success arrary",
			args: args{
				req: map[string]any{
					"id": []string{"b2b7fa03-7972-4910-a13e-60b9d63c8dcf", "a8e21b84-2160-46db-91dd-54af2b261643"},
				},
				bv: model.BodyVariable{
					Path: "id",
				},
			},
			wantErr: false,
		},
		{
			name: "success prefix",
			args: args{
				req: map[string]any{
					"id": "prefix|b2b7fa03-7972-4910-a13e-60b9d63c8dcf",
				},
				bv: model.BodyVariable{
					Path:   "id",
					Prefix: "prefix|",
				},
			},
			wantErr: false,
		},
		{
			name: "success suffix",
			args: args{
				req: map[string]any{
					"id": "b2b7fa03-7972-4910-a13e-60b9d63c8dcf::suffix",
				},
				bv: model.BodyVariable{
					Path:   "id",
					Suffix: "::suffix",
				},
			},
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				req: map[string]any{
					"id": "uuid",
				},
				bv: model.BodyVariable{
					Path: "id",
				},
			},
			wantErr: true,
		},
		{
			name: "not a string",
			args: args{
				req: map[string]any{
					"id": 77,
				},
				bv: model.BodyVariable{
					Path: "id",
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
			if err := UUID(string(req), tt.args.bv); (err != nil) != tt.wantErr {
				t.Errorf("UUID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
