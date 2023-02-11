package helper

import (
	"crypto/md5"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// 生成用户id
func GenUid(name string, email string) string {
	u := uuid.NewV4()
	str := u.String() + name + email
	id := md5.Sum([]byte(str))

	return fmt.Sprintf("%x", id)
}

// 生成文件id
func GenFid(key string) string {
	u := uuid.NewV4()
	str := u.String() + key
	id := md5.Sum([]byte(str))

	return fmt.Sprintf("%x", id)
}

// 生成用户空间文件id
func GenUserFid(user, name string) string {
	u := uuid.NewV4()
	str := u.String() + user + name
	id := md5.Sum([]byte(str))

	return fmt.Sprintf("%x", id)
}
