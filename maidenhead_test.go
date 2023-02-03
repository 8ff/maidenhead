package maidenhead

import (
	"fmt"
	"math"
	"testing"
)

func TestGetGrid(t *testing.T) {
	tests := []struct {
		lat      float64
		long     float64
		expected string
		err      error
	}{
		{45.5231, -122.6765, "CN85pm", nil}, // Portland, OR
		{40.7128, -74.0060, "FN20xr", nil},  // New York, NY
		{51.5074, -0.1278, "IO91wm", nil},   // London, UK
		{-33.8599, 151.2090, "QF56od", nil}, // Sydney, Australia
		{35.6895, 139.6917, "PM95uq", nil},  // Tokyo, Japan
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%.5f %.5f", test.lat, test.long), func(t *testing.T) {
			result, err := GetGrid(test.lat, test.long)
			if test.err != nil {
				if err == nil {
					t.Fatalf("Expected error %v, but got nil", test.err)
				}
				if err.Error() != test.err.Error() {
					t.Fatalf("Expected error %v, but got %v", test.err, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if result != test.expected {
				t.Fatalf("Expected %q, but got %q", test.expected, result)
			}
		})
	}
}

func TestConvertGridLocation(t *testing.T) {
	tests := []struct {
		location string
		lat      float64
		long     float64
		err      error
	}{
		{"CN85pm", 45.5231, -122.6765, nil}, // Portland, OR
		{"FN20xr", 40.7128, -74.0060, nil},  // New York, NY
		{"IO91wm", 51.5074, -0.1278, nil},   // London, UK
		{"QF56od", -33.8599, 151.2090, nil}, // Sydney, Australia
		{"PM95uq", 35.6895, 139.6917, nil},  // Tokyo, Japan
	}

	for _, tt := range tests {
		lat, long, err := GetCoordinates(tt.location)
		if tt.err != nil {
			if err == nil {
				t.Error("Expected error, but got nil")
			} else if err.Error() != tt.err.Error() {
				t.Errorf("Expected error message %v, but got %v", tt.err, err)
			}
			continue
		}

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
			continue
		}

		if math.Abs(tt.lat-lat) > 0.1 {
			t.Errorf("Latitude mismatch, expected %f, got %f", tt.lat, lat)
		}
		if math.Abs(tt.long-long) > 0.1 {
			t.Errorf("Longitude mismatch, expected %f, got %f", tt.long, long)
		}
	}
}
