package body

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/g8rswimmer/http-loki/internal/httpx"
	"github.com/g8rswimmer/http-loki/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	type args struct {
		req  *httpx.Request
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
			name: "body success",
			args: args{
				req: &httpx.Request{
					Body: map[string]any{
						"other":    3,
						"new_path": "this is what it should say",
						"others":   true,
					},
				},
				resp: map[string]any{
					"replace": "nothing",
					"hello":   42.0,
				},
				bv: model.BodyVariable{
					Path: "replace",
					VariableParams: model.VariableParams{
						Args: []string{reqBodyValue, "new_path"},
					},
				},
			},
			want: map[string]any{
				"replace": "this is what it should say",
				"hello":   42.0,
			},
			wantErr: false,
		},
		{
			name: "query success",
			args: args{
				req: &httpx.Request{
					Body: map[string]any{
						"other":  3,
						"others": true,
					},
					QueryParameters: url.Values{
						"new_path": []string{"this is what it should say"},
					},
				},
				resp: map[string]any{
					"replace": "nothing",
					"hello":   42.0,
				},
				bv: model.BodyVariable{
					Path: "replace",
					VariableParams: model.VariableParams{
						Args: []string{reqQueryValue, "new_path"},
					},
				},
			},
			want: map[string]any{
				"replace": "this is what it should say",
				"hello":   42.0,
			},
			wantErr: false,
		},
		{
			name: "body success number",
			args: args{
				req: &httpx.Request{
					Body: map[string]any{
						"other":    3,
						"new_path": 27.0,
						"others":   true,
					},
				},
				resp: map[string]any{
					"replace": "nothing",
					"hello":   42.0,
				},
				bv: model.BodyVariable{
					Path: "replace",
					VariableParams: model.VariableParams{
						Args: []string{reqBodyValue, "new_path"},
					},
				},
			},
			want: map[string]any{
				"replace": 27.0,
				"hello":   42.0,
			},
			wantErr: false,
		},
		{
			name: "query success number",
			args: args{
				req: &httpx.Request{
					Body: map[string]any{
						"other":  3,
						"others": true,
					},
					QueryParameters: url.Values{
						"new_path": []string{"27.0"},
					},
				},
				resp: map[string]any{
					"replace": "nothing",
					"hello":   42.0,
				},
				bv: model.BodyVariable{
					Path: "replace",
					VariableParams: model.VariableParams{
						Args: []string{reqQueryValue, "new_path", floatType},
					},
				},
			},
			want: map[string]any{
				"replace": 27.0,
				"hello":   42.0,
			},
			wantErr: false,
		},
		{
			name: "body success prefix",
			args: args{
				req: &httpx.Request{
					Body: map[string]any{
						"other":    3,
						"new_path": "this is what it should say",
						"others":   true,
					},
				},
				resp: map[string]any{
					"replace": "nothing",
					"hello":   42.0,
				},
				bv: model.BodyVariable{
					Path: "replace",
					VariableParams: model.VariableParams{
						Args:   []string{reqBodyValue, "new_path"},
						Prefix: "Hi, ",
					},
				},
			},
			want: map[string]any{
				"replace": "Hi, this is what it should say",
				"hello":   42.0,
			},
			wantErr: false,
		},
		{
			name: "body success suffix",
			args: args{
				req: &httpx.Request{
					Body: map[string]any{
						"other":    3,
						"new_path": "this is what it should say",
						"others":   true,
					},
				},
				resp: map[string]any{
					"replace": "nothing",
					"hello":   42.0,
				},
				bv: model.BodyVariable{
					Path: "replace",
					VariableParams: model.VariableParams{
						Args:   []string{reqBodyValue, "new_path"},
						Suffix: ", Bye",
					},
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
				req: &httpx.Request{
					Body: map[string]any{
						"other":    3,
						"new_path": "this is what it should say",
						"others":   true,
					},
				},
				resp: map[string]any{
					"replace": "nothing",
					"hello":   42.0,
				},
				bv: model.BodyVariable{
					Path: "replace",
					VariableParams: model.VariableParams{
						Args: []string{},
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc, err := json.Marshal(tt.args.req.Body)
			if err != nil {
				t.Errorf("request encoding error %v", err)
				return
			}
			tt.args.req.EncodedBody = string(enc)

			resp, err := json.Marshal(tt.args.resp)
			if err != nil {
				t.Errorf("response encoding error %v", err)
				return
			}

			newResp, err := Path(tt.args.req, string(resp), tt.args.bv)
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
				assert.Equal(t, tt.want, got)
			}

		})
	}
}
