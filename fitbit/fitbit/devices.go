package fitbit

import (
	"encoding/json"
)

// https://dev.fitbit.com/build/reference/web-api/devices/

// Device contains information about a fitbit device.
type Device struct {
	Battery       string        `json:"battery"`
	BatteryLevel  int           `json:"batteryLevel,omitempty"`
	DeviceVersion string        `json:"deviceVersion"`
	Features      []interface{} `json:"features"`
	ID            string        `json:"id"`
	LastSyncTime  string        `json:"lastSyncTime"`
	Mac           string        `json:"mac,omitempty"`
	Type          string        `json:"type"`
}

// DevicesJSON returns a list of user devices as a raw JSON.
// The currently loggen-in user profile if 0.
func (f *Fitbit) DevicesJSON(userID uint64) ([]byte, error) {
	return f.makeGETRequest("https://api.fitbit.com/1/user/" + convertToRequestID(userID) + "/devices.json")
}

// Devices converts the raw JSON to the Device slice type.
func (f *Fitbit) Devices(jsn []byte) ([]Device, error) {
	device := []Device{}
	if err := json.Unmarshal(jsn, &device); err != nil {
		return device, err
	}

	return device, nil
}
