package list

import "testing"

func TestCreateList(t *testing.T) {
	l := NewList[int]()

	if l == nil {
		t.Errorf("NewList() = %v; want a non-nil value", l)
	}

}
