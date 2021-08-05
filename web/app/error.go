package app

const (
	SuccessCode = 0
	SuccessMsg  = "success"
)

const (
	InnerError      = 1001
	InvalidArgument = 1002
	PoolExist       = 1003
	PoolNotExist    = 1004
	LogoExist       = 1005
	LogoNotExist    = 1006
	SubExist        = 1007
	SubNotExist     = 1008
)

var ErrorMap = map[int]string{
	InnerError:      "inner error",
	InvalidArgument: "invalid argument",
	PoolExist:       "pool exist",
	PoolNotExist:    "pool is not exist",
	LogoExist:       "logo exist",
	LogoNotExist:    "logo is not exist",
	SubExist:        "sub exist",
	SubNotExist:     "sub not exist",
}

type HpError struct {
	Inner error
	Code  int
}

func (e *HpError) Error() string {
	if e.Inner != nil {
		return e.Inner.Error()
	} else if e.Code != 0 {
		return ErrorMap[e.Code]
	} else {
		return "not defined error"
	}
}
