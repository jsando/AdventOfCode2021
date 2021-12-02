package day01

import (
	"reflect"
	"testing"
)

func Test_CountIncreases(t *testing.T) {
	if countIncreases([]int{}) != 0 {
		t.Error()
	}
	if countIncreases([]int{1}) != 0 {
		t.Error()
	}
	if countIncreases([]int{1, 2}) != 1 {
		t.Error()
	}
	if countIncreases([]int{1, 1}) != 0 {
		t.Error()
	}
	if countIncreases([]int{1, 1, 2, 2, 3, 3}) != 2 {
		t.Error()
	}
}

func Test_3Window(t *testing.T) {
	if !reflect.DeepEqual(get3WindowSum([]int{}),[]int{}) {
		t.Error()
	}
	if !reflect.DeepEqual(get3WindowSum([]int{1, 2, 3}),[]int{6}) {
		t.Error()
	}
	if !reflect.DeepEqual(get3WindowSum([]int{1, 2, 3, 4}),[]int{6, 9}) {
		t.Error()
	}
}
