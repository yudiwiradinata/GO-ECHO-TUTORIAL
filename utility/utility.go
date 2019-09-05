package utility

import "strings"

//CreateKeyRedis -
func CreateKeyRedis(s string) string {
	return strings.ToUpper(RedisKey + "_" + s)
}
