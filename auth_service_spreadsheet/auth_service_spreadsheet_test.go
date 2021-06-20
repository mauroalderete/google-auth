package auth_service_spreadsheet

import (
	"io/fs"
	"io/ioutil"
	"os"
	"testing"
)

var testFilePrepare = []struct {
	path    string
	content string
}{
	{"./credential.json.test", "{\"installed\":{\"client_id\":\"abc\",\"project_id\":\"abc\",\"auth_uri\":\"abc\",\"token_uri\":\"abc\",\"auth_provider_x509_cert_url\":\"abc\",\"client_secret\":\"abc\",\"redirect_uris\":[\"urn:ietf:wg:oauth:2.0:oob\",\"http://localhost\"]}}"},
	{"./token.json.test", "{\"access_token\":\"abc\",\"token_type\":\"Bearer\",\"refresh_token\":\"abc\",\"expiry\":\"2021-06-19T21:15:54.8451088-03:00\"}"},
}

var testCorpus = []struct {
	name           string
	credentialPath string
	tokenPath      string
	readonly       bool
	expected       bool
}{
	{"1", "./credential.json.test", "./token.json.test", true, true},
	{"2", "./credential.json.test", "./token.json.test", false, true},
	{"3", "./otro", "./token.json.test", false, false},
	{"4", "./credential.json.test", "./otro", false, false},
}

func TestAuthServiceSpreadsheetInitialize(t *testing.T) {
	for _, pr := range testFilePrepare {
		ioutil.WriteFile(pr.path, []byte(pr.content), fs.ModeAppend.Perm())
	}

	for _, tc := range testCorpus {
		t.Run(tc.name, func(t *testing.T) {
			var auth AuthServiceSpreadsheet
			got := auth.Initialize(tc.credentialPath, tc.tokenPath, tc.readonly)
			if (got == nil) != tc.expected {
				t.Errorf("Load Expected: %v, got: %v", tc.expected, got)
			} else {
				if auth.Readonly != tc.readonly {
					t.Errorf("Readonly Expected: %v, got: %v", tc.readonly, auth.Readonly)
				}
			}
		})
	}

	for _, pr := range testFilePrepare {
		os.Remove(pr.path)
	}
}
