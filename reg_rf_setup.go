package nrf24l01

import (
	"github.com/Georges760/go-regdis"
)

func dissectRfSetup(val uint8) (ret string) {
	ret = "RF Setup Register\n"
	elements := []regdis.Element{}
	elements = append(elements, regdis.Element{
		BitOffset:  7,
		BitSize:    1,
		ResetValue: 0,
		Name:       "CONT_WAVE",
		Type:       "R/W",
		Desc:       "Enable continuous carrier transmit when high",
		Semantic:   regdis.SemanticEnable,
	})
	elements = append(elements, regdis.Element{
		BitOffset:  6,
		BitSize:    1,
		ResetValue: 0,
		Name:       "Reserved",
		Type:       "R/W",
		Desc:       "Only '0' allowed",
	})
	elements = append(elements, regdis.Element{
		BitOffset:  5,
		BitSize:    1,
		ResetValue: 0,
		Name:       "RF_DR_LOW",
		Type:       "R/W",
		Desc:       "Set RF Data Rate to 250kbps",
		Semantic:   regdis.SemanticEnable,
	})
	elements = append(elements, regdis.Element{
		BitOffset:  4,
		BitSize:    1,
		ResetValue: 0,
		Name:       "PLL_LOCK",
		Type:       "R/W",
		Desc:       "Force PLL lock signal",
		Semantic:   regdis.SemanticEnable,
	})
	elements = append(elements, regdis.Element{
		BitOffset:  3,
		BitSize:    1,
		ResetValue: 1,
		Name:       "RF_DR_HIGH",
		Type:       "R/W",
		Desc:       "Select between the high speed data rates",
		Semantic: map[uint64]string{
			0: "1Mbps",
			1: "2Mbps",
		},
	})
	elements = append(elements, regdis.Element{
		BitOffset:  1,
		BitSize:    2,
		ResetValue: 3,
		Name:       "RF_PWR",
		Type:       "R/W",
		Desc:       "Set RF output power in TX mode",
		Semantic: map[uint64]string{
			0: "-18dBm (MIN)",
			1: "-12dBm (LOW)",
			2: "-6dBm (HIGH)",
			3: "0dBm (MAX)",
		},
	})
	elements = append(elements, regdis.Element{
		BitOffset: 0,
		BitSize:   1,
		Name:      "Obsolete",
		Desc:      "Don't care",
	})
	ret += regdis.Dissect(uint64(val), elements)
	return
}
