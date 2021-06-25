package pages

import (
	"fmt"
	"regexp"
	"strings"
)

var tagCheck = regexp.MustCompile("^[[:alnum:]/]+$")

//ResolveTag - outputs a string array of branch and name
func ResolveTag(tag string) ([]string, error) {
	rawTag := strings.Split(tag, "/")

	if len(rawTag) > 2 || len(rawTag) == 0 || (len(rawTag) > 1 && rawTag[1] == "") || !tagCheck.Match([]byte(tag)) {
		return []string{}, fmt.Errorf("wrong tag format: '%s'", tag)
	}

	if rawTag[0] == "default" && len(rawTag) > 1 {
		return []string{}, fmt.Errorf("Branch 'default' is reserved for the system. Please choose another name.")
	}

	if len(rawTag) == 1 && rawTag[0] != "default" {
		rawTag = append(rawTag, "latest")
	}

	return rawTag, nil
}
