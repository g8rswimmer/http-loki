package validate

import (
	"testing"

	"github.com/g8rswimmer/http-loki/internal/model"
)

func TestIgnore(t *testing.T) {
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
			name: "success prefix",
			args: args{
				value: "prefix something",
				params: model.VariableParams{
					Prefix: "prefix ",
				},
			},
			wantErr: false,
		},
		{
			name: "fail prefix",
			args: args{
				value: "prefix something",
				params: model.VariableParams{
					Prefix: "error ",
				},
			},
			wantErr: true,
		},
		{
			name: "success sufix",
			args: args{
				value: "prefix something suffix",
				params: model.VariableParams{
					Suffix: "suffix",
				},
			},
			wantErr: false,
		},
		{
			name: "fail sufix",
			args: args{
				value: "prefix something suffix",
				params: model.VariableParams{
					Suffix: "error",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Ignore(tt.args.value, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("Ignore() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
