package validate

import (
	"testing"

	"github.com/g8rswimmer/http-loki/internal/model"
)

func TestRegEx(t *testing.T) {
	type args struct {
		value  string
		params model.VariableParams
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
				params: model.VariableParams{
					Args: []string{"p([a-z]+)ch"},
				},
			},
			wantErr: false,
		},
		{
			name: "success prefix",
			args: args{
				value: "nope peach",
				params: model.VariableParams{
					Args:   []string{"p([a-z]+)ch"},
					Prefix: "nope ",
				},
			},
			wantErr: false,
		},
		{
			name: "fail prefix",
			args: args{
				value: "nope peach",
				params: model.VariableParams{
					Args:   []string{"p([a-z]+)ch"},
					Prefix: "error ",
				},
			},
			wantErr: true,
		},
		{
			name: "success suffix",
			args: args{
				value: "peach after",
				params: model.VariableParams{
					Args:   []string{"p([a-z]+)ch"},
					Suffix: " after",
				},
			},
			wantErr: false,
		},
		{
			name: "fail suffix",
			args: args{
				value: "peach after",
				params: model.VariableParams{
					Args:   []string{"p([a-z]+)ch"},
					Suffix: " error",
				},
			},
			wantErr: true,
		},
		{
			name: "fail",
			args: args{
				value: "nope",
				params: model.VariableParams{
					Args: []string{"p([a-z]+)ch"},
				},
			},
			wantErr: true,
		},
		{
			name: "no args",
			args: args{
				value: "peach",
				params: model.VariableParams{
					Args: []string{},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RegEx(tt.args.value, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("RegEx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
