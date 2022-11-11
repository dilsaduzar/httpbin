package handler

import "encoding/json"

type msgErr struct {
	ErrMsg string
}

func errMsg(msg string) string {
	Msg := msgErr{msg}
	outErr, err := json.Marshal(&Msg)
	if err != nil {
		return `{"Error code": "-1"}`
	}
	return string(outErr)
}
