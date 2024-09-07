package tiime

import "time"

const DaysInYear = 365

var totalDaysAtMonthStart = [12]int64{
	0,
	31,
	59,
	90,
	120,
	151,
	181,
	212,
	243,
	273,
	304,
	334,
}

func DayOfYearToMonthDay(dayOfYear int64) (month, day int64) {
	switch {
	case dayOfYear <= totalDaysAtMonthStart[time.January]:
		return 1, dayOfYear - totalDaysAtMonthStart[0]
	case dayOfYear <= totalDaysAtMonthStart[time.February]:
		return 2, dayOfYear - totalDaysAtMonthStart[1]
	case dayOfYear <= totalDaysAtMonthStart[time.March]:
		return 3, dayOfYear - totalDaysAtMonthStart[2]
	case dayOfYear <= totalDaysAtMonthStart[time.April]:
		return 4, dayOfYear - totalDaysAtMonthStart[3]
	case dayOfYear <= totalDaysAtMonthStart[time.May]:
		return 5, dayOfYear - totalDaysAtMonthStart[4]
	case dayOfYear <= totalDaysAtMonthStart[time.June]:
		return 6, dayOfYear - totalDaysAtMonthStart[5]
	case dayOfYear <= totalDaysAtMonthStart[time.July]:
		return 7, dayOfYear - totalDaysAtMonthStart[6]
	case dayOfYear <= totalDaysAtMonthStart[time.August]:
		return 8, dayOfYear - totalDaysAtMonthStart[7]
	case dayOfYear <= totalDaysAtMonthStart[time.September]:
		return 9, dayOfYear - totalDaysAtMonthStart[8]
	case dayOfYear <= totalDaysAtMonthStart[time.October]:
		return 10, dayOfYear - totalDaysAtMonthStart[9]
	case dayOfYear <= totalDaysAtMonthStart[time.November]:
		return 11, dayOfYear - totalDaysAtMonthStart[10]
	case dayOfYear <= DaysInYear:
		return 12, dayOfYear - totalDaysAtMonthStart[11]
	default:
		panic("invalid day of year")
	}
}

func MonthDayToDayOfYear(month time.Month, day int64) int64 {
	switch month {
	case time.January:
		return day
	case time.February:
		return day + totalDaysAtMonthStart[1]
	case time.March:
		return day + totalDaysAtMonthStart[2]
	case time.April:
		return day + totalDaysAtMonthStart[3]
	case time.May:
		return day + totalDaysAtMonthStart[4]
	case time.June:
		return day + totalDaysAtMonthStart[5]
	case time.July:
		return day + totalDaysAtMonthStart[6]
	case time.August:
		return day + totalDaysAtMonthStart[7]
	case time.September:
		return day + totalDaysAtMonthStart[8]
	case time.October:
		return day + totalDaysAtMonthStart[9]
	case time.November:
		return day + totalDaysAtMonthStart[10]
	case time.December:
		return day + totalDaysAtMonthStart[11]
	default:
		panic("invalid month")
	}
}
