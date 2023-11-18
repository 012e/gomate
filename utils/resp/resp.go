package resp

import "fmt"


type BaseFail struct {
	Message string `json:"message,omitempty"`
	Success bool   `json:"success,omitempty"`
}

type BaseOk struct {
	Message string `json:"message,omitempty"`
	Success bool   `json:"success,omitempty"`
}

func Fail(msg string) BaseFail {
	return BaseFail{Message: msg, Success: false}
}
func Failf(msg string, args ...any) BaseFail {
	return Fail(fmt.Sprintf(msg, args...))
}

func FailUnknow() BaseFail {
	return Fail("something went wrong")
}

func Ok(msg string) BaseFail {
	return BaseFail{Message: msg, Success: true}
}
func Okf(msg string, args ...any) BaseFail {
	return Ok(fmt.Sprintf(msg, args...))
}
