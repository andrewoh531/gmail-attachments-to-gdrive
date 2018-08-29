package clients

import (
	"testing"
	"time"
)

func Test_getStartDate(t *testing.T) {

	t.Run("getStartDate should return Sunday from a week ago", func(t *testing.T) {
		startDate := getStartDate()

		if startDate.Weekday() != time.Sunday {
			t.Errorf("Returned date should be Sunday but was %s", startDate.Weekday())
		}

		if startDate.After(time.Now().Add(- time.Duration(24 * 7) * time.Hour)) {
			t.Errorf("Date is not from a week ago but is %s", startDate)
		}
	})
}


//func Test_retrieve(t *testing.T) {
//	t.Run("testing retrieve", func(t *testing.T) {
//		Retrieve("dummy-refresh")
//	})
//}
