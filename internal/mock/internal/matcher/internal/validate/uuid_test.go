package validate

import (
	"testing"

	"github.com/g8rswimmer/http-loki/internal/model"
)

func TestUUID(t *testing.T) {
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
				value:  "b2b7fa03-7972-4910-a13e-60b9d63c8dcf",
				params: model.VariableParams{},
			},
			wantErr: false,
		},
		{
			name: "success prefix",
			args: args{
				value: "prefix|b2b7fa03-7972-4910-a13e-60b9d63c8dcf",
				params: model.VariableParams{
					Prefix: "prefix|",
				},
			},
			wantErr: false,
		},
		{
			name: "fail prefix",
			args: args{
				value: "error|b2b7fa03-7972-4910-a13e-60b9d63c8dcf",
				params: model.VariableParams{
					Prefix: "prefix|",
				},
			},
			wantErr: true,
		},
		{
			name: "success suffix",
			args: args{
				value: "b2b7fa03-7972-4910-a13e-60b9d63c8dcf::suffix",
				params: model.VariableParams{
					Suffix: "::suffix",
				},
			},
			wantErr: false,
		},
		{
			name: "fail suffix",
			args: args{
				value: "b2b7fa03-7972-4910-a13e-60b9d63c8dcf::error",
				params: model.VariableParams{
					Suffix: "::suffix",
				},
			},
			wantErr: true,
		},
		{
			name: "invalid",
			args: args{
				value:  "uuid",
				params: model.VariableParams{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UUID(tt.args.value, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("UUID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
