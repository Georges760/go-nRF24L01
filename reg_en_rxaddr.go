package nrf24l01

import (
	"github.com/Georges760/go-regdis"
)

func dissectEnRxaddr(val uint8) (ret string) {
	ret = "Enable RX Address Register\n"
	elements := []regdis.Element{}
	elements = append(elements, regdis.Element{
		BitOffset:  7,
		BitSize:    2,
		ResetValue: 0,
		Name:       "Reserved",
		Type:       "R/W",
		Desc:       "Only '00' allowed",
	})
	elements = append(elements, regdis.Element{
		BitOffset:  5,
		BitSize:    1,
		ResetValue: 0,
		Name:       "ERX_P5",
		Type:       "R/W",
		Desc:       "Enable data pipe 5",
		Semantic:   regdis.SemanticEnable,
	})
	elements = append(elements, regdis.Element{
		BitOffset:  4,
		BitSize:    1,
		ResetValue: 0,
		Name:       "ERX_P4",
		Type:       "R/W",
		Desc:       "Enable data pipe 4",
		Semantic:   regdis.SemanticEnable,
	})
	elements = append(elements, regdis.Element{
		BitOffset:  3,
		BitSize:    1,
		ResetValue: 0,
		Name:       "ERX_P3",
		Type:       "R/W",
		Desc:       "Enable data pipe 3",
		Semantic:   regdis.SemanticEnable,
	})
	elements = append(elements, regdis.Element{
		BitOffset:  2,
		BitSize:    1,
		ResetValue: 0,
		Name:       "ERX_P2",
		Type:       "R/W",
		Desc:       "Enable data pipe 2",
		Semantic:   regdis.SemanticEnable,
	})
	elements = append(elements, regdis.Element{
		BitOffset:  1,
		BitSize:    1,
		ResetValue: 1,
		Name:       "ERX_P1",
		Type:       "R/W",
		Desc:       "Enable data pipe 1",
		Semantic:   regdis.SemanticEnable,
	})
	elements = append(elements, regdis.Element{
		BitOffset:  0,
		BitSize:    1,
		ResetValue: 1,
		Name:       "ERX_P0",
		Type:       "R/W",
		Desc:       "Enable data pipe 0",
		Semantic:   regdis.SemanticEnable,
	})
	ret += regdis.Dissect(uint64(val), elements)
	return
}
