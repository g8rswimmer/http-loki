package query

import (
	"testing"

	"github.com/g8rswimmer/http-loki/internal/model"
)

func TestIgnore(t *testing.T) {
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
			name: "success prefix",
			args: args{
				value: "prefix something",
				qv: model.QueryVariable{
					VariableParams: model.VariableParams{
						Prefix: "prefix ",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "success sufix",
			args: args{
				value: "prefix something suffix",
				qv: model.QueryVariable{
					VariableParams: model.VariableParams{
						Suffix: "suffix",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ignore(tt.args.value, tt.args.qv); (err != nil) != tt.wantErr {
				t.Errorf("Ignore() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
