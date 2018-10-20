// Package controllers connects with DAL and business logic with data models
package controllers

import (
	"fmt"
	"time"
)

// rediskey generate redis key
func rediskey(index string, uid int64, from time.Time) string {
	return fmt.Sprintf("dongfeng_core_%v_%v_%v", index, uid, from.Format("20060102"))
}
