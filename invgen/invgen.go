package main

import (
    "flag"
    "fmt"
    "os"
    "time"
    "invgen/timetracking/clockify"
    "invgen/config"
)

func main() {
    var year int
    var month int
    var configFile string
    flag.IntVar(&year, "year", time.Now().Year(), "Report Year. (Required)")
    flag.IntVar(&month, "month", 0, "Report Month. (Required)")
    flag.StringVar(&configFile, "config", "/root/.invgen.conf", "Alternative config (Optional)")
    flag.Parse()

    if month == 0 {
        flag.PrintDefaults()
        os.Exit(1)
    }

    startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
    endDate := startDate.AddDate(0, 1, 0).Add(time.Nanosecond * -1)

    fmt.Printf("Searching Clockify for billable hours from %s to %s\n", startDate, endDate)
    conf := config.ParseConfig()
    c := clockify.NewClockifyAPI(*conf)
    totalHours := c.GetTotalHoursWorked(startDate, endDate)
    fmt.Printf("TOTAL HOURS CALCULATED :: %#v\n\n", totalHours)
}