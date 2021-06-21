package vo

import (
	"encoding/json"
	"fmt"
	"github.com/cosmopolitann/clouddb/sugar"
)

type ResponseModel struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   error       `json:"error"`
	Count   int64       `json:"count"`
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
	fmt.Printf("这是 返回数据的格式 %T\n", resmodel)
	b, e := json.Marshal(resmodel)

	if e != nil {
		sugar.Log.Error("Marshal is failed.")
	}
	sugar.Log.Info("response info is successful.")
	sugar.Log.Info("这是 返回数据  ", string(b))
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
	sugar.Log.Info("response info is successful.")
	sugar.Log.Info("这是 返回数据 2 ", string(b))
	return string(b)
}
