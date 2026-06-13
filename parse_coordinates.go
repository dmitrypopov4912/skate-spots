package main

import(
	"errors"
	"strconv"
)

var(
	ParseError = errors.New("Incorrect input format")
	OutOfRangeError = errors.New("Coordinate value out of range")
)

func parseCoordinates(val string, min, max float64) (float64, error){
	coord, err := strconv.ParseFloat(val, 64)
	if err != nil{
		return 0, ParseError
	}
	if coord < min || coord > max{
		return 0, OutOfRangeError
	}
	return coord, nil
}

func ValidateCoordinates(latStr, lngStr string) (float64, float64, error){
	lat, err := parseCoordinates(latStr, -90, 90)
	if err != nil{
		return 0, 0, err
	}
	lng, err := parseCoordinates(lngStr, -180, 180)
	if err != nil{
		return 0, 0, err
	}
	return lat, lng, nil
}