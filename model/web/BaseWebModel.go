package web

import "Hest/model/enum"

type HestResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Obj  interface{} `json:"obj"`
}

func SUCCESS() HestResp {
	return HestResp{enum.SUCCESS, "", nil, nil}
}

func SUCCESSR(data interface{}) HestResp {
	return HestResp{enum.SUCCESS, "", data, nil}
}

func FAILER(code int, msg string) HestResp {
	return HestResp{code, msg, nil, nil}
}

func ERROR(msg string) HestResp {
	return HestResp{enum.FAILER, msg, nil, nil}
}
