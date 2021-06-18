package time_operations

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	testLayout = "06-01-02 15:04:05.999"
)

type TestF64ToTimeItem struct {
	input float64
	expectedTime time.Time
	errorExpected bool
}

type TestTimeToF64Item struct {
	input time.Time
	expectedF64 float64
}

type TestF64ToDurationItem struct {
	input float64
	expectedDuration time.Duration
	errorExpected bool
}

func strToTime(str string) time.Time {
	t, _ := time.ParseInLocation(testLayout, str, time.Local)
	return t
}

func strToDuration(str string) time.Duration {
	d, _ := time.ParseDuration(str)
	return d
}

func TestF64ToTime(t *testing.T) {
	testCases := []TestF64ToTimeItem {
		{181020.193521, strToTime("18-10-20 19:35:21"), false},
		{181020.193521123, strToTime("18-10-20 19:35:21.123"), false},
		{1181020.193521123, time.Time{}, true},
		{184520.193521123, time.Time{}, true},
	}

	for _, c := range testCases {
		result, err := F64ToTime(c.input)
		if c.errorExpected {
			assert.Error(t, err)
		}
		assert.Equal(t, c.expectedTime, result)
	}
}

func TestTimeToF64(t *testing.T) {
	testCases := []TestTimeToF64Item {
		{strToTime("18-10-20 19:35:21"), 181020.193521},
		{strToTime("18-10-20 19:35:21.123"), 181020.193521123},
	}

	for _, c := range testCases {
		result := TimeToF64(c.input)
		assert.Equal(t, c.expectedF64, result)
	}
}

func TestF64ToDuration(t *testing.T) {
	testCases := []TestF64ToDurationItem {
		{99.00, strToDuration("2376h"), false},
		{0.01, strToDuration("1h"), false},
		{0.010000005, strToDuration("1h5ms"), false},
		{0.015959500, strToDuration("1h59m59s500ms"), false},
		{199.00, 0, true},
	}

	for _, c := range testCases {
		result, err := F64ToDuration(c.input)
		if c.errorExpected {
			assert.Error(t, err)
		}
		assert.Equal(t, c.expectedDuration, result)
	}
}