package clients

import "time"

func getStartDate() time.Time {
	oneWeekDuration := time.Duration(24 * 7) * time.Hour
	oneWeekAgo := time.Now().Add(-oneWeekDuration)
	return sundayOfThatWeek(oneWeekAgo)
}

func sundayOfThatWeek(date time.Time) time.Time {
	weekday := date.Weekday()
	return date.Add(- time.Duration(int(weekday *  24)) * time.Hour)
}

