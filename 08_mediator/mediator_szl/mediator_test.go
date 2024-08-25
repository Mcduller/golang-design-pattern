package mediator_szl

import (
	"mediator_szl/cd_driver"
	"mediator_szl/cpu_driver"
	"mediator_szl/player/sound_card"
	"mediator_szl/player/video_card"
	"mediator_szl/types"
	"testing"
)

func TestMediator(t *testing.T) {
	mediator := GetMediatorInstance()
	mediator.CDDriver = cd_driver.NewCDDriver()
	mediator.CPUDriver = cpu_driver.NewCpuDriver()
	mediator.VideoCard = video_card.NewNormalVideoCard()
	mediator.SoundCard = sound_card.NewNormalSoundCard()

	mediator.CDDriver.ReadDataV2(types.Data{
		Video: "img",
		Sound: "music",
	})

	wantCPUSound := "music"
	if got := mediator.CPUDriver.GetSound(); got != wantCPUSound {
		t.Errorf("want:%s,got:%s", wantCPUSound, got)
	}

}
