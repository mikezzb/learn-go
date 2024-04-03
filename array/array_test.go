package array

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {

	got := SumMultiArrs([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
