package model

import (
	"blog-backend/config"

	"github.com/gomodule/redigo/redis"
)

// 存储图片id到redis数据库
func SaveImgCode(code, uuid string) error {

	//从连接池获取连接
	conn := RedisPool.Get()

	defer conn.Close()

	//写数据库
	_, err := conn.Do("setex", uuid, 60*5, code) //时效机制 5分钟
	return err
}

func CheckImg(uuid, captchaCode string) bool {
	conn := RedisPool.Get()
	defer conn.Close()

	rsp, _ := redis.String(conn.Do("get", uuid))

	return rsp == captchaCode
}

// 存储code到redis数据库
func SavePhoneCode(code, phoneNumber string) error {

	conn := RedisPool.Get()
	defer conn.Close()

	_, err := conn.Do("setex", phoneNumber, 60*5, code) //时效机制 5分钟

	return err
}

func CheckCode(phoneNumber, code string) bool {
	conn := RedisPool.Get()
	defer conn.Close()

	res, _ := redis.String(conn.Do("get", phoneNumber))
	return res == code
}

func SaveUserInfo(phoneNumber, password string) error {

	// 使用配置包中的DB实例
	result := config.DB.Create(&UserInfo{
		PhoneNumber: phoneNumber,
		UserName:    phoneNumber,
		PassWord:    password,
	})

	return result.Error
}
