package ds1820

import "testing"

const CONTENT = "blablablablabla\nblablablablbla t=22123"
const EXPECTED_TEMP = 22.123

func dummyReader() (string, error) {
	return CONTENT, nil
}

func TestReadSensorFile(t *testing.T) {
	temp, err := readTemperatureSensor(dummyReader)
	if err != nil {
		t.Errorf("temperature could not be read")
	}
	if temp != EXPECTED_TEMP {
		t.Errorf("temp=%f, want %f", temp, EXPECTED_TEMP)
	}
}
