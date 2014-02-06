package gomobiledetect

import "testing"

func TestGetMobileDetectionRules(t *testing.T) {
	rules := NewRules()
	count := len(rules.phoneDevices) + len(rules.tabletDevices) + len(rules.operatingSystems) + len(rules.browsers)
	values := rules.getMobileDetectionRules()
	valuesLength := len(values)
	if count != valuesLength {
		t.Errorf("Values length should be the same (count %d, values %d)", count, valuesLength)
	}
}

// public function testRules()
// {
//     $md = new Mobile_Detect;
//     $count = array_sum(array(
//         count(Mobile_Detect::getPhoneDevices()),
//         count(Mobile_Detect::getTabletDevices()),
//         count(Mobile_Detect::getOperatingSystems()),
//         count(Mobile_Detect::getBrowsers())
//     ));
//     $rules = $md->getRules();
//     $this->assertEquals($count, count($rules));
// }
