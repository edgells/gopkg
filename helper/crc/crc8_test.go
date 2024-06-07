package crc

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
)

func Crc8Bit(data []byte) []byte {
	var byteData byte
	for _, val := range data {
		byteData += val
	}

	return append(data, byteData)
}

func TestCrc8(t *testing.T) {
	data := []byte{0xFA, 0x01, 0xF6, 0x02, 0x80, 0x02}
	t.Log(crc8(data))
}

func TestCrc8bit(t *testing.T) {
	data := []byte{0xFA, 0x01, 0xFD, 0x80, 0xF5, 0xFF, 0x00, 0x00, 0x03, 0x20}

	t.Logf("%x\n", Crc8Bit(data))
}

func TestCrc8Pos(t *testing.T) {
	hexString := "FA 01 FD 02 80 02 00 FA 00"
	hexString = strings.ReplaceAll(hexString, " ", "")
	data, err := hex.DecodeString(hexString)
	if err != nil {
		return
	}

	t.Log(fmt.Sprintf("%x", Crc8Bit(data)))
}
