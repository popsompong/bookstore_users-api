package date_utils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDBLayout   = "2006-01-02 15:04:05"
)

var (
	countryTz = map[string]string{
		"Hungary":  "Europe/Budapest",
		"Egypt":    "Africa/Cairo",
		"Thailand": "Asia/Bangkok",
	}
)

func GetNow(country string) time.Time {
	loc, err := time.LoadLocation(countryTz[country])
	if err != nil {
		return time.Now().UTC()
	}
	return time.Now().In(loc)
}

func GetNowString() string {
	return GetNow("Thailand").Format(apiDateLayout)
}

func GetNowDBFormat() string {
	return GetNow("Thailand").Format(apiDBLayout)
}
