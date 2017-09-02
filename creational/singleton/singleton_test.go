package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()

	if counter1 == nil {
		t.Fatal("expected pointer to the counter on calling GetInstance() not to be nill")
	}

	expectedCounter := counter1

	currentCount := counter1.AddOne()

	if currentCount != 1 {
		t.Fatal("expected current count to be 1 on calling AddOne(), but current count is: %d", currentCount)
	}

	counter2 := GetInstance()

	if counter2 != expectedCounter {
		t.Fatal("expected the next call to GetInstance() return the exact same counter. but got different instance")
	}

	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Fatal("expected second call to AddOne() return the value of 2 but got: %d", currentCount)
	}
}
