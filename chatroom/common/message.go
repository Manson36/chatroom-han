package message

const (
	LoginMesType 	= "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType 	= "Register"
	RegisterResMesType	= "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType				= "SmsMes"
)

//定义几个用户状态常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
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

	UsersId []int //返回用户在线列表
}

type RegisterMes struct {
	User User `json:"user"`
}

type RegisterResMes struct {
	Code int `json:"code"`//返回的状态码，400表示用户Id已被占用，200表示注册成功
	Error string `json:"error"`
}

//为了配合服务端发送用户状态变化，创建
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

type SmsMes struct {
	Content	string `json:"content"`
	User //匿名结构体，继承
}