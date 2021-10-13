package fitbit

import (
	"encoding/json"
)

// Profile contains profile information of an user.
type Profile struct {
	User struct {
		Age                    int    `json:"age"`
		Ambassador             bool   `json:"ambassador"`
		AutoStrideEnabled      bool   `json:"autoStrideEnabled"`
		Avatar                 string `json:"avatar"`
		Avatar150              string `json:"avatar150"`
		Avatar640              string `json:"avatar640"`
		AverageDailySteps      int    `json:"averageDailySteps"`
		ChallengesBeta         bool   `json:"challengesBeta"`
		ClockTimeDisplayFormat string `json:"clockTimeDisplayFormat"`
		Corporate              bool   `json:"corporate"`
		CorporateAdmin         bool   `json:"corporateAdmin"`
		Country                string `json:"country"`
		DateOfBirth            string `json:"dateOfBirth"`
		DisplayName            string `json:"displayName"`
		DisplayNameSetting     string `json:"displayNameSetting"`
		DistanceUnit           string `json:"distanceUnit"`
		EncodedID              string `json:"encodedId"`
		FamilyGuidanceEnabled  bool   `json:"familyGuidanceEnabled"`
		Features               struct {
			ExerciseGoal             bool `json:"exerciseGoal"`
			PhoneNumberFriendFinding struct {
				Algorithm string   `json:"algorithm"`
				Countries []string `json:"countries"`
				Salt      string   `json:"salt"`
			} `json:"phoneNumberFriendFinding"`
		} `json:"features"`
		FoodsLocale              string  `json:"foodsLocale"`
		FullName                 string  `json:"fullName"`
		Gender                   string  `json:"gender"`
		GlucoseUnit              string  `json:"glucoseUnit"`
		Height                   float64 `json:"height"`
		HeightUnit               string  `json:"heightUnit"`
		IsBugReportEnabled       bool    `json:"isBugReportEnabled"`
		IsChild                  bool    `json:"isChild"`
		IsCoach                  bool    `json:"isCoach"`
		LanguageLocale           string  `json:"languageLocale"`
		LegalTermsAcceptRequired bool    `json:"legalTermsAcceptRequired"`
		Locale                   string  `json:"locale"`
		MemberSince              string  `json:"memberSince"`
		MfaEnabled               bool    `json:"mfaEnabled"`
		OffsetFromUTCMillis      int     `json:"offsetFromUTCMillis"`
		SdkDeveloper             bool    `json:"sdkDeveloper"`
		SleepTracking            string  `json:"sleepTracking"`
		StartDayOfWeek           string  `json:"startDayOfWeek"`
		StrideLengthRunning      float64 `json:"strideLengthRunning"`
		StrideLengthRunningType  string  `json:"strideLengthRunningType"`
		StrideLengthWalking      float64 `json:"strideLengthWalking"`
		StrideLengthWalkingType  string  `json:"strideLengthWalkingType"`
		SwimUnit                 string  `json:"swimUnit"`
		Timezone                 string  `json:"timezone"`
		TopBadges                []Badge `json:"topBadges"`
		WaterUnit                string  `json:"waterUnit"`
		WaterUnitName            string  `json:"waterUnitName"`
		Weight                   float64 `json:"weight"`
		WeightUnit               string  `json:"weightUnit"`
	} `json:"user"`
}

// Profile returns the current user profile if 0 is used or the profile of a friend with matching ID.
func (f *FitBit) Profile(userID uint64) (Profile, []byte, error) {
	contents, err := f.makeGETRequest("https://api.fitbit.com/1/user/" + convertToRequestID(userID) + "/profile.json")
	if err != nil {
		return Profile{}, []byte{}, err
	}

	profile := Profile{}
	if err := json.Unmarshal(contents, &profile); err != nil {
		return Profile{}, contents, err
	}

	return profile, contents, nil
}

// SetProfile updates the profile information of a user. ID 0 is a current user.
func (f *FitBit) SetProfile(userID uint64, params map[string]string) (Profile, error) {
	contents, err := f.makePOSTRequest("https://api.fitbit.com/1/user/"+convertToRequestID(userID)+"/profile.json", params)
	if err != nil {
		return Profile{}, err
	}

	profile := Profile{}
	if err := json.Unmarshal(contents, &profile); err != nil {
		return Profile{}, err
	}

	return profile, nil
}
