package nrf24l01

import (
	"github.com/Georges760/go-regdis"
	"strconv"
)

func dissectSetupRetr(val uint8) (ret string) {
	ret = "Setup Automatic Retransmission Register\n"
	elements := []regdis.Element{}
	ard := regdis.Element{
		BitOffset:  4,
		BitSize:    4,
		ResetValue: 0,
		Name:       "ARD",
		Type:       "R/W",
		Desc:       "Auto Retransmit Delay",
		Semantic:   map[uint64]string{},
	}
	arc := regdis.Element{
		BitOffset:  0,
		BitSize:    4,
		ResetValue: 3,
		Name:       "ARC",
		Type:       "R/W",
		Desc:       "Auto Retransmit Count",
		Semantic: map[uint64]string{
			0: "Re-Transmit disabled",
		},
	}
	var i uint64
	for i = 0; i < 16; i++ {
		ard.Semantic[i] = "Wait " + strconv.Itoa(int(i+1)*250) + "uS"
		if i > 0 {
			arc.Semantic[i] = "Up to " + strconv.Itoa(int(i)) + " Re-Transmit on fail AA"
		}
	}
	elements = append(elements, ard)
	elements = append(elements, arc)
	ret += regdis.Dissect(uint64(val), elements)
	return
}
