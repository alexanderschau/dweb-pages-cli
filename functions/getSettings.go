package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"go.alxs.xyz/dweb-pages/types"
)

func GetSettings() types.Settings {
	if _, err := os.Stat(".dweb-pages/settings.json"); os.IsNotExist(err) {
		fmt.Println("No project initialized")
		os.Exit(1)
	}
	out, err := ioutil.ReadFile(".dweb-pages/settings.json")
	if err != nil {
		panic(err)
	}
	var settings types.Settings
	json.Unmarshal(out, &settings)
	return settings
}
