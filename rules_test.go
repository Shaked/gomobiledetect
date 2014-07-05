package mobiledetect

import "testing"

func TestGetMobileDetectionRules(t *testing.T) {
	rules := NewRules()
	count := len(rules.phoneDevices) + len(rules.tabletDevices) + len(rules.operatingSystems) + len(rules.browsers)
	values := rules.mobileDetectionRules()
	valuesLength := len(values)
	if count != valuesLength {
		t.Errorf("Values length should be the same (count %d, values %d)", count, valuesLength)
	}
}
