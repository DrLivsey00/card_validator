package service

import (
	"errors"
	"strconv"
	"time"
	"unicode"
)

type CardSrv struct {
}

func NewCardSrv() *CardSrv {
	return &CardSrv{}
}

func (c *CardSrv) IsValid(number string, expMonth, expYear int) (bool, error, string) {
	validNumber, err, code := IsValidNumber(number)
	if err != nil {
		return false, err, code
	}
	validExpDate, err, code := IsValidExpDate(expMonth, expYear)
	if err != nil {
		return false, err, code
	}

	return validNumber && validExpDate, nil, ""
}

func IsValidNumber(number string) (bool, error, string) {
	if len(number) < 13 || len(number) > 19 {
		return false, errors.New("number length should be between 13 and 19"), "001"
	}
	for _, char := range number {
		if !unicode.IsDigit(char) {
			return false, errors.New("number must contain only digits"), "002"
		}
	}
	sum := 0
	for i, char := range number {
		digit, _ := strconv.Atoi(string(char))
		if (len(number)-i)%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	if !(sum%10 == 0) {
		return false, errors.New("failing Luhn Algorithm Check "), "006"
	}
	return true, nil, ""
}

func IsValidExpDate(expMonth, expYear int) (bool, error, string) {
	currentYear, currentMonth, _ := time.Now().Date()
	if expMonth < 1 || expMonth > 12 {
		return false, errors.New("expiration month must be between 1 and 12"), "003"
	}
	if currentYear > expYear {
		return false, errors.New("invalid expiration year"), "004"
	}

	if expYear == currentYear && int(currentMonth) > expMonth {
		return false, errors.New("expiration month must be less than current month"), "005"
	}
	return true, nil, ""
}
