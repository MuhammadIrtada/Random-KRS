package helper

import (
	"strings"
	"time"
)

func StringToHours(jam string) float64{
	hour, _ := time.ParseDuration(strings.Split(jam, ".")[0] + "h" + strings.Split(jam, ".")[1] + "m")
	return hour.Hours()
}