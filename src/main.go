package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var HelmArgs []string

func main() {

	for _, arg := range os.Args[1:] {
		if strings.HasSuffix(arg, ".yaml") && replaceSecrets(arg) {
			HelmArgs = append(HelmArgs, "with-secrets-"+arg)
		} else {
			HelmArgs = append(HelmArgs, arg)
		}
	}

	fmt.Println(strings.Join(HelmArgs, " "))

}

func replaceSecrets(valuesFile string) bool {

	var fileLines []string
	var linesWithSecrets []string

	secretsmanagerPattern := regexp.MustCompile(`.+(@secretsmanager/arn:aws:secretsmanager:[a-z]{2}-[a-z]+-[0-9]:[0-9]+:[a-z]+.+)$`)

	readFile, err := os.Open(valuesFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	fileAsSecret := false
	for _, line := range fileLines {
		match := secretsmanagerPattern.FindStringSubmatch(line)
		if len(match) > 0 {
			fileAsSecret = true
			fmt.Fprintln(os.Stderr, "replacing secret : ", match[1])

			secret := strings.Split(match[1], "/")

			secretValue := "'" + getSecretsmanagerSecret(secret[1], secret[2]) + "'"

			linesWithSecrets = append(linesWithSecrets, strings.Replace(line, match[1], secretValue, 1))
		} else {
			linesWithSecrets = append(linesWithSecrets, line)
		}
	}
	if fileAsSecret {
		writeFileWithSecrets("with-secrets-"+valuesFile, linesWithSecrets)
		return true
	}
	return false
}

func writeFileWithSecrets(fileName string, lines []string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		writer.WriteString(line + "\n")
	}
	writer.Flush()
}
