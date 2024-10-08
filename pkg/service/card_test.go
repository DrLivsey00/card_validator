package service_test

import (
	"github.com/DrLivsey00/card_checker/pkg/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidNumber(t *testing.T) {
	validCardNumber := "4111111111111111"
	isValid, err, code := service.IsValidNumber(validCardNumber)

	assert.True(t, isValid, "Expected card number to be valid")
	assert.NoError(t, err, "Expected no error for valid card number")
	assert.Empty(t, code, "Expected no error code for valid card number")
}

func TestIsValidNumber_InvalidLength(t *testing.T) {
	invalidCardNumber := "123"
	isValid, err, code := service.IsValidNumber(invalidCardNumber)

	assert.False(t, isValid, "Expected card number to be invalid due to length")
	assert.Error(t, err, "Expected an error for invalid card number")
	assert.Equal(t, "001", code, "Expected error code '001' for invalid length")
}

func TestIsValidNumber_InvalidCharacters(t *testing.T) {
	// Test an invalid card number (non-digit characters)
	invalidCardNumber := "4111a1111111111"
	isValid, err, code := service.IsValidNumber(invalidCardNumber)
	assert.False(t, isValid, "Expected card number to be invalid due to non-digit characters")
	assert.Error(t, err, "Expected an error for invalid card number")
	assert.Equal(t, "002", code, "Expected error code '002' for non-digit characters")
}

func Test_Luhn_Alg(t *testing.T) {
	invalidCardNumber := "1111111111111"
	isValid, err, code := service.IsValidNumber(invalidCardNumber)
	assert.False(t, isValid, "Expected card number to be invalid")
	assert.Error(t, err, "Expected an error for invalid card number")
	assert.Equal(t, "006", code, "Expected error code '006' for wrong Luhn test")
}

func TestIsValidExpDate_WrongMonth(t *testing.T) {
	month := 13
	year := 2028
	isvalid, err, code := service.IsValidExpDate(month, year)
	assert.False(t, isvalid, "Expected date to be invalid")
	assert.Error(t, err, "Expected an error for invalid date")
	assert.Equal(t, "003", code, "Expected error code '003' for wrong month")
}

func TestIsValidExpDate_WrongYear(t *testing.T) {
	month := 12
	year := 2006
	isValid, err, code := service.IsValidExpDate(month, year)
	assert.False(t, isValid, "Expected date to be invalid")
	assert.Error(t, err, "Expected an error for invalid date")
	assert.Equal(t, "004", code, "Expected error code '004' for wrong year")
}

func TestIsValidExpDate_InvalidMonth2(t *testing.T) {
	month := 9
	year := 2024
	isValid, err, code := service.IsValidExpDate(month, year)
	assert.False(t, isValid, "Expected date to be invalid")
	assert.Error(t, err, "Expected an error for invalid date")
	assert.Equal(t, "005", code, "Expected error code '005' for wrong month in the same year")
}
func TestIsValidExpDate_ValidDate(t *testing.T) {
	month := 10
	year := 2028
	isValid, err, code := service.IsValidExpDate(month, year)
	assert.True(t, isValid, "Expected date to be valid")
	assert.NoError(t, err, "Expected no error for valid date")
	assert.Equal(t, "", code, "Expected no error code for valid date")
}
