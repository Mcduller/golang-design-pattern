package domain

import "mediator_szl/types"

type CPUDriver interface {
	GetSound() string
	GetVideo() string

	Process(data types.Data)
}
