package activities

// Daily is a type for daily resource path options.
// https://dev.fitbit.com/build/reference/web-api/activity/#activity-time-series
type Daily string

// Intra is a type for intra-day resource path options.
// https://dev.fitbit.com/build/reference/web-api/activity/#get-activity-intraday-time-series
type Intra string

// Resource path options for various activities.
const (
	DailyCalories             Daily = "activities/calories"
	DailyCaloriesBMR          Daily = "activities/caloriesBMR"
	DailySteps                Daily = "activities/steps"
	DailyDistance             Daily = "activities/distance"
	DailyFloors               Daily = "activities/floors"
	DailyElevation            Daily = "activities/elevation"
	DailyMinutesSedentary     Daily = "activities/minutesSedentary"
	DailyMinutesLightlyActive Daily = "activities/minutesLightlyActive"
	DailyMinutesFairlyActive  Daily = "activities/minutesFairlyActive"
	DailyMinutesVeryActive    Daily = "activities/minutesVeryActive"
	DailyActivityCalories     Daily = "activities/activityCalories"

	DailyTrackerCalories             Daily = "activities/tracker/calories"
	DailyTrackerSteps                Daily = "activities/tracker/steps"
	DailyTrackerDistance             Daily = "activities/tracker/distance"
	DailyTrackerFloors               Daily = "activities/tracker/floors"
	DailyTrackerElevation            Daily = "activities/tracker/elevation"
	DailyTrackerMinutesSedentary     Daily = "activities/tracker/minutesSedentary"
	DailyTrackerMinutesLightlyActive Daily = "activities/tracker/minutesLightlyActive"
	DailyTrackerMinutesFairlyActive  Daily = "activities/tracker/minutesFairlyActive"
	DailyTrackerMinutesVeryActive    Daily = "activities/tracker/minutesVeryActive"
	DailyTrackerActivityCalories     Daily = "activities/tracker/activityCalories"

	IntraCalories  Intra = "activities/calories"
	IntraSteps     Intra = "activities/steps"
	IntraDistance  Intra = "activities/distance"
	IntraFloors    Intra = "activities/floors"
	IntraElevation Intra = "activities/elevation"
)
