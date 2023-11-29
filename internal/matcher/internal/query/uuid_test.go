package query

import (
	"testing"

	"github.com/g8rswimmer/http-loki/internal/model"
)

func TestUUID(t *testing.T) {
	type args struct {
		value string
		bv    model.QueryVariable
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				value: "b2b7fa03-7972-4910-a13e-60b9d63c8dcf",
				bv:    model.QueryVariable{},
			},
			wantErr: false,
		},
		{
			name: "success prefix",
			args: args{
				value: "prefix|b2b7fa03-7972-4910-a13e-60b9d63c8dcf",
				bv: model.QueryVariable{
					VariableParams: model.VariableParams{
						Prefix: "prefix|",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "success suffix",
			args: args{
				value: "b2b7fa03-7972-4910-a13e-60b9d63c8dcf::suffix",
				bv: model.QueryVariable{
					VariableParams: model.VariableParams{
						Suffix: "::suffix",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				value: "uuid",
				bv:    model.QueryVariable{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := uuid(tt.args.value, tt.args.bv); (err != nil) != tt.wantErr {
				t.Errorf("UUID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
