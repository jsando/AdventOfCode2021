package helpers

import (
	"reflect"
	"testing"
)

func Test_StringToIntSlice(t *testing.T) {
	expected := []int{3, 2, 1}
	slice := StringToIntSlice("3  2  1", " ")
	if !reflect.DeepEqual(expected, slice) {
		t.Errorf("Got %+v", slice)
	}
	slice = StringToIntSlice("3,2,1", ",")
	if !reflect.DeepEqual(expected, slice) {
		t.Errorf("Got %+v", slice)
	}
}
