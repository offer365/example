package tools

import (
	"fmt"
	"os"
	"path/filepath"
)



// 运行时间
func RunTime(now, start int64) string {
	if now < start {
		return ""
	}
	online := now - start
	d := online / 86400
	h := (online - d*86400) / 3600
	m := (online - d*86400 - h*3600) / 60
	s := online - d*86400 - h*3600 - m*60
	return fmt.Sprintf("%02d天%02d小时%02d分钟%02d秒.", d, h, m, s)
}

// 获取运行的根目录
func GetRootDir() (p string) {
	return filepath.Dir(os.Args[0])
}


