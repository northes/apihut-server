package logic

import (
	"fmt"
)

func ToLocation(lat, lng float64) string {
	return fmt.Sprintf("[%f,%f]", lat, lng)
}