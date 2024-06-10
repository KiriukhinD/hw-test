package main

import (
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	testCases := []struct {
		fromURL   string
		toPath    string
		offset    int64
		limit     int64
		expectErr bool
	}{
		{
			fromURL:   "testdata/input.txt",
			toPath:    "/tmp/out.txt",
			offset:    0,
			limit:     0,
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		err := Copy(tc.fromURL, tc.toPath, tc.offset, tc.limit)

		if tc.expectErr && err == nil {
			t.Errorf("Expected an error but got nil")
		}

		if !tc.expectErr && err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		// Clean up the destination file after each test
		_ = os.Remove(tc.toPath)
	}
}
