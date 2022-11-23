package rpc

import (
	"Hest/model/entity"
	"encoding/json"
	"io"
	"net/http"
)

func ReadBaseDescribe() (*[]entity.BaseDescribe, error) {
	info, err := http.Get("tcloud.longsys.com/swagger-resources")
	if err != nil {
		return nil, err
	}
	var content []byte
	r := make([]byte, 1024)
	for {
		rows, err := info.Body.Read(r)
		if err == io.EOF {
			break
		}
		content = append(content, r[:rows]...)
	}
	var target []entity.BaseDescribe
	err = json.Unmarshal(content, &target)
	return &target, err
}
