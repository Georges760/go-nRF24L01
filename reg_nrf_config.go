package nrf24l01

import (
	"github.com/Georges760/go-regdis"
)

func dissectNrfConfig(val uint8) (ret string) {
	config := []regdis.Element{}
	config = append(config, regdis.Element{
		BitOffset:  7,
		BitSize:    1,
		ResetValue: 0,
		Name:       "Reserved",
		Type:       "R/W",
		Desc:       "Only 'O' allowed",
	})
	config = append(config, regdis.Element{
		BitOffset:  6,
		BitSize:    1,
		ResetValue: 0,
		Name:       "MASK_RX_DR",
		Type:       "R/W",
		Desc:       "Mask interrupt caused by RX_DR",
		Semantic: map[uint64]string{
			1: "Interrupt not reflected on the IRQ pin",
			0: "Reflect RX_DR as active low on the IRQ pin",
		},
	})
	config = append(config, regdis.Element{
		BitOffset:  5,
		BitSize:    1,
		ResetValue: 0,
		Name:       "MASK_TX_DS",
		Type:       "R/W",
		Desc:       "Mask interrupt caused by TX_DS",
		Semantic: map[uint64]string{
			1: "Interrupt not reflected on the IRQ pin",
			0: "Reflect TX_DS as active low on the IRQ pin",
		},
	})
	config = append(config, regdis.Element{
		BitOffset:  4,
		BitSize:    1,
		ResetValue: 0,
		Name:       "MASK_MAX_RT",
		Type:       "R/W",
		Desc:       "Mask interrupt caused by MAX_RT",
		Semantic: map[uint64]string{
			1: "Interrupt not reflected on the IRQ pin",
			0: "Reflect MAX_RT as active low on the IRQ pin",
		},
	})
	config = append(config, regdis.Element{
		BitOffset:  3,
		BitSize:    1,
		ResetValue: 1,
		Name:       "EN_CRC",
		Type:       "R/W",
		Desc:       "Enable CRC, forced high if one of the bit in EN_AA is high",
		Semantic: map[uint64]string{
			0: "Disabled",
			1: "Enabled",
		},
	})
	config = append(config, regdis.Element{
		BitOffset:  2,
		BitSize:    1,
		ResetValue: 0,
		Name:       "CRCO",
		Type:       "R/W",
		Desc:       "CRC encoding scheme",
		Semantic: map[uint64]string{
			0: "1 byte",
			1: "2 bytes",
		},
	})
	config = append(config, regdis.Element{
		BitOffset:  1,
		BitSize:    1,
		ResetValue: 0,
		Name:       "PWR_UP",
		Type:       "R/W",
		Desc:       "",
		Semantic: map[uint64]string{
			0: "POWER DOWN",
			1: "POWER UP",
		},
	})
	config = append(config, regdis.Element{
		BitOffset:  0,
		BitSize:    1,
		ResetValue: 0,
		Name:       "PRIM_RX",
		Type:       "R/W",
		Desc:       "RX_TX control",
		Semantic: map[uint64]string{
			0: "PTX",
			1: "PRX",
		},
	})
	ret = regdis.Dissect(uint64(val), config)
	return
}
