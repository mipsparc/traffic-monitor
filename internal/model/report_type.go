package model

import "log/slog"

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
	Unknown        ReportType = "unknown"
)

func (rt ReportType) String() string {
	switch rt {
	case Crash:
		return "単独事故"
	case Collision:
		return "衝突事故"
	case StopInLiveLane:
		return "本線停車"
	case StopInShoulder:
		return "路肩停車"
	case WrongWay:
		return "逆走"
	case Pedestrian:
		return "歩行者侵入"
	case Animal:
		return "動物侵入"
	case Smoke:
		return "煙"
	case Fire:
		return "火災"
	case TooLowSpeed:
		return "異常に低速"
	case TooFastSpeed:
		return "極端な速度超過"
	case Debris:
		return "障害物"
	case FacilityDamage:
		return "施設破損"
	case Other:
		return "その他"
	case Unknown:
		return "不明"
	default:
		slog.Error("登録されていないレポートタイプを受領しました")
		return "不明なレポートタイプ"
	}
}
