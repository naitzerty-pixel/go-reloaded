package main

import "strconv"

// HexToDec convertit un nombre hexadécimal en décimal (ex : "1E" -> "30").
func HexToDec(s string) (string, error) {
	n, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(n, 10), nil
}

// BinToDec convertit un nombre binaire en décimal (ex : "10" -> "2").
func BinToDec(s string) (string, error) {
	n, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(n, 10), nil
}
