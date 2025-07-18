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
		{"../testdata/mock_secrets.yaml", true},
		{"../testdata/mock_nosecrets.yaml", false},
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
		{"../testdata/mock_secrets.yaml", "../testdata/with-secrets-mock_secrets.yaml"},
		{"../testdata/mock_nosecrets.yaml", "../testdata/with-secrets-mock_nosecrets.yaml"},
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
