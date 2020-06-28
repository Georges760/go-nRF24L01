package nrf24l01

import (
	"github.com/Georges760/go-regdis"
)

func dissectSetupAw(val uint8) (ret string) {
	ret = "Setup of Addresses Widths Register\n"
	elements := []regdis.Element{}
	elements = append(elements, regdis.Element{
		BitOffset:  2,
		BitSize:    6,
		ResetValue: 0,
		Name:       "Reserved",
		Type:       "R/W",
		Desc:       "Only '000000' allowed",
	})
	elements = append(elements, regdis.Element{
		BitOffset:  0,
		BitSize:    2,
		ResetValue: 3,
		Name:       "AW",
		Type:       "R/W",
		Desc:       "RX/TX Address field width",
		Semantic: map[uint64]string{
			0: "illegal",
			1: "3 bytes",
			2: "4 bytes",
			3: "5 bytes",
		},
	})
	ret += regdis.Dissect(uint64(val), elements)
	return
}
