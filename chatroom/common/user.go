package message

type User struct {
	UserId int `json:"userId"`    	 //注意：必须将字段序列化与Redis中存储的相同的样式，否则必定出错
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}
