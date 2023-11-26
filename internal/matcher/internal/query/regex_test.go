package query

import (
	"testing"

	"github.com/g8rswimmer/http-loki/internal/model"
)

func TestRegEx(t *testing.T) {
	type args struct {
		value string
		qv    model.QueryVariable
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				value: "peach",
				qv: model.QueryVariable{
					VariableParams: model.VariableParams{
						Args: []string{"p([a-z]+)ch"},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "success prefix",
			args: args{
				value: "nope peach",
				qv: model.QueryVariable{
					VariableParams: model.VariableParams{
						Args:   []string{"p([a-z]+)ch"},
						Prefix: "nope ",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "success suffix",
			args: args{
				value: "peach after",
				qv: model.QueryVariable{
					VariableParams: model.VariableParams{
						Args:   []string{"p([a-z]+)ch"},
						Suffix: " after",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				value: "nope",
				qv: model.QueryVariable{
					VariableParams: model.VariableParams{
						Args: []string{"p([a-z]+)ch"},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "no args",
			args: args{
				value: "peach",
				qv: model.QueryVariable{
					VariableParams: model.VariableParams{
						Args: []string{},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := regex(tt.args.value, tt.args.qv); (err != nil) != tt.wantErr {
				t.Errorf("RegEx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
