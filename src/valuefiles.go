package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func containsSecrets(yamlFileName string) bool {

	yamlContent, err := os.ReadFile(yamlFileName)
	if err != nil {
		panic(err)
	}

	if strings.Contains(string(yamlContent), secretPrefix) {
		fmt.Println("Secrets found in file: ", yamlFileName)
		return true
	}
	fmt.Println("No secrets found in file: ", yamlFileName)
	return false
}

// read a YAML file, replaces any secrets found with their values from AWS Secrets Manager,
// writes the modified content to a new file prefixed with "with-secrets-".
// returns the name of the new file.
func replaceSecrets(yamlFileName string) string {

	var linesWithSecrets []string

	readFile, err := os.Open(yamlFileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		match := secretsmanagerPattern.FindStringSubmatch(fileScanner.Text())
		if len(match) > 0 {

			fmt.Fprintln(os.Stdout, "replacing secret : ", match[1])
			secretValue := "'" + getSecretsmanagerSecret(match[1]) + "'"

			linesWithSecrets = append(linesWithSecrets, strings.Split(fileScanner.Text(), "@")[0]+secretValue)

		} else {
			linesWithSecrets = append(linesWithSecrets, fileScanner.Text())
		}
	}

	return writeFileWithSecrets(yamlFileName, linesWithSecrets)
}

func writeFileWithSecrets(valueFileName string, lines []string) string {

	directory := filepath.Dir(valueFileName)
	fileName := filepath.Base(valueFileName)

	file, err := os.Create(directory + "/" + "with-secrets-" + fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		writer.WriteString(line + "\n")
	}
	writer.Flush()

	return directory + "/" + "with-secrets-" + fileName
}
