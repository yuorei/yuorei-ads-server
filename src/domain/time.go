package domain

import (
	"time"
)

func NowJST() time.Time {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	nowJST := time.Now().In(jst)

	return nowJST
}
