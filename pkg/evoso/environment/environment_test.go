package environment

import (
	"testing"
)

func TestEnvironmentTimeQUantumChannel(t *testing.T) {
	t.Log("TestEnvironmentTimeQuantumChannel")
	newEnvironment := NewEnvironment()

	for cont := 0; cont < 10; cont++ {
		testThing := NewThing(TimeQuantum, nil)
		newEnvironment.AddThing(testThing)
	}

}
