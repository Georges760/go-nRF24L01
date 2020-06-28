package nrf24l01

import (
	"testing"
)

func Test_InterpretTransaction(t *testing.T) {
	// Case 1
	spi := []SpiTx{
		{[2]byte{0x2A, 0x0E}},
		{[2]byte{0xD2, 0x00}},
		{[2]byte{0xF0, 0x00}},
		{[2]byte{0xF0, 0x00}},
		{[2]byte{0xF0, 0x00}},
		{[2]byte{0xF0, 0x00}},
	}
	ret, err := InterpretTransaction(spi, false)
	if err != nil {
		t.Fatal(err)
	}
	if ret != "Write Register RX_ADDR_P0 : 0xd2f0f0f0f0" {
		t.Fatal("Return String Mismatch")
	}

	// Case 2
	spi = []SpiTx{
		{[2]byte{0xE2, 0x0E}},
	}
	ret, err = InterpretTransaction(spi, false)
	if err != nil {
		t.Fatal(err)
	}
	if ret != "Flush RX FIFO" {
		t.Fatal("Return String Mismatch")
	}

	// Case 3
	spi = []SpiTx{
		{[2]byte{0x1C, 0x0E}},
		{[2]byte{0xFF, 0x00}},
	}
	ret, err = InterpretTransaction(spi, false)
	if err != nil {
		t.Fatal(err)
	}
	if ret != "Read Register DYNPD : 0x00" {
		t.Fatal("Return String Mismatch")
	}
}
