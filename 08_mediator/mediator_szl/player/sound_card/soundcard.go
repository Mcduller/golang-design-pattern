package sound_card

import (
	"fmt"
	"mediator_szl/domain"
)

type normalSoundCard struct {
	Data string
}

func (n *normalSoundCard) Play(content string) string {
	n.Data = content
	return fmt.Sprintf("normal sound card play:%s", content)
}

func NewNormalSoundCard() domain.Player {
	return &normalSoundCard{}
}
