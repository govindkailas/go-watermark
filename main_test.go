package main

import (
	"image"
	"reflect"
	"testing"
)

func Test_createWatermark(t *testing.T) {
	type args struct {
		name     string
		bgWidth  int
		bgHeight int
	}
	tests := []struct {
		name string
		args args
		want *image.RGBA
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createWatermark(tt.args.name, tt.args.bgWidth, tt.args.bgHeight); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createWatermark() = %v, want %v", got, tt.want)
			}
		})
	}
}
