package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var HelmArgs []string
var secretPrefix = "@secretsmanager@"
var secretsmanagerPattern = regexp.MustCompile(`.+` + secretPrefix + `(arn:aws:secretsmanager:[a-z]{2}-[a-z]+-[0-9]:[0-9]+:[a-z]+.+)$`)

func main() {

	for _, arg := range os.Args[1:] {
		if strings.HasSuffix(arg, ".yaml") {
			if containsSecrets(arg) {

				// helm command will be run with the new file containing secrets
				HelmArgs = append(HelmArgs, replaceSecrets(arg, getSecretsmanagerSecret))
			} else {
				HelmArgs = append(HelmArgs, arg)
			}
		}

		fmt.Println(strings.Join(HelmArgs, " "))

	}
}
