package process

import "fmt"

//因为UserMgr的实例在服务端有且只有一个；并且在很多地方使用到，所以定义为一个全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//完成对UserMgr 的初始化工作
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

//完成对onlineUsers的添加、修改
func (this *UserMgr) AddOnlineUsers(up *UserProcess) {
	this.onlineUsers[up.UserId] = up  //map的添加方法, 修改方法
}

func (this *UserMgr) DelOnlineUsers(UserId int) {
	delete(this.onlineUsers, UserId)  //map的删除方法
}

//返回当前所有在线用户列表
func (this *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return this.onlineUsers
}

//根据Id返回对应的用户
func (this *UserMgr) GetOnlineUsersById(userId int) (up *UserProcess, err error) {
	up, ok := this.onlineUsers[userId] //从map中获取值
	if !ok {
		err = fmt.Errorf("用户%d 不存在或不在线", userId)
		return
	}
	return
}