package nrf24l01

import (
	"errors"
	"strconv"
)

type Register uint8

/* Memory Map */
const (
	REG_NRF_CONFIG  Register = 0x00
	REG_EN_AA       Register = 0x01
	REG_EN_RXADDR   Register = 0x02
	REG_SETUP_AW    Register = 0x03
	REG_SETUP_RETR  Register = 0x04
	REG_RF_CH       Register = 0x05
	REG_RF_SETUP    Register = 0x06
	REG_NRF_STATUS  Register = 0x07
	REG_OBSERVE_TX  Register = 0x08
	REG_CD          Register = 0x09
	REG_RX_ADDR_P0  Register = 0x0A
	REG_RX_ADDR_P1  Register = 0x0B
	REG_RX_ADDR_P2  Register = 0x0C
	REG_RX_ADDR_P3  Register = 0x0D
	REG_RX_ADDR_P4  Register = 0x0E
	REG_RX_ADDR_P5  Register = 0x0F
	REG_TX_ADDR     Register = 0x10
	REG_RX_PW_P0    Register = 0x11
	REG_RX_PW_P1    Register = 0x12
	REG_RX_PW_P2    Register = 0x13
	REG_RX_PW_P3    Register = 0x14
	REG_RX_PW_P4    Register = 0x15
	REG_RX_PW_P5    Register = 0x16
	REG_FIFO_STATUS Register = 0x17
	REG_DYNPD       Register = 0x1C
	REG_FEATURE     Register = 0x1D
	/* P specific */
	REG_RPD Register = 0x09
)

const REGISTER_MASK = 0x1F
const PIPE_MASK = 0x07

/* Bit Mnemonics */
const (
	MASK_RX_DR  = 6
	MASK_TX_DS  = 5
	MASK_MAX_RT = 4
	EN_CRC      = 3
	CRCO        = 2
	PWR_UP      = 1
	PRIM_RX     = 0
	ENAA_P5     = 5
	ENAA_P4     = 4
	ENAA_P3     = 3
	ENAA_P2     = 2
	ENAA_P1     = 1
	ENAA_P0     = 0
	ERX_P5      = 5
	ERX_P4      = 4
	ERX_P3      = 3
	ERX_P2      = 2
	ERX_P1      = 1
	ERX_P0      = 0
	AW          = 0
	ARD         = 4
	ARC         = 0
	PLL_LOCK    = 4
	RF_DR       = 3
	RF_PWR      = 6
	RX_DR       = 6
	TX_DS       = 5
	MAX_RT      = 4
	RX_P_NO     = 1
	TX_FULL     = 0
	PLOS_CNT    = 4
	ARC_CNT     = 0
	TX_REUSE    = 6
	FIFO_FULL   = 5
	TX_EMPTY    = 4
	RX_FULL     = 1
	RX_EMPTY    = 0
	DPL_P5      = 5
	DPL_P4      = 4
	DPL_P3      = 3
	DPL_P2      = 2
	DPL_P1      = 1
	DPL_P0      = 0
	EN_DPL      = 2
	EN_ACK_PAY  = 1
	EN_DYN_ACK  = 0
	/* P specific */
	RF_DR_LOW   = 5
	RF_DR_HIGH  = 3
	RF_PWR_LOW  = 1
	RF_PWR_HIGH = 2
)

type Command uint8

/* Command Mnemonics */
const (
	R_REGISTER    Command = 0x00
	W_REGISTER    Command = 0x20
	ACTIVATE      Command = 0x50
	R_RX_PL_WID   Command = 0x60
	R_RX_PAYLOAD  Command = 0x61
	W_TX_PAYLOAD  Command = 0xA0
	W_ACK_PAYLOAD Command = 0xA8
	FLUSH_TX      Command = 0xE1
	FLUSH_RX      Command = 0xE2
	REUSE_TX_PL   Command = 0xE3
	NOP           Command = 0xFF
	/* P specific */
	W_TX_PAYLOAD_NO_ACK Command = 0xB0
)

const (
	MOSI = 0
	MISO = 1
)

type SpiTx struct {
	Data [2]byte
}

func InterpretTransaction(txs []SpiTx) (ret string, err error) {
	if len(txs) == 0 {
		err = errors.New("At least 1 SPI TX is needed")
		return
	}
	// Parse first MOSI byte to interprete Command
	switch txs[0].Data[MOSI] {
	case byte(FLUSH_TX):
		ret = "Flush TX FIFO"
	case byte(FLUSH_RX):
		ret = "Flush RX FIFO"
	case byte(REUSE_TX_PL):
		ret = "Reuse last transmitted payload"
	case byte(NOP):
		ret = "NOP"
	case byte(R_RX_PL_WID):
		ret = "RX payload width"
		if len(txs) != 2 {
			err = errors.New("2 SPI TX are needed")
			return
		}
		ret += " : " + strconv.Itoa(int(txs[1].Data[MISO]))
	case byte(R_RX_PAYLOAD):
		ret = "Read RX-payload"
		/* 1 to 32 bytes Data */
		if len(txs) < 2 || len(txs) > 33 {
			err = errors.New("Between 2 and 33 SPI TX are needed")
			return
		}
		ret += dumpPayload(txs, MISO)
	case byte(W_TX_PAYLOAD):
		ret = "Write TX-payload"
		/* 1 to 32 bytes Data */
		if len(txs) < 2 || len(txs) > 33 {
			err = errors.New("Between 2 and 33 SPI TX are needed")
			return
		}
		ret += dumpPayload(txs, MOSI)
	case byte(W_TX_PAYLOAD_NO_ACK):
		ret = "Write TX-payload without ACK"
		/* 1 to 32 bytes Data */
		if len(txs) < 2 || len(txs) > 33 {
			err = errors.New("Between 2 and 33 SPI TX are needed")
			return
		}
		ret += dumpPayload(txs, MOSI)
	default:
		/* Special case of Command containing value */
		if (txs[0].Data[MOSI] &^ REGISTER_MASK) == byte(R_REGISTER) {
			ret = "Read Register " + getRegisterName(Register(txs[0].Data[MOSI]&REGISTER_MASK))
			/* 1 to 5 bytes Data */
			if len(txs) < 2 || len(txs) > 6 {
				err = errors.New("Between 2 and 6 SPI TX are needed")
				return
			}
			ret += dumpPayload(txs, MISO)
		} else {
			if (txs[0].Data[MOSI] &^ REGISTER_MASK) == byte(W_REGISTER) {
				ret = "Write Register " + getRegisterName(Register(txs[0].Data[MOSI]&REGISTER_MASK))
				/* 1 to 5 bytes Data */
				if len(txs) < 2 || len(txs) > 6 {
					err = errors.New("Between 2 and 6 SPI TX are needed")
					return
				}
				ret += dumpPayload(txs, MOSI)
			} else {
				if (txs[0].Data[MOSI] &^ PIPE_MASK) == byte(W_ACK_PAYLOAD) {
					ret = "Write ACK Payload for pipe 0x" + strconv.FormatUint(uint64(txs[0].Data[MOSI]&PIPE_MASK), 16)
					/* 1 to 32 bytes Data */
					if len(txs) < 2 || len(txs) > 33 {
						err = errors.New("Between 2 and 33 SPI TX are needed")
						return
					}
					ret += dumpPayload(txs, MOSI)
				} else {
					err = errors.New("Unknown Command")
				}
			}
		}
	}
	return
}

func dumpPayload(txs []SpiTx, stream int) (ret string) {
	ret += " : 0x"
	for i := 1; i < len(txs); i++ {
		if txs[i].Data[stream] < 16 {
			ret += "0"
		}
		ret += strconv.FormatUint(uint64(txs[i].Data[stream]), 16)
	}
	return
}

func getRegisterName(reg Register) (name string) {
	switch reg {
	case REG_NRF_CONFIG:
		name = "NRF_CONFIG"
	case REG_EN_AA:
		name = "EN_AA"
	case REG_EN_RXADDR:
		name = "EN_RXADDR"
	case REG_SETUP_AW:
		name = "SETUP_AW"
	case REG_SETUP_RETR:
		name = "SETUP_RETR"
	case REG_RF_CH:
		name = "RF_CH"
	case REG_RF_SETUP:
		name = "RF_SETUP"
	case REG_NRF_STATUS:
		name = "NRF_STATUS"
	case REG_OBSERVE_TX:
		name = "OBSERVE_TX"
	case REG_CD:
		name = "CD"
	case REG_RX_ADDR_P0:
		name = "RX_ADDR_P0"
	case REG_RX_ADDR_P1:
		name = "RX_ADDR_P1"
	case REG_RX_ADDR_P2:
		name = "RX_ADDR_P2"
	case REG_RX_ADDR_P3:
		name = "RX_ADDR_P3"
	case REG_RX_ADDR_P4:
		name = "RX_ADDR_P4"
	case REG_RX_ADDR_P5:
		name = "RX_ADDR_P5"
	case REG_TX_ADDR:
		name = "TX_ADDR"
	case REG_RX_PW_P0:
		name = "RX_PW_P0"
	case REG_RX_PW_P1:
		name = "RX_PW_P1"
	case REG_RX_PW_P2:
		name = "RX_PW_P2"
	case REG_RX_PW_P3:
		name = "RX_PW_P3"
	case REG_RX_PW_P4:
		name = "RX_PW_P4"
	case REG_RX_PW_P5:
		name = "RX_PW_P5"
	case REG_FIFO_STATUS:
		name = "FIFO_STATUS"
	case REG_DYNPD:
		name = "DYNPD"
	case REG_FEATURE:
		name = "FEATURE"
	default:
		name = "Unknown"
	}
	return
}
