package domain

import (
	"log"
	"time"
)

func NowJST() time.Time {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Println("Failed to load location: ", err)
		return time.Now()
	}
	nowJST := time.Now().In(jst)

	return nowJST
}
