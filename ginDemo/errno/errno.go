package errno

import "encoding/json"

var (
	//OK
	OK = NewError(0, "OK")
	//服务级错误
	ErrServer    = NewError(10001, "服务异常，请联系管理员")
	ErrParam     = NewError(10002, "参数错误")
	ErrSignParam = NewError(10003, "签名参数错误")

	//模块级错误 - 用户模块
	ErrUserPhone   = NewError(20101, "用户手机号不合法")
	ErrUserCaptcha = NewError(20102, "用户验证码有误")

	//...
)

var _ Error = (*err)(nil)

type Error interface {
	//i为了避免被其他包实现
	i()
	//WithData 设置成功时返回的数据
	WithData(data interface{}) Error
	//WithID 设置当前请求的唯一ID
	WithID(id string) Error
	//ToString 返回JSON格式的详情
	ToString() string
}

type err struct {
	Code int         `json:"code"`         //业务编码
	Msg  string      `json:"msg"`          //错误描述
	Data interface{} `json:"data"`         //成功时返回的数据
	ID   string      `json:"id,omitempty"` //当前请求的唯一ID，便于问题定位，忽略也可以
}

func NewError(code int, msg string) Error {
	return &err{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func (e *err) i() {}

func (e *err) WithData(data interface{}) Error {
	e.Data = data
	return e
}

func (e *err) WithID(id string) Error {
	e.ID = id
	return e
}

// ToString 返回JSON格式的错误详情
func (e *err) ToString() string {
	raw, _ := json.Marshal(e)
	return string(raw)
}
