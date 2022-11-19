package ds1820

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Reader func() (string, error)

func readTemperatureSensor(reader Reader) (float64, error) {

	data, err := reader()
	if err != nil {
		return 0.0, err
	}

	lines := strings.Split(strings.ReplaceAll(data, "\r\n", "\n"), "\n")

	r, _ := regexp.Compile(".* t=(.*)")
	for _, line := range lines {
		match := r.FindStringSubmatch(line)
		if len(match) > 1 {
			temp, err := strconv.ParseFloat(match[1], 64)
			return temp / 1000, err
		}
	}
	return 0.0, errors.New("No temperature found in ds1820 data.")
}
