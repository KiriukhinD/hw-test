package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int             `validate:"min:18|max:50"`
		Email  string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole        `validate:"in:admin,stuff"`
		Phones []string        `validate:"len:11"`
		meta   json.RawMessage //nolint:unused
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			// Place your code here.
		},
		// ...
		// Place your code here.
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()

			// Place your code here.
			_ = tt
		})
	}
}
func TestValidateField(t *testing.T) {
	type testCase struct {
		fieldName   string
		value       reflect.Value
		validator   string
		expectedErr string
	}

	testCases := []testCase{
		{
			fieldName:   "TestString",
			value:       reflect.ValueOf("hello"),
			validator:   "len:5",
			expectedErr: "",
		},
		{
			fieldName:   "TestNumber",
			value:       reflect.ValueOf(10),
			validator:   "min:15",
			expectedErr: "TestNumber should be greater than 15",
		},
		{
			fieldName:   "TestEmail",
			value:       reflect.ValueOf("example.com"),
			validator:   "regexp:^\\w+@\\w+\\.\\w+$",
			expectedErr: "TestEmail should match pattern ^\\w+@\\w+\\.\\w+$",
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		err := validateField(tc.fieldName, tc.value, tc.validator)

		if err != nil && err.Error() != tc.expectedErr {
			t.Errorf("Validation failed for %s with value %v and validator %s. Expected: %s, Got: %s",
				tc.fieldName, tc.value, tc.validator, tc.expectedErr, err.Error())
		}
	}
}
