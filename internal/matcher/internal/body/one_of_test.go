package body

import (
	"testing"

	"github.com/g8rswimmer/http-loki/internal/model"
)

func Test_oneOf(t *testing.T) {
	type args struct {
		value string
		bv    model.BodyVariable
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
				bv: model.BodyVariable{
					VariableParams: model.VariableParams{
						Args: []string{"one", "two", "three"},
					},
					Path: "/",
					Func: "oneOf",
				},
			},
			wantErr: false,
		},
		{
			name: "not found",
			args: args{
				value: "four",
				bv: model.BodyVariable{
					VariableParams: model.VariableParams{
						Args: []string{"one", "two", "three"},
					},
					Path: "/",
					Func: "oneOf",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := oneOf(tt.args.value, tt.args.bv); (err != nil) != tt.wantErr {
				t.Errorf("oneOf() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
