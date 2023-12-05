package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var noError = func(t assert.TestingT, err error, i ...interface{}) bool {
	return !assert.Nil(t, err)
}

var noDigitFoundError = func(t assert.TestingT, err error, i ...interface{}) bool {
	return assert.ErrorIs(t, err, ErrNoDigitFound)
}

func Test_decodeLine(t *testing.T) {
	tests := []struct {
		name    string
		line    string
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{"valid test 1", "1abc2", 12, noError},
		{"valid test 2", "pqr3stu8vwx", 38, noError},
		{"valid test 3", "a1b2c3d4e5f", 15, noError},
		{"valid test 4", "treb7uchet", 77, noError},
		{"invalid data 1", "trebuchet", 0, noDigitFoundError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeLine(tt.line)
			if !tt.wantErr(t, err, fmt.Sprintf("getDigit(%v)", tt.line)) {
				return
			}

			assert.EqualValues(t, tt.want, got)
		})
	}
}

func Test_getNumber(t *testing.T) {

	tests := []struct {
		name    string
		data    string
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{"valid data 1", "1abc2", 1, noError},
		{"valid data 2", "pqr3stu8vwx", 3, noError},
		{"valid data 3", "a1b2c3d4e5f", 1, noError},
		{"valid data 4", "treb7uchet", 7, noError},
		{"invalid data 1", "trebuchet", 0, noDigitFoundError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getNumber(tt.data)
			if !tt.wantErr(t, err, fmt.Sprintf("getNumber(%v)", tt.data)) {
				return
			}
			assert.Equalf(t, tt.want, got, "getNumber(%v)", tt.data)
		})
	}
}
