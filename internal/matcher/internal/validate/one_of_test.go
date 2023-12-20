package validate

import (
	"testing"

	"github.com/g8rswimmer/http-loki/internal/model"
)

func TestOneOf(t *testing.T) {
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
			name: "found",
			args: args{
				value: "three",
				params: model.VariableParams{
					Args: []string{"one", "two", "three"},
				},
			},
			wantErr: false,
		},
		{
			name: "not found",
			args: args{
				value: "four",
				params: model.VariableParams{
					Args: []string{"one", "two", "three"},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := OneOf(tt.args.value, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("OneOf() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
