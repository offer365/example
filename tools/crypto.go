package tools

import (
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

// 可以获取唯一的相应的密码值，这是目前为止最难破解的，但耗时比较久100ms
func Scrypt(pwd []byte, salt []byte, ) ([]byte, error) {
	return scrypt.Key(pwd, salt, 1<<15, 8, 1, 32)
}

// 在数据库中存储密码与校验密码

// 用密码生产一个密文
func GenerateFromPassword(pwd []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(pwd, 12)
}

// 将密文与密码比较。error==nil 则相等。
func CompareHashAndPassword(cipher, pwd []byte) error {
	return bcrypt.CompareHashAndPassword(cipher, pwd)
}
