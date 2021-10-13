package main

import (
	"encoding/json"
	"fmt"

	"main/fitbit"
	"main/fitbit/scopes"
)

func main() {
	fb, err := fitbit.New("23BKXX", "03e904359d3e6cb2c3cb0666074a453c", "http://localhost:41259", []string{
		scopes.Activity,
		scopes.Settings,
		scopes.Location,
		scopes.Social,
		scopes.Heartrate,
		scopes.Profile,
		scopes.Sleep,
		scopes.Nutrition,
		scopes.Weight,
	})

	if err != nil {
		panic(fmt.Sprintf("Cannot create fitbit session, exiting: %s", err))
	}

	fmt.Println("{")
	/*d, b, err := fb.Profile(0)
	if err != nil {
		panic(fmt.Sprintf("Cannot get profile, exiting: %s", err))
	}*/

	d, b, err := fb.HeartIntraday("2021-10-12", "1sec")
	if err != nil {
		fmt.Println("\"raw1sec\":", string(b), "}")
		panic(fmt.Sprintf("Cannot get intraday 1sec, exiting: %s", err))
	}

	fmt.Println("\"raw1sec\":", string(b), ",")

	jb, _ := json.Marshal(d)
	fmt.Println("\"marsh1sec\":", string(jb), ",")

	d, b, err = fb.HeartIntraday("2021-10-12", "1min")
	if err != nil {
		panic(fmt.Sprintf("Cannot get intraday 1min, exiting: %s", err))
	}

	fmt.Println("\"raw1min\":", string(b), ",")

	jb, _ = json.Marshal(d)
	fmt.Println("\"marsh1min\":", string(jb), "")

	fmt.Println("}")
}
