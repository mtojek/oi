package main

import (
	"testing"
	"strings"
	"bytes"
)

func TestBasic(t *testing.T) {
	dataFile := `[019FD16A-9AAC-44A2-BF0A-4A8A0C8BC1D8] Auth failed due to incorrect user password
[1700A1FB-F5D4-4D3D-839D-C02B5519767D] Auth successful
[C7E42377-2E95-46D0-A0A5-851745DAD8C7] Auth successful
[C270893B-E0E3-4B27-A624-07C27CF742D7] Auth failed due to network issue
[669E846A-E6D7-4F97-972F-5714C45EDC59] Auth failed due to timeout
[9F49A673-8470-451E-AC99-CBB8D90915BA] Auth failed due to network issue
[FF10E36F-D0B5-41AF-829B-DEB80193688F] Auth successful
[BA50A7E1-F458-4B1A-A4A9-32618ACD8AF5] Auth rejected
[579E9469-9E06-4349-8F51-65E803ADF245] Auth failed due to incorrect user password
[0619D5C7-F364-4427-ADD3-6AD95B7D5B66] Auth failed due to timeout`
	patternsFile := `019FD16A-9AAC-44A2-BF0A-4A8A0C8BC1D8
C270893B-E0E3-4B27-A624-07C27CF742D7
669E846A-E6D7-4F97-972F-5714C45EDC59
9F49A673-8470-451E-AC99-CBB8D90915BA
579E9469-9E06-4349-8F51-65E803ADF245`
	expected := `[019FD16A-9AAC-44A2-BF0A-4A8A0C8BC1D8] Auth failed due to incorrect user password
[C270893B-E0E3-4B27-A624-07C27CF742D7] Auth failed due to network issue
[669E846A-E6D7-4F97-972F-5714C45EDC59] Auth failed due to timeout
[9F49A673-8470-451E-AC99-CBB8D90915BA] Auth failed due to network issue
[579E9469-9E06-4349-8F51-65E803ADF245] Auth failed due to incorrect user password
`
	check(t, dataFile, patternsFile, expected, ALL_PATTERNS_FOUND)
}

func TestFilterAll(t *testing.T) {
	dataFile := `[019FD16A-9AAC-44A2-BF0A-4A8A0C8BC1D8] Auth failed due to incorrect user password
[1700A1FB-F5D4-4D3D-839D-C02B5519767D] Auth successful
[C7E42377-2E95-46D0-A0A5-851745DAD8C7] Auth successful
[C270893B-E0E3-4B27-A624-07C27CF742D7] Auth failed due to network issue
[669E846A-E6D7-4F97-972F-5714C45EDC59] Auth failed due to timeout
[9F49A673-8470-451E-AC99-CBB8D90915BA] Auth failed due to network issue
[FF10E36F-D0B5-41AF-829B-DEB80193688F] Auth successful
[BA50A7E1-F458-4B1A-A4A9-32618ACD8AF5] Auth rejected
[579E9469-9E06-4349-8F51-65E803ADF245] Auth failed due to incorrect user password
[0619D5C7-F364-4427-ADD3-6AD95B7D5B66] Auth failed due to timeout
`
	check(t, dataFile, dataFile, dataFile, ALL_PATTERNS_FOUND)
}

func TestNotAllPatternsFound(t *testing.T) {
	dataFile := `[019FD16A-9AAC-44A2-BF0A-4A8A0C8BC1D8] Auth failed due to incorrect user password
[0619D5C7-F364-4427-ADD3-6AD95B7D5B66] Auth failed due to timeout`
	patternsFile := `019FD16A-9AAC-44A2-BF0A-4A8A0C8BC1D8
C270893B-E0E3-4B27-A624-07C27CF742D7`
	expected := `[019FD16A-9AAC-44A2-BF0A-4A8A0C8BC1D8] Auth failed due to incorrect user password
`
	check(t, dataFile, patternsFile, expected, NOT_ALL_PATTERNS_FOUND)
}

func TestNoPatterns(t *testing.T) {
	dataFile := `[019FD16A-9AAC-44A2-BF0A-4A8A0C8BC1D8] Auth failed due to incorrect user password
[0619D5C7-F364-4427-ADD3-6AD95B7D5B66] Auth failed due to timeout`
	patternsFile := ``
	expected := ``
	check(t, dataFile, patternsFile, expected, NO_PATTERNS)
}

func check(t *testing.T, dataFile, patternsFile, expectedOutput string, expectedStatus int) {
	data := strings.NewReader(dataFile)
	patterns := strings.NewReader(patternsFile)

	filtered := new(bytes.Buffer)
	actualStatus := OrderedIntersect(data, patterns, filtered)
	actual := filtered.String()

	if expectedOutput != actual {
		t.Fatalf("\nActual output:\n%s\nExpected output:\n%s", actual, expectedOutput)
	}
	if expectedStatus != actualStatus {
		t.Fatalf("\nActual status: %d\nExpected status: %d", actualStatus, expectedStatus)
	}
}