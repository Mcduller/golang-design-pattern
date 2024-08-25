package video_card

import (
	"fmt"
	"mediator_szl/domain"
)

type normalVideoCard struct {
	data string
}

func (n *normalVideoCard) Play(content string) string {
	n.data = content
	return fmt.Sprintf("normal video crd play:%s", content)
}

func NewNormalVideoCard() domain.Player {
	return &normalVideoCard{}
}
