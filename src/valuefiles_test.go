package main

import "testing"

func SecretBackendMock(secretName string) string {

	return "mocked secret value"
}

func TestContainsSecrets(t *testing.T) {

	tests := []struct {
		fileName string
		expected bool
	}{
		{"../testdata/secrets.yaml", true},
		{"../testdata/nosecrets.yaml", false},
	}

	for _, test := range tests {
		result := containsSecrets(test.fileName)
		if result != test.expected {
			t.Errorf("Expected %v for file %s, got %v", test.expected, test.fileName, result)
		}
	}
}

func TestOutputfileName(t *testing.T) {
	tests := []struct {
		fileName               string
		expectedOutputFilename string
	}{
		{"../testdata/secrets.yaml", "../testdata/with-secrets-secrets.yaml"},
		{"../testdata/nosecrets.yaml", "../testdata/with-secrets-nosecrets.yaml"},
	}

	for _, test := range tests {
		result := replaceSecrets(test.fileName, SecretBackendMock)
		if result != test.expectedOutputFilename {
			t.Errorf("Expected %s for file %s, got %s", test.expectedOutputFilename, test.fileName, result)
		} else {
			t.Logf("Output file name for %s is : %s", test.fileName, result)
		}
	}
}
