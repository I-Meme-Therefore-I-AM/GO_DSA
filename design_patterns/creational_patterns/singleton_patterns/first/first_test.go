package first

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()

	if counter1 == nil {
		t.Error("A new connection object must have been made")
	}

	expectedCounter := counter1
	currentCount := counter1.AddOne()

	if currentCount != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 but it is %d\n", currentCount)
	}

	counter2 := GetInstance()

	if counter2 != expectedCounter {
		t.Error("Singleton instance must be different")
	}

	currentCount = counter2.AddOne()

	if currentCount != 2 {
		t.Errorf("After calling 'AddOne' using the second counter, the current count must be 2 but was %d\n", currentCount)
	}
}
