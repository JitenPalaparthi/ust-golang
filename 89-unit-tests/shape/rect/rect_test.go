package rect_test

import (
	"test-demo/shape/rect"
	"testing"
)

func TestNew(t *testing.T) {
	expectedRect := new(rect.Rect)
	expectedRect.L = 1
	expectedRect.B = 2

	actualRect := rect.New(1, 2) // what is the expected result?
	// 1. actualRect should not be nil
	if actualRect == nil {
		t.Log("actualRect should not be nil")
		t.Fail()
	}
	//time.Sleep(time.Second * 5)
	// 2. actualResult and ExpectedResult must be same
	if *actualRect != *expectedRect {
		t.Log("actualRect and expectedRect are not same")
		t.Fail()
	}
}
