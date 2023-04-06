package email

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/antlabs/mock/stringx"
)

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

var emailSuffixesTab = []string{
	"@gmail.com",
	"@yahoo.com",
	"@qq.com",
	"@163.com",
	"@126.com",
	"@sina.com",
	"@sohu.com",
	"@hotmail.com",
	"@aol.com",
	"@live.com",
	"@msn.com",
	"@ymail.com",
	"@icloud.com",
	"@mac.com",
	"@me.com",
	"@gmx.com",
	"@mail.com",
	"@inbox.com",
	"@live.cn",
	"@163.net",
	"@yeah.net",
	"@sogou.com",
	"@139.com",
	"@wo.com.cn",
	"@21cn.com",
	"@188.com",
	"@foxmail.com",
	"@outlook.com",
}

func emailSuffixes() string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(emailSuffixesTab))
	return emailSuffixesTab[index]
}

func Email() (string, error) {
	//username := randString(10)
	username, err := stringx.StringRange(3, 10)
	if err != nil {
		return "", err
	}

	suffix := emailSuffixes()
	return fmt.Sprintf("%s%s", username, suffix), nil
}

// func randString(length int) string {
// 	b := make([]byte, length)
// 	for i := range b {
// 		b[i] = charset[rand.Intn(len(charset))]
// 	}
// 	return strings.ToLower(string(b))
// }
