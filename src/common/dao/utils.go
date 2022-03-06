package dao

import (
	"fmt"
	"strings"
)

func JoinNumberConditions(ids []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ",", -1), "[]")
}
