package notifications

type Notification struct {
	Event     string      `json:"event"`
	Data      interface{} `json:"data"`
	patientID int32       `json:"patient_id"`
	Timestamp int64       `json:"timestamp"`
}

const (
	EventNewMotion      = "new_motion"
	EventNewTemperature = "new_temperature"
	EventNewFoodStatus  = "new_food_status"
	EventNewHumidity    = "new_humidity"
)
