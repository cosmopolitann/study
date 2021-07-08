package vo

import "errors"

// 通用错误定义
var ErrorAffectZero error = errors.New("row affect zero")
var ErrorRowNotExists error = errors.New("row not exists")
var ErrorRowIsExists error = errors.New("row is exists")
var ErrorRepeatHandle error = errors.New("repeat handle")
