package main

import (
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrDataNotFound = errors.New("record not found")
)

type User struct {
	Id uint64 `json:"Id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

type Dao interface {
	Get(id uint64) interface{}
	List() interface{}
	Create()
	Update()
	Delete(id uint64)
}

type UserDao struct {}

// Dao层获取到底层错误，使用errors的Wrap进行包装
func(user *UserDao) Get(id uint64) (*User, error) {
user := User{}
err := db.Where("id = ?",id).Find(user).Error

if errors.Is(err,sql.ErrNoRows){
retrun errors.Wrap(err,fmt.Sprintf("find user null,user id: %v",id))
}
return &user,nil
}

// 业务层获取到错误直接往上层抛
type UserService struct {}

func (s *Service) FindUserByID(userID int) (*model.User, error) {
	return dao.FindUserByID(userID)
}
