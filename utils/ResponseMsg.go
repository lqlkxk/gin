package utils

type ResponseMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success() ResponseMsg {
	ret := ResponseMsg{Code: 200, Msg: "success", Data: ""}
	return ret
}

func SuccessWithData(data interface{}) ResponseMsg {
	ret := ResponseMsg{Code: 200, Msg: "success", Data: data}
	return ret
}

func ErrWithMsg(msg string) ResponseMsg {
	ret := ResponseMsg{Code: 400, Msg: msg, Data: ""}
	return ret
}
func BuildMsg(code int, msg string, data interface{}) ResponseMsg {
	ret := ResponseMsg{Code: code, Msg: msg, Data: data}
	return ret
}
