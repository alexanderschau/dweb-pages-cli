package functions

import (
	"encoding/json"
	"io/ioutil"

	"go.alxs.xyz/dweb-pages/types"
)

func UpdateSettings(settings types.Settings) {
	jsn, err := json.Marshal(settings)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(".dweb-pages/settings.json", jsn, 0644)
	Then(settings)
}
