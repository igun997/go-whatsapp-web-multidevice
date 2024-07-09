package validations

import (
	"context"
	"github.com/stretchr/testify/assert"
	domainUser "main/domains/user"
	pkgError "main/pkg/error"
	"testing"
)

func TestValidateUserAvatar(t *testing.T) {
	type args struct {
		request domainUser.AvatarRequest
	}
	tests := []struct {
		name string
		args args
		err  any
	}{
		{
			name: "should success",
			args: args{request: domainUser.AvatarRequest{
				Phone:       "1728937129312@s.whatsapp.net",
				IsPreview:   false,
				IsCommunity: false,
			}},
			err: nil,
		},
		{
			name: "should error with empty phone",
			args: args{request: domainUser.AvatarRequest{
				Phone:       "",
				IsPreview:   false,
				IsCommunity: false,
			}},
			err: pkgError.ValidationError("phone: cannot be blank."),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUserAvatar(context.Background(), tt.args.request)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestValidateUserInfo(t *testing.T) {
	type args struct {
		request domainUser.InfoRequest
	}
	tests := []struct {
		name string
		args args
		err  any
	}{
		{
			name: "should success",
			args: args{request: domainUser.InfoRequest{
				Phone: "1728937129312@s.whatsapp.net",
			}},
			err: nil,
		},
		{
			name: "should error with empty phone",
			args: args{request: domainUser.InfoRequest{
				Phone: "",
			}},
			err: pkgError.ValidationError("phone: cannot be blank."),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUserInfo(context.Background(), tt.args.request)
			assert.Equal(t, tt.err, err)
		})
	}
}
