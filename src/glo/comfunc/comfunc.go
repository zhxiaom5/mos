package comfunc

import (
	"fmt"
	"mos/src/glo/encrypt"
	"strings"
	"time"
)

// DefaultPassword const Params
const DefaultPassword string = "123456"

// TokenInvaild Const Params
const TokenInvaild string = "Token invaild"

// FormatTs 格式化时间戳为字符串
func FormatTs(ts int64) (res string) {
	res = time.Unix(ts, 0).Format("2006-01-02 15:04:05")
	return
}

// FormatShortTs 格式化时间戳为字符串
func FormatShortTs(ts int64) (res string) {
	res = time.Unix(ts, 0).Format("20060102150405")
	return
}

// FlorPageInt 分页整除判断
func FlorPageInt(pageSize int, total int) (res int) {
	res = 0
	if pageSize <= 0 {
		return
	}
	mix := total % pageSize
	if mix != 0 {
		res = total/pageSize + 1
	} else {
		res = total / pageSize
	}
	return
}

// GetDefaultPage 校验页数，小于0返回默认1
func GetDefaultPage(page int) (res int) {
	if page <= 0 {
		res = 1
		return
	}
	res = page
	return
}

// GetDefaultPageSize 校验每页行数，小于0返回默认20
func GetDefaultPageSize(pageSize int) (res int) {
	if pageSize <= 0 {
		res = 20
		return
	}
	res = pageSize
	return
}

// EncryptToken Init Token
func EncryptToken(srcStr string, timeTs int64, key string) (token string) {
	encStr := fmt.Sprintf("%s_%d", srcStr, timeTs)
	token, _ = encrypt.AesEncryptString([]byte(encStr), []byte(key))
	return
}

// DecryptToken Token
func DecryptToken(token string, key string) (srcStr string) {
	encStr, _ := encrypt.AesDecryptString(token, []byte(key))
	srcStr = strings.Split(string(encStr), "_")[0]
	return
}

// GetDayByTimeStampRange 获取时间戳之间的日期
func GetDayByTimeStampRange(startTs, endTs int64) []string {
	var (
		retDay []string
		tmpTs  int64
		i      int64
	)
	if startTs > endTs {
		tmpTs = startTs
		startTs = endTs
		endTs = tmpTs
	}
	i = 0
	for {
		tmpTs = i*86400 + startTs
		// tmpTs = endTs - i*86400
		retDay = append(retDay, time.Unix(tmpTs, 0).Format("2006-01-02"))
		if tmpTs > endTs {
			break
		}
		i++

		// if tmpTs == startTs {
		// 	retDay = append(retDay, time.Unix(tmpTs, 0).Format("2006-01-02"))
		// }

		//fmt.Println(tmpTs)
	}
	return retDay
}

// GetTodayFirstTs 获取当天凌晨时间戳
func GetTodayFirstTs() int64 {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", timeStr)
	timeNumber := t.Unix()
	return timeNumber
}

// StrArrayIndexOf 检查字符串数组中是否存在
func StrArrayIndexOf(arr []string, val string) (index int, status bool) {
	index = 0
	status = false
	for i, r := range arr {
		if r == val {
			index = i
			status = true
			break
		}
	}
	return
}
