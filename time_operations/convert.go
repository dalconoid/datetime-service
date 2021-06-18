package time_operations

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	layoutF64ToTime = "060102150405.000"
	layoutTimeToF64 = "06,1,2,15,4,5.000"
	formatLayout = "2006-01-02 15:04:05.999999999 -0700 MST"
)

// F64ToTime - converts time from F64 format to time.Time format
func F64ToTime(f64time float64) (time.Time, error){
	f64time *= 1000000
	f64dt := int(f64time)
	dtStr := fmt.Sprintf("%012d", f64dt)
	f64ms := f64time - float64(f64dt)
	msStr := fmt.Sprintf("%.3f", f64ms)
	str := dtStr + msStr[1:]

	t, err := time.ParseInLocation(layoutF64ToTime, str, time.Local)
	if err != nil {
		return  t, fmt.Errorf("invalid time format: %s", err.Error())
	}

	return t, nil
}

// TimeToF64 - converts time from time.Time format to F64 format
func TimeToF64(t time.Time) float64 {
	var f64time float64 = 0
	str := t.Format(layoutTimeToF64)
	parts := strings.Split(str, ",")
	year, _ := strconv.ParseFloat(parts[0], 64)
	month, _ := strconv.ParseFloat(parts[1], 64)
	day, _ := strconv.ParseFloat(parts[2], 64)
	hour, _ := strconv.ParseFloat(parts[3], 64)
	min, _ := strconv.ParseFloat(parts[4], 64)
	sec, _ := strconv.ParseFloat(parts[5], 64)
	intPart := year * 10000 + month * 100 + day
	decimalPart := math.Round(hour * 10000000 + min * 100000 + sec * 1000)/1000000000
	f64time +=  intPart + decimalPart
	return f64time
}

// F64ToDuration - converts time from F64 format to time.Duration format
func F64ToDuration(f64time float64) (time.Duration, error) {
	var mult time.Duration
	if f64time == 0 {
		return 0, nil
	} else if f64time > 0 {
		mult = 1
	} else {
		mult = -1
	}

	f64time = math.Abs(f64time)
	days := int(f64time)
	if days > 99 {
		return 0, fmt.Errorf("invalid delta format: days are limited to 99")
	}

	str := fmt.Sprintf("%.9f", f64time) //99.103050255
	hMS := strings.Split(str, ".")[1] //103050255

	hrs, _ := strconv.Atoi(hMS[:2])
	if hrs > 24 {
		return 0, fmt.Errorf("invalid format: hours")
	}
	hrs += days * 24
	mins, _ := strconv.Atoi(hMS[2:4])
	if mins > 59 {
		return 0, fmt.Errorf("invalid format: minutes")
	}
	secs, _ := strconv.Atoi(hMS[4:6])
	if secs > 59 {
		return 0, fmt.Errorf("invalid format: seconds")
	}
	millis, _ := strconv.Atoi(hMS[6:])

	durationStr := fmt.Sprintf("%dh%dm%ds%dms", hrs, mins, secs, millis)

	d, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, err
	}

	d *= mult
	return d, nil
}

// TimeToString - converts time from time.Time format to string
func TimeToString(t time.Time) string {
	return t.Format(formatLayout)
}