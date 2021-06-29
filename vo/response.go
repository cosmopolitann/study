package vo

import (
	"encoding/json"

	"github.com/cosmopolitann/clouddb/sugar"
)

type ResponseModel struct {
	Code    int         `json:"code"`    //验证码
	Message string      `json:"message"` //消息
	Data    interface{} `json:"data"`    //数据
	Error   error       `json:"error"`   //错误
	Count   int64       `json:"count"`   //数量
}

func BuildResp() *ResponseModel {
	return &ResponseModel{
		Code: 200}
}
func ResponseSuccess(item ...interface{}) string {
	resmodel := BuildResp()
	if len(item) >= 1 {
		resmodel.Data = item[0]
	}
	if len(item) >= 2 {
		resmodel.Count = item[1].(int64)
	}
	if len(item) >= 3 {
		resmodel.Message = item[2].(string)
	}
	b, e := json.Marshal(resmodel)
	if e != nil {
		sugar.Log.Error("Marshal is failed.")
	}
	sugar.Log.Info("Response info is successful.")
	sugar.Log.Info("Response Data:", string(b))
	return string(b)
}
func ResponseErrorMsg(code int, msg string) string {
	resmodel := BuildResp()
	resmodel.Message = msg
	resmodel.Code = code
	b, e := json.Marshal(resmodel)

	if e != nil {
		sugar.Log.Error("Marshal is failed.")
	}
	sugar.Log.Info("Response info is successful.")
	sugar.Log.Info("Response Data:", string(b))
	return string(b)
}
