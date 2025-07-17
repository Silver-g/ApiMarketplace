package boundary

import (
	"testing"
)

type regUserCase struct {
	name    string
	input   UserRequest
	wantErr error
}

func TestUserValidate(t *testing.T) {
	tests := []regUserCase{
		{
			name:    "ValidInput",
			input:   UserRequest{Username: "_validuser_", Password: "password123"},
			wantErr: nil,
		},
		{
			name:    "EmptyInput",
			input:   UserRequest{Username: "", Password: ""},
			wantErr: ErrEmptyLoginPassword,
		},
		{
			name:    "ShortUsername",
			input:   UserRequest{Username: "ab", Password: "password123"},
			wantErr: ErrUsernameLength,
		},
		{
			name:    "LongUsername",
			input:   UserRequest{Username: "averyveryverylongusername", Password: "password123"},
			wantErr: ErrUsernameLength,
		},
		{
			name:    "ShortPassword",
			input:   UserRequest{Username: "validuser", Password: "short"},
			wantErr: ErrPasswordLength,
		},
		{
			name:    "OnlySpace",
			input:   UserRequest{Username: "___   ___", Password: "password123"},
			wantErr: ErrUsernameSpaces,
		},
		{
			name:    "ProhibitedCharacters",
			input:   UserRequest{Username: "bad;user", Password: "password123"},
			wantErr: ErrUsernameProhibited,
		},
		{
			name:    "InvalidCharacters",
			input:   UserRequest{Username: "имяПользователя", Password: "password123"},
			wantErr: ErrUsernameInvalidChars,
		},
		{
			name:    "SpaceInUsername",
			input:   UserRequest{Username: "user name", Password: "password123"},
			wantErr: ErrUsernameInvalidChars,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := UserValidate(tc.input)
			if err != tc.wantErr {
				t.Errorf("UserValidate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
