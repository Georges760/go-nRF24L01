package nrf24l01

import (
	"github.com/Georges760/go-regdis"
	"strconv"
)

func dissectRfCh(val uint8) (ret string) {
	ret = "RF Channel Register\n"
	elements := []regdis.Element{}
	elements = append(elements, regdis.Element{
		BitOffset:  7,
		BitSize:    1,
		ResetValue: 0,
		Name:       "Reserved",
		Type:       "R/W",
		Desc:       "Only '0' allowed",
	})
	rf_ch := regdis.Element{
		BitOffset:  0,
		BitSize:    7,
		ResetValue: 2,
		Name:       "RF_CH",
		Type:       "R/W",
		Desc:       "Sets the frequency channel nRF24L01 operates on",
		Semantic:   map[uint64]string{},
	}
	var i uint64
	for i = 0; i < 128; i++ {
		rf_ch.Semantic[i] = "Channel " + strconv.Itoa(int(i)) + " (" + strconv.Itoa(int(2400+i)) + " MHz)"
	}
	elements = append(elements, rf_ch)
	ret += regdis.Dissect(uint64(val), elements)
	return
}
