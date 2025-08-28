package authorization

type Authorization struct {
	NeedInputUsername    bool
	IsFirstAuthorization bool
	AllowedPass          bool
}

func NewAuthorization() *Authorization {
	return &Authorization{
		NeedInputUsername:    false,
		IsFirstAuthorization: true,
		AllowedPass:          false,
	}
}
