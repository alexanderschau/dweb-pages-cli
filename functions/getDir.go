package functions

import (
	"fmt"
	"os"
)

func checkDir() error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(path)
	return nil
}
