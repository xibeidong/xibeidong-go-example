package State

import (
	"testing"
)

func TestState(t *testing.T) {
	tv := &BigTV{}
	tv.setState(&powerOffState{})
	tv.switchChannel()
	tv.close()
	tv.open()
	tv.switchChannel()
	tv.close()
}
