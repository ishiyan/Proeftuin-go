package fitbit

import (
	"encoding/json"
)

// BadgeList contains a list of badges.
type BadgeList struct {
	Badges []Badge `json:"badges"`
}

// Badge contains information about a badge.
type Badge struct {
	BadgeGradientEndColor   string        `json:"badgeGradientEndColor"`
	BadgeGradientStartColor string        `json:"badgeGradientStartColor"`
	BadgeType               string        `json:"badgeType"`
	Category                string        `json:"category"`
	Cheers                  []interface{} `json:"cheers"` // FIXME: unknown data
	DateTime                string        `json:"dateTime"`
	Description             string        `json:"description"`
	EarnedMessage           string        `json:"earnedMessage,omitempty"`
	EncodedID               string        `json:"encodedId"`
	Image100Px              string        `json:"image100px"`
	Image125Px              string        `json:"image125px"`
	Image300Px              string        `json:"image300px"`
	Image50Px               string        `json:"image50px"`
	Image75Px               string        `json:"image75px"`
	MarketingDescription    string        `json:"marketingDescription"`
	MobileDescription       string        `json:"mobileDescription"`
	Name                    string        `json:"name"`
	ShareImage640Px         string        `json:"shareImage640px"`
	ShareText               string        `json:"shareText"`
	ShortDescription        string        `json:"shortDescription"`
	ShortName               string        `json:"shortName"`
	TimesAchieved           int           `json:"timesAchieved"`
	Value                   int           `json:"value,omitempty"`
	Unit                    string        `json:"unit,omitempty"`
}

// BadgesJSON returns a list of user badges as a raw JSON.
// The currently loggen-in user profile if 0.
func (f *Fitbit) BadgesJSON(userID uint64) ([]byte, error) {
	return f.makeGETRequest("https://api.fitbit.com/1/user/" + convertToRequestID(userID) + "/badges.json")
}

// Badges converts the raw JSON to the BadgeList type.
func (f *Fitbit) Badges(jsn []byte) (BadgeList, error) {
	badgeList := BadgeList{}
	if err := json.Unmarshal(jsn, &badgeList); err != nil {
		return BadgeList{}, err
	}

	return badgeList, nil
}
