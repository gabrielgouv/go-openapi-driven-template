package dateutil

import "time"

func NowAsString() string {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	date := time.Now().In(loc).String()
	return date
}
