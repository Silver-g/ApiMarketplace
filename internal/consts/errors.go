package consts

const (
	ErrEmptyLoginPasswordMsg   = "error login and password cannot be empty"
	ErrUsernameLengthMsg       = "error username must be between 5 and 40 characters"
	ErrPasswordLengthMsg       = "error password must be at least 9 characters"
	ErrUsernameSpacesMsg       = "error username cannot consist only of spaces or underscores"
	ErrUsernameProhibitedMsg   = "error username contains prohibited characters"
	ErrUsernameInvalidCharsMsg = "error username contains invalid characters"
	ErrPasswordHashingMsg      = "failed to hash password"
	ErrUserAlreadyExistsMsg    = "user with this username already exists"
	ErrUserNotFoundMsg         = "error user not found"
	ErrIncorrectPasswordMsg    = "Incorrect password"
)
