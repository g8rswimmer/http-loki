package validate

import (
	"encoding/json"
	"testing"
)

func TestUUID(t *testing.T) {
	type args struct {
		req  any
		path string
		in2  []string
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
					"id": "uuid",
				},
				path: "id",
			},
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				req: map[string]any{
					"id": "uuid",
				},
				path: "id",
			},
			wantErr: true,
		},
		{
			name: "not a string",
			args: args{
				req: map[string]any{
					"id": 77,
				},
				path: "id",
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
			if err := UUID(string(req), tt.args.path, tt.args.in2); (err != nil) != tt.wantErr {
				t.Errorf("UUID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
