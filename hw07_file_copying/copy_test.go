package main

import (
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	testCases := []struct {
		fromPath  string
		toPath    string
		offset    int64
		limit     int64
		expectErr bool
	}{
		{
			fromPath:  "https://github.com/KiriukhinD/hw-test/blob/8edc22474bc58da6d48053cb9a40fe1f2eafbd55/hw07_file_copying/testdata/input.txt",
			toPath:    "https://github.com/KiriukhinD/hw-test/blob/8edc22474bc58da6d48053cb9a40fe1f2eafbd55/hw07_file_copying/testdata/out_offset0_limit0.txt",
			offset:    0,
			limit:     0,
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		err := Copy(tc.fromPath, tc.toPath, tc.offset, tc.limit)

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
