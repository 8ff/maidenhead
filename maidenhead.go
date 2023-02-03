package maidenhead

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func GetGrid(latitude, longitude float64) (string, error) {
	// Shift latitude and longitude
	shiftedLatitude := latitude + 90
	shiftedLongitude := longitude + 180

	// Calculate field part
	fieldLatitude := int(math.Floor(shiftedLatitude / 10))
	fieldLongitude := int(math.Floor(shiftedLongitude / 20))
	fieldLatitudeChar, err := numToLetter(fieldLatitude, true)
	if err != nil {
		return "", fmt.Errorf("error calculating field latitude: %v", err)
	}
	fieldLongitudeChar, err := numToLetter(fieldLongitude, true)
	if err != nil {
		return "", fmt.Errorf("error calculating field longitude: %v", err)
	}
	field := fieldLongitudeChar + fieldLatitudeChar

	// Calculate square part
	squareLatitude := int(math.Floor((shiftedLatitude/10 - float64(fieldLatitude)) * 10))
	squareLongitude := int(math.Floor((shiftedLongitude/20 - float64(fieldLongitude)) * 10))
	square := strconv.Itoa(squareLongitude) + strconv.Itoa(squareLatitude)

	// Calculate subsquare part
	subsquareLatitude := int(math.Floor(((shiftedLatitude/10-float64(fieldLatitude))*10 - float64(squareLatitude)) * 24))
	subsquareLongitude := int(math.Floor(((shiftedLongitude/20-float64(fieldLongitude))*10 - float64(squareLongitude)) * 24))
	subsquareLatitudeChar, err := numToLetter(subsquareLatitude, false)
	if err != nil {
		return "", fmt.Errorf("error calculating subsquare latitude: %v", err)
	}
	subsquareLongitudeChar, err := numToLetter(subsquareLongitude, false)
	if err != nil {
		return "", fmt.Errorf("error calculating subsquare longitude: %v", err)
	}
	subsquare := subsquareLongitudeChar + subsquareLatitudeChar

	// Concatenate field, square, and subsquare parts
	locator := field + square + subsquare
	return locator, nil
}

func GetCoordinates(location string) (float64, float64, error) {
	if len(location) != 4 && len(location) != 6 {
		return 0, 0, fmt.Errorf("grid location must be either 4 or 6 digits")
	}

	location = strings.ToLower(location)

	l := make([]int, 6)
	var err error
	l[0], err = letterToNum(string(location[0]))
	if err != nil {
		return 0, 0, fmt.Errorf("longitude field value: %w", err)
	}
	l[1], err = letterToNum(string(location[1]))
	if err != nil {
		return 0, 0, fmt.Errorf("latitude field value: %w", err)
	}

	l[2], err = strconv.Atoi(string(location[2]))
	if err != nil {
		return 0, 0, fmt.Errorf("longitude sqare value: %w", err)
	}
	l[3], err = strconv.Atoi(string(location[3]))
	if err != nil {
		return 0, 0, fmt.Errorf("latitude sqare value: %w", err)
	}

	if len(location) == 6 {
		l[4], err = letterToNum(string(location[4]))
		if err != nil {
			return 0, 0, fmt.Errorf("longitude subsquare value: %w", err)
		}
		l[5], err = letterToNum(string(location[5]))
		if err != nil {
			return 0, 0, fmt.Errorf("latitude subsquare value: %w", err)
		}
	}

	long := (float64(l[0]) * 20) + (float64(l[2]) * 2) + (float64(l[4]) / 12) - 180
	lat := (float64(l[1]) * 10) + float64(l[3]) + (float64(l[5]) / 24) - 90

	return lat, long, nil
}

func letterToNum(input string) (output int, err error) {
	if len(input) != 1 {
		return 0, errors.New("invalid input: input must be a single character")
	}
	if input < "a" || input > "x" {
		return 0, errors.New("invalid input: input must be a letter between a and x")
	}
	output = int(input[0] - 'a')
	return
}

func numToLetter(input int, capital bool) (output string, err error) {
	if input < 0 || input > 23 {
		return "", errors.New("invalid input: input must be a number between 0 and 23")
	}
	if capital {
		output = string(rune(input + 'A'))
	} else {
		output = string(rune(input + 'a'))
	}
	return
}
