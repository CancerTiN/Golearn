package process

import (
	"fmt"
)

var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

func init() {
	userMgr = &UserMgr{onlineUsers: make(map[int]*UserProcess, 1024)}
}

func (um *UserMgr) AddOnlineUser(up *UserProcess) {
	um.onlineUsers[up.UserId] = up
}

func (um *UserMgr) DelOnlineUser(userId int) {
	delete(um.onlineUsers, userId)
}

func (um *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return um.onlineUsers
}

func (um *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := um.onlineUsers[userId]
	if !ok {
		// err = errors.New("用户不在线")
		err = fmt.Errorf("用户（%d）不在线", userId)
	}
	return
}
