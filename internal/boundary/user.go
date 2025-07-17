package boundary

import (
	"ApiMarketplace/internal/consts"
	"ApiMarketplace/internal/domain"
	"errors"
	"regexp"
)

type UserRequest struct {
	Username string `json:"login"`
	Password string `json:"password"`
}
type RegisterUserResponse struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
}
type LoginUserResponse struct {
	JwtToken string `json:"authtoken"`
}

// Кастомные ошибки
var (
	ErrEmptyLoginPassword   = errors.New(consts.ErrEmptyLoginPasswordMsg)
	ErrUsernameLength       = errors.New(consts.ErrUsernameLengthMsg)
	ErrPasswordLength       = errors.New(consts.ErrPasswordLengthMsg)
	ErrUsernameSpaces       = errors.New(consts.ErrUsernameSpacesMsg)
	ErrUsernameProhibited   = errors.New(consts.ErrUsernameProhibitedMsg)
	ErrUsernameInvalidChars = errors.New(consts.ErrUsernameInvalidCharsMsg)
)

// регулярка разрешены латинские буквы цифры подчеркивания
var registerRegexp = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)

//var registerRegexp = regexp.MustCompile(`^[a-zA-Zа-яА-Я0-9_]+$`) опционально

// Набор запрещённых символов
var prohibitedChars = map[rune]struct{}{
	'\'': {}, '"': {}, ';': {}, '\\': {}, '`': {}, '=': {},
	'(': {}, ')': {}, '{': {}, '}': {}, '[': {}, ']': {},
}

func isOnlySpacesOrUnderscores(s string) bool {
	for _, ch := range s {
		if ch != ' ' && ch != '_' && ch != '\t' {
			return false
		}
	}
	return true
}

func containsProhibitedChars(s string) bool {
	for _, ch := range s {
		if _, exists := prohibitedChars[ch]; exists {
			return true
		}
	}
	return false
}

func RegisterUserResponseMaping(userReq domain.RegisterUserDB) RegisterUserResponse {
	return RegisterUserResponse{
		ID:    userReq.Id,
		Login: userReq.Username,
	}
}

func RegisterUserMaping(userReq UserRequest) domain.RegisterUserInternal {
	return domain.RegisterUserInternal{
		Username: userReq.Username,
		Password: userReq.Password,
	}
}
func LoginUserMaping(userReq UserRequest) domain.LoginUserInternal {
	return domain.LoginUserInternal{
		Username: userReq.Username,
		Password: userReq.Password,
	}
}
func LoginUserDbMaping(userReq domain.LoginUserInternal) domain.LoginUserDB {
	return domain.LoginUserDB{
		Username: userReq.Username,
	}
}
func LoginUserResponseMapping(token string) LoginUserResponse {
	return LoginUserResponse{JwtToken: token}
}

// func RegisterUserHashMapping(passwordHash string) domain.RegisterUserInternal {
// 	return domain.RegisterUserInternal{Password: passwordHash}
// }

func RegisterUserDbMaping(userReq domain.RegisterUserInternal) domain.RegisterUserDB {
	return domain.RegisterUserDB{
		Username:     userReq.Username,
		PasswordHash: userReq.Password,
	}
}

func UserValidate(userReq UserRequest) error {
	if userReq.Username == "" || userReq.Password == "" {
		return ErrEmptyLoginPassword
	}
	if isOnlySpacesOrUnderscores(userReq.Username) {
		return ErrUsernameSpaces
	}
	if containsProhibitedChars(userReq.Username) {
		return ErrUsernameProhibited
	}
	if !registerRegexp.MatchString(userReq.Username) {
		return ErrUsernameInvalidChars
	}
	if len(userReq.Username) < 5 || len(userReq.Username) > 20 {
		return ErrUsernameLength
	}
	if len(userReq.Password) < 9 {
		return ErrPasswordLength
	}

	return nil
}
