package tools

import "time"

func GetTimeInFormatISO8601() (timeStr string) {
	gmt, err := GetGMTLocation()

	if err != nil {
		panic(err)
	}
	return time.Now().In(gmt).Format("2006-01-02T15:04:05Z")
}

func GetGMTLocation() (*time.Location, error) {
	if LoadLocationFromTZData != nil && TZData != nil {
		return LoadLocationFromTZData("GMT", TZData)
	} else {
		return time.LoadLocation("GMT")
	}
}

var TZData []byte = nil
var LoadLocationFromTZData func(name string, data []byte) (*time.Location, error) = nil

func GetTimeInFormatRFC2616() (timeStr string) {
	gmt, err := GetGMTLocation()

	if err != nil {
		panic(err)
	}
	return time.Now().In(gmt).Format("Mon, 02 Jan 2006 15:04:05 GMT")
}

