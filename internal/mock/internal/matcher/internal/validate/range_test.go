package validate

import (
	"testing"

	"github.com/g8rswimmer/http-loki/internal/model"
)

func TestIntRange(t *testing.T) {
	type args struct {
		value  int
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
				value: 6,
				params: model.VariableParams{
					Args: []string{"-10", "10"},
				},
			},
			wantErr: false,
		},
		{
			name: "range fail",
			args: args{
				value: 100,
				params: model.VariableParams{
					Args: []string{"-10", "10"},
				},
			},
			wantErr: true,
		},
		{
			name: "arg length fail",
			args: args{
				value: 6,
				params: model.VariableParams{
					Args: []string{"-10"},
				},
			},
			wantErr: true,
		},
		{
			name: "arg NaN less",
			args: args{
				value: 6,
				params: model.VariableParams{
					Args: []string{"hi", "10"},
				},
			},
			wantErr: true,
		},
		{
			name: "arg NaN great",
			args: args{
				value: 6,
				params: model.VariableParams{
					Args: []string{"-10", "bye"},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IntRange(tt.args.value, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("IntRange() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
