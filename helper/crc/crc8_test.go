package crc

import "testing"

func TestCrc8(t *testing.T) {
	data := []byte{0xFA, 0x01, 0xF6, 0x02, 0x80, 0x02}
	t.Log(crc8(data))
}
