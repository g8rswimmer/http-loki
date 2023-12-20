package query

import (
	"testing"

	"github.com/g8rswimmer/http-loki/internal/model"
)

func Test_oneOf(t *testing.T) {
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
			name: "found",
			args: args{
				value: "one",
				qv: model.QueryVariable{
					VariableParams: model.VariableParams{
						Args: []string{"three", "two", "one"},
					},
					Func: "oneOf",
				},
			},
			wantErr: false,
		},
		{
			name: "not found",
			args: args{
				value: "four",
				qv: model.QueryVariable{
					VariableParams: model.VariableParams{
						Args: []string{"three", "two", "one"},
					},
					Func: "oneOf",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := oneOf(tt.args.value, tt.args.qv); (err != nil) != tt.wantErr {
				t.Errorf("oneOf() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
