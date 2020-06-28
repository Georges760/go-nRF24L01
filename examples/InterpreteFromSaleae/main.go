package main

import (
	"flag"
	"fmt"
	"github.com/Georges760/go-nRF24L01"
	"github.com/Georges760/go-saleae"
	"log"
	"os"
)

func main() {
	// Check args
	pfile := flag.String("f", "", "CSV file from Saleae Decoded Protocol export")
	var cmdTimeout float64
	flag.Float64Var(&cmdTimeout, "t", 0.002, "Command Timeout in Second")
	var debug bool
	flag.BoolVar(&debug, "d", false, "Enable DEBUG mode")
	flag.Parse()
	if *pfile == "" {
		flag.PrintDefaults()
		log.Fatal("-f arg is mandatory")
	}
	// Open CSV file
	recordFile, err := os.Open(*pfile)
	if err != nil {
		log.Fatal("Error opening CSV file:", err)
	}
	// Parse it
	spi, err := saleae.ParseCSV(recordFile)
	if err != nil {
		log.Fatal("Error parsing CSV file:", err)
	}
	// Interpret as nRF24L01 Commands
	spi24 := []nrf24l01.SpiTx{}
	lastTxSecond := float64(spi[0].Second)
	for _, tx := range spi {
		if tx.Second-lastTxSecond > cmdTimeout {
			if debug {
				log.Println("spi24:", spi24)
			}
			txt, err := nrf24l01.InterpretTransaction(spi24)
			fmt.Println(txt)
			if err != nil {
				log.Println(err)
			}
			spi24 = nil
		}
		spi24 = append(spi24, nrf24l01.SpiTx{[2]byte{byte(tx.Mosi), byte(tx.Miso)}})
		lastTxSecond = tx.Second
	}
	txt, err := nrf24l01.InterpretTransaction(spi24)
	fmt.Println(txt)
	if err != nil {
		log.Println(err)
	}
	// Exit
	os.Exit(0)
}
