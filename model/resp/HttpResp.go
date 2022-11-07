package resp

type Resp struct {
	code int
	msg  string
	date interface{}
}

func Success(data interface{}) Resp {
	return Resp{
		200,
		"",
		data,
	}
}

func Error(msg string) Resp {
	return Resp{
		505,
		msg,
		nil,
	}
}
