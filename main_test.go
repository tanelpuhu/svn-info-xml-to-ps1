package main

import (
	"testing"
)

func getInput(root string, relativeURL string) []byte {
	return []byte(`<?xml version="1.0" encoding="UTF-8"?>
		<info>
			<entry path="." revision="123" kind="dir">
				<url>https://example.com` + root + relativeURL + `</url>
				<relative-url>^` + relativeURL + `</relative-url>
				<repository>
					<root>https://example.com` + root + `</root>
					<uuid>f045a106-09e3-4ed6-aece-6702c4977c5b</uuid>
				</repository>
				<wc-info>
					<wcroot-abspath>/path/to/checkout/directory</wcroot-abspath>
					<schedule>normal</schedule>
					<depth>infinity</depth>
				</wc-info>
				<commit revision="123">
					<author>tanel</author>
					<date>2000-01-01T12:13:14.123456Z</date>
				</commit>
			</entry>
		</info>
	`)
}

func assertResult(t *testing.T, result string, err error, expectation string) {
	if err != nil {
		t.Errorf("Error was returned: %s", err)
	}
	if result != expectation {
		t.Errorf("Wrong result was returned: %s, expected %s", result, expectation)
	}
}

func TestInvalidInput(t *testing.T) {
	_, err := getPS1String([]byte{}, "/home/user/svn/")
	if err == nil {
		t.Errorf("Invalid input did not cause error")
	}
}

func TestTrunkNotInCWD(t *testing.T) {
	input := getInput("/repo", "/trunk/app")
	result, err := getPS1String(input, "/home/user/svn/")
	// repo + trunk because: repo always, trunk cause its not in CWD
	assertResult(t, result, err, ":repo:trunk")
}

func TestTrunkInCWD(t *testing.T) {
	input := getInput("/repo", "/trunk/app")
	result, err := getPS1String(input, "/home/user/svn/checkout/trunk/app/web/static")
	// just repo because: repo always, trunk is already in CWD
	assertResult(t, result, err, ":repo")
}

func TestBranchInCWD(t *testing.T) {
	input := getInput("/repo", "/branches/fixes/app/web/")
	result, err := getPS1String(input, "/home/user/svn/checkout/branches/fixes/app/web")
	// just repo because: repo always, "branches" and "fixes" are already in CWD
	assertResult(t, result, err, ":repo")
}

func TestBranchNotInCWD(t *testing.T) {
	input := getInput("/repo", "/branches/fixes/app/web/")
	result, err := getPS1String(input, "/home/user/tmp/branches/app/web")
	// repo + "fixes" because: repo always, "branches" is already in CWD
	assertResult(t, result, err, ":repo:fixes")
}

func TestBranchNorBranchesNotInCWD(t *testing.T) {
	input := getInput("/repo", "/branches/fixes/app/web/")
	result, err := getPS1String(input, "/home/user/tmp/app/web")
	//  repo + "branches" + "fixes" because: repo always, "branches" and "fixes" are not in CWD
	assertResult(t, result, err, ":repo:branches:fixes")
}
