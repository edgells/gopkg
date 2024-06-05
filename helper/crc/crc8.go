package main

const (
	firstPadding  = 0xFF
	secondPadding = 0x17B
	threePadding  = 0x7B
)

func crc8(data []byte) (byte, error) {
	var crc byte
	for i := 0; i < len(data); i++ {
		crc ^= data[i]
		for j := 0; j < 8; j++ {
			if crc&byte(0x80) > 0 {
				crc = (crc << 1) ^ 0x07
			} else {
				crc <<= 1
			}
		}
	}

	return crc, nil
}
