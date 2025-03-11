package main

import (
	"testing"
	"time"
)

func TestFormatDuration(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		want     string
	}{
		{
			name:     "hours only",
			duration: 2 * time.Hour,
			want:     "2h",
		},
		{
			name:     "minutes only",
			duration: 30 * time.Minute,
			want:     "30m",
		},
		{
			name:     "hours and minutes",
			duration: 2*time.Hour + 30*time.Minute,
			want:     "2h 30m",
		},
		{
			name:     "zero duration",
			duration: 0,
			want:     "0m",
		},
		{
			name:     "single hour",
			duration: 1 * time.Hour,
			want:     "1h",
		},
		{
			name:     "single minute",
			duration: 1 * time.Minute,
			want:     "1m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatDuration(tt.duration)
			if got != tt.want {
				t.Errorf("formatDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
