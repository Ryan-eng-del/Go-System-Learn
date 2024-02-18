package block_net

import "testing"
func TestBlockNet(t *testing.T) {
	BlockIONet()
}

func TestNotBlockNet(t *testing.T) {
	NotBlockIONet()
}

func TestNotBlockChannelNet(t *testing.T) {
	NotBlockIOChannelNet()
}

func TestNotBlockChannelGoRoutineNet(t *testing.T) {
	NotBlockIOChannelGoRoutineNet()
}