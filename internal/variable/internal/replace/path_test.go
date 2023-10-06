package replace

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/g8rswimmer/http-loki/internal/model"
)

func TestPath(t *testing.T) {
	type args struct {
		req  any
		resp any
		bv   model.BodyVariable
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
				bv: model.BodyVariable{
					Path: "replace",
					Args: []string{"new_path"},
				},
			},
			want: map[string]any{
				"replace": "this is what it should say",
				"hello":   42.0,
			},
			wantErr: false,
		},
		{
			name: "success number",
			args: args{
				req: map[string]any{
					"other":    3,
					"new_path": 27.0,
					"others":   true,
				},
				resp: map[string]any{
					"replace": "nothing",
					"hello":   42.0,
				},
				bv: model.BodyVariable{
					Path: "replace",
					Args: []string{"new_path"},
				},
			},
			want: map[string]any{
				"replace": 27.0,
				"hello":   42.0,
			},
			wantErr: false,
		},
		{
			name: "success prefix",
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
				bv: model.BodyVariable{
					Path:   "replace",
					Args:   []string{"new_path"},
					Prefix: "Hi, ",
				},
			},
			want: map[string]any{
				"replace": "Hi, this is what it should say",
				"hello":   42.0,
			},
			wantErr: false,
		},
		{
			name: "success suffic",
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
				bv: model.BodyVariable{
					Path:   "replace",
					Args:   []string{"new_path"},
					Suffix: ", Bye",
				},
			},
			want: map[string]any{
				"replace": "this is what it should say, Bye",
				"hello":   42.0,
			},
			wantErr: false,
		},
		{
			name: "no args",
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
				bv: model.BodyVariable{
					Path: "replace",
					Args: []string{},
				},
			},
			want:    nil,
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
			resp, err := json.Marshal(tt.args.resp)
			if err != nil {
				t.Errorf("response encoding error %v", err)
				return
			}
			newResp, err := Path(string(req), string(resp), tt.args.bv)
			switch {
			case (err != nil) != tt.wantErr:
				t.Errorf("Path() error = %v, wantErr %v", err, tt.wantErr)
				return
			case err != nil:
			default:
				var got map[string]any
				if err := json.Unmarshal([]byte(newResp), &got); err != nil {
					t.Errorf("new response encoding %v", err)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Path() = %v, want %v", got, tt.want)
				}
			}

		})
	}
}
