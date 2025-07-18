package main

import "testing"

func TestArnParsing(t *testing.T) {
	tests := []struct {
		arn             string
		expectedParsing secret
	}{
		{"arn:aws:secretsmanager:us-east-1:012345678910:secret:test01-Ftxaat/token1",
			secret{
				Region:              "us-east-1",
				AccountId:           "012345678910",
				SecretName:          "test01-Ftxaat",
				SecretKey:           "token1",
				SecretArnWithoutKey: "arn:aws:secretsmanager:us-east-1:012345678910:secret:test01-Ftxaat"},
		},
		{"arn:aws:secretsmanager:eu-west-3:012345678910:secret:test01/Ftxaat/token1",
			secret{
				Region:              "eu-west-3",
				AccountId:           "012345678910",
				SecretName:          "test01/Ftxaat",
				SecretKey:           "token1",
				SecretArnWithoutKey: "arn:aws:secretsmanager:eu-west-3:012345678910:secret:test01/Ftxaat"},
		},
	}

	for _, test := range tests {
		result := secretsmanagerArnParser(test.arn)
		if result != test.expectedParsing {
			t.Errorf("KO! - Arn parsing\n for %s,\nExpected \n%s\ngot\n%s", test.arn, test.expectedParsing, result)
		} else {
			t.Logf("OK - Arn parsing %s is : %s", test.arn, result)
		}
	}
}
