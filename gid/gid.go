package gid

import (
	"github.com/google/uuid"
)

// 使用google的uuid库生成一个uuid
func GID() string {
	return uuid.NewString()
}
