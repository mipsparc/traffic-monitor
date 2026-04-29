package model

type ReportType string

const (
	Crash          ReportType = "crash"
	Collision      ReportType = "collision"
	StopInLiveLane ReportType = "stop_in_live_lane"
	StopInShoulder ReportType = "stop_in_shoulder"
	WrongWay       ReportType = "wrong_way"
	Pedestrian     ReportType = "pedestrian"
	Animal         ReportType = "animal"
	Smoke          ReportType = "smoke"
	Fire           ReportType = "fire"
	TooLowSpeed    ReportType = "too_low_speed"
	TooFastSpeed   ReportType = "too_fast_speed"
	Debris         ReportType = "debris"
	FacilityDamage ReportType = "facility_damage"
	Other          ReportType = "other"
)

var reportTypes = []ReportType{
	Crash, Collision, StopInLiveLane, StopInShoulder, WrongWay,
	Pedestrian, Animal, Smoke, Fire, TooLowSpeed, TooFastSpeed,
	Debris, FacilityDamage, Other,
}

// Validate returns true if the report type is valid
func (rt ReportType) Validate() bool {
	for _, t := range reportTypes {
		if rt == t {
			return true
		}
	}
	return false
}
