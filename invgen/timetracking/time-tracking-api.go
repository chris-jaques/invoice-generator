package timetracking

type TimeTrackingAPI interface {
	getTotalHoursWorked(year int, month int) int
}
