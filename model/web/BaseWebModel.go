package web

import "Hest/model/enum"

type HestResp struct {
	code int
	msg  string
	data interface{}
	obj  interface{}
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
