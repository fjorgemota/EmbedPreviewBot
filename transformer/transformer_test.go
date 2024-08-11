package transformer

import (
	"testing"
)

func TestTransformURL(t *testing.T) {
	tests := []struct {
		name           string
		inputURL       string
		expectedURL    string
		expectingError bool
	}{
		{
			name:           "Twitter URL without query string",
			inputURL:       "https://twitter.com/someuser/status/12345",
			expectedURL:    "https://vxtwitter.com/someuser/status/12345",
			expectingError: false,
		},
		{
			name:           "Twitter URL with query string",
			inputURL:       "https://twitter.com/someuser/status/12345?ref_src=twsrc%5Etfw",
			expectedURL:    "https://vxtwitter.com/someuser/status/12345",
			expectingError: false,
		},
		{
			name:           "Instagram URL without query string",
			inputURL:       "https://instagram.com/reels/abcd",
			expectedURL:    "https://ddinstagram.com/reels/abcd",
			expectingError: false,
		},
		{
			name:           "Instagram URL with query string",
			inputURL:       "https://instagram.com/reels/abcd?utm_source=ig_web_copy_link",
			expectedURL:    "https://ddinstagram.com/reels/abcd",
			expectingError: false,
		},
		{
			name:           "Invalid URL",
			inputURL:       "https://invalid.com/somepath",
			expectedURL:    "https://invalid.com/somepath",
			expectingError: false,
		},
		{
			name:           "Malformed URL",
			inputURL:       "test",
			expectedURL:    "",
			expectingError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TransformURL(tt.inputURL)
			if (err != nil) != tt.expectingError {
				t.Errorf("TransformURL() error = %v, expectingError %v", err, tt.expectingError)
				return
			}
			if got != tt.expectedURL {
				t.Errorf("TransformURL() = %v, want %v", got, tt.expectedURL)
			}
		})
	}
}
