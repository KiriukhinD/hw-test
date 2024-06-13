package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

func (v ValidationErrors) Error() string {
	errors := []string{}
	for _, err := range v {
		errors = append(errors, fmt.Sprintf("%s: %s", err.Field, err))
	}
	return strings.Join(errors, "; ")
}

type ValidationErrors []ValidationError

func Validate(v interface{}) error {
	valErrors := make(ValidationErrors, 0)

	value := reflect.ValueOf(v)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		return errors.New("input must be a struct")
	}

	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		fieldValue := value.Field(i)

		validateTag := field.Tag.Get("validate")
		validators := strings.Split(validateTag, "|")

		for _, validator := range validators {
			err := validateField(field.Name, fieldValue, validator)
			if err != nil {
				valErrors = append(valErrors, ValidationError{Field: field.Name, Err: err})
			}
		}
	}

	if len(valErrors) > 0 {
		return valErrors
	}

	return nil
}

func validateField(fieldName string, value reflect.Value, validator string) error {
	rules := strings.Split(validator, ":")

	switch rules[0] {
	case "len":
		expectedLen, _ := strconv.Atoi(rules[1])
		if value.Len() != expectedLen {
			return fmt.Errorf("%s length must be %s", fieldName, rules[1])
		}
	case "regexp":
		pattern := rules[1]
		match, _ := regexp.MatchString(pattern, value.String())
		if !match {
			return fmt.Errorf("%s should match pattern %s", fieldName, pattern)
		}
	case "min":
		minValue, _ := strconv.Atoi(rules[1])
		if value.Int() < int64(minValue) {
			return fmt.Errorf("%s should be greater than %s", fieldName, rules[1])
		}
	case "max":
		maxValue, _ := strconv.Atoi(rules[1])
		if value.Int() > int64(maxValue) {
			return fmt.Errorf("%s should be less than %s", fieldName, rules[1])
		}
	case "in":
		validValues := strings.Split(rules[1], ",")
		valid := false
		for _, validValue := range validValues {
			if value.String() == validValue {
				valid = true
				break
			}
		}
		if !valid {
			return fmt.Errorf("%s should be one of %s", fieldName, rules[1])
		}
	default:
		return fmt.Errorf("unknown validator: %s", rules[0])
	}

	return nil
}
