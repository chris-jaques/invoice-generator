package clockify

import (
	// "bytes"
	"encoding/json"
	"net/http"
	// "io/ioutil"
	"fmt"
	"os"
	"invgen/config"
	"time"
	"strings"
	"strconv"
)

type ClockifyAPI struct {
	apiKey string
	workspaceId string
	userId string
}

type TimeEntry struct {
	Id string
	ProjectId string
	TimeInterval struct {
		Duration string
	}
}

func NewClockifyAPI(conf config.Config) *ClockifyAPI {
	c := conf.Clockify
	return &ClockifyAPI{c.ApiKey, c.WorkspaceId, c.UserId}
}

func (c ClockifyAPI) GetTotalHoursWorked(start time.Time, end time.Time) float64 {
	s := start.Format(time.RFC3339)
	e := end.Format(time.RFC3339)
	// jsonData := map[string]string{"start": start.Format(time.RFC3339), "end": end.Format(time.RFC3339)}
    // jsonValue, _ := json.Marshal(jsonData)
	url := fmt.Sprintf("https://api.clockify.me/api/v1/workspaces/%s/user/%s/time-entries?start=%s&end=%s", c.workspaceId, c.userId, s, e)
	// fmt.Printf("URL = %s\n",url)
	httpReq, _ := http.NewRequest("GET", url, nil)
	httpReq.Header.Set("X-Api-Key", c.apiKey)
	client := &http.Client{}

	response, err := client.Do(httpReq)

	if err != nil {
		fmt.Printf("The HTTP request failed with error:\n\t%s\n", err)
		os.Exit(1)
	} 
	defer response.Body.Close()

	var entries []TimeEntry
	json.NewDecoder(response.Body).Decode(&entries)
	// fmt.Printf("results:\n%#v\n\n",entries[0])
	var totalHours float64 = 0
	for _, entry := range entries {
		duration := strings.TrimPrefix(entry.TimeInterval.Duration, "PT")
		totalHours += parseDuration(duration)
	}

	return totalHours
}

// Duration Strings are in the form
// 2H30M
// 8H
// 15M
func parseDuration(duration string) float64 {
	// fmt.Println(duration)
	tmp := strings.Split(duration, "H")
	hours,_ := strconv.ParseFloat(tmp[0], 64)
	if len(tmp) > 1 {
		tmp = strings.Split(tmp[1], "M")
	} else {
		tmp = strings.Split(tmp[0], "M")
	}
	minutes,_ := strconv.ParseFloat(tmp[0], 64)

	return hours + minutes/60
}