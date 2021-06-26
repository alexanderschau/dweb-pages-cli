package functions

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"go.alxs.xyz/dweb-pages/types"
)

func Then(settings types.Settings) {
	if _, err := os.Stat(settings.Then); os.IsNotExist(err) {
		return
	}
	cmd := exec.Command("sh", settings.Then)
	cmd.Env = append(cmd.Env, fmt.Sprintf("currentCID=%s", settings.Current))
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}
