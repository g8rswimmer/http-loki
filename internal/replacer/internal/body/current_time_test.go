package body

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/g8rswimmer/http-loki/internal/model"
)

func TestCurrentTime(t *testing.T) {
	type args struct {
		resp any
		bv   model.BodyVariable
	}
	tests := []struct {
		name       string
		args       args
		wantLayout string
		wantKey    string
		wantErr    bool
	}{
		{
			name: "success RFC3339",
			args: args{
				resp: map[string]any{
					"created_at": "nothing",
					"hello":      42.0,
				},
				bv: model.BodyVariable{
					Path: "created_at",
					VariableParams: model.VariableParams{
						Args: []string{"RFC3339"},
					},
				},
			},
			wantKey:    "created_at",
			wantLayout: time.RFC3339,
			wantErr:    false,
		},
		{
			name: "success layout",
			args: args{
				resp: map[string]any{
					"created_at": "nothing",
					"hello":      42.0,
				},
				bv: model.BodyVariable{
					Path: "created_at",
					VariableParams: model.VariableParams{
						Args: []string{"02 Jan 06"},
					},
				},
			},
			wantKey:    "created_at",
			wantLayout: "02 Jan 06",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := json.Marshal(tt.args.resp)
			if err != nil {
				t.Errorf("response encoding error %v", err)
				return
			}
			newResp, err := CurrentTime(nil, string(resp), tt.args.bv)
			switch {
			case (err != nil) != tt.wantErr:
				t.Errorf("CurrentTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			case err != nil:
			default:
				var got map[string]any
				if err := json.Unmarshal([]byte(newResp), &got); err != nil {
					t.Errorf("new response encoding %v", err)
					return
				}
				field, ok := got[tt.wantKey]
				if !ok {
					t.Errorf("CurrentTime() response key not present %s", tt.wantKey)
					return
				}
				timestamp, ok := field.(string)
				if !ok {
					t.Errorf("CurrentTime() response field not string %T", field)
					return
				}
				ct, err := time.Parse(tt.wantLayout, timestamp)
				if err != nil {
					t.Errorf("CurrentTime() timestamp can't be parsed %s %v", timestamp, err)
					return
				}
				now := time.Now().Add(time.Millisecond)
				if ct.After(time.Now()) {
					t.Errorf("CurrentTime() timestamp can't after curent time %v %v", ct, now)
					return
				}
			}
		})
	}
}
