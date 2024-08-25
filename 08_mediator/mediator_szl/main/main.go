package main

import (
	"fmt"
	"mediator_szl"
	"mediator_szl/cd_driver"
	"mediator_szl/cpu_driver"
	"mediator_szl/player/sound_card"
	"mediator_szl/player/video_card"
	"mediator_szl/types"
)

func main() {
	mediator := mediator_szl.GetMediatorInstance()
	mediator.CDDriver = cd_driver.NewCDDriver()
	mediator.CPUDriver = cpu_driver.NewCpuDriver()
	mediator.VideoCard = video_card.NewNormalVideoCard()
	mediator.SoundCard = sound_card.NewNormalSoundCard()

	mediator.CDDriver.ReadDataV2(types.Data{
		Video: "img",
		Sound: "music",
	})

	wantCPUSound := "music1"
	if got := mediator.CPUDriver.GetSound(); got != wantCPUSound {
		fmt.Printf("want:%s,got:%s", wantCPUSound, got)
	}
}
