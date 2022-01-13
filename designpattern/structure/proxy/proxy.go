package proxy

import (
	"log"
	"time"
)

// 静态代理
// 不改变原始类，引入代理类进行功能的扩展
type UserScene interface {
	Login(username, password string) error
}

type User struct {

} 

func (u * User) Login(username, password string) error {
	// 业务逻辑
	return nil
}

type UserProxy struct {
	User *User
}


func (u * UserProxy) Login(username, password string) error {
	// 扩展功能
	start := time.Now()

	if err := u.User.Login(username, password); err != nil {
		return err
	}

	// 扩展的功能
	log.Printf("user login cost time: %s", time.Since(start))
	return nil
}
