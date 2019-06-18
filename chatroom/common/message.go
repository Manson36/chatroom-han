package message

const (
	LoginMesType 	= "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType 	= "Register"
	RegisterResMesType	= "RegisterResMes"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	UserId int  `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code int `json:"code"`//返回的状态码，500表示未注册，200表示登陆成功
	Error string `json:"error"`
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code int `json:"code"`//返回的状态码，400表示用户Id已被占用，200表示注册成功
	Error string `json:"error"`
}