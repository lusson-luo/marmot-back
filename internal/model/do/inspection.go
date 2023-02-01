package do

import "time"

type Inspection struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Count        int       `json:"count"`
	SuccessCount int       `json:"successCount"`
	FailedCount  int       `json:"failedCount"`
	Connection   bool      `json:"connection"`
	Availability bool      `json:"availability"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
}
