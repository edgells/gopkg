package crc

import (
	"encoding/hex"
)

// crc16Modbus 函数计算给定指令字符串的CRC16校验码
func crc16Modbus(instruction string) (str string, err error) { // 将指令字符串转换为字节数组
	data, err := hex.DecodeString(instruction)
	if err != nil {
		return
	}
	// 初始化CRC值为0xFFFF
	crc := uint16(0xFFFF)
	// 遍历数据字节
	for _, b := range data {
		//将当前字节与CRC值进行异或操作
		crc ^= uint16(b)
		//对CRC值进行8次右移操作
		for i := 0; i < 8; i++ {
			//如果最低位为1，则进行异或运算
			if crc&0x0001 != 0 {
				crc >>= 1
				crc ^= 0xA001
			} else {
				//如果最低位为0，则直接右移
				crc >>= 1
			}
		}
	}
	//将计算出的CRC16值转换为字节数组
	crcBytes := make([]byte, 2)
	crcBytes[0] = byte(crc & 0xFF)
	crcBytes[1] = byte((crc >> 8) & 0xFF)
	//将CRC16校验码转换为十六进制字符串
	return hex.EncodeToString(crcBytes), nil
}

func crc16ModbusWithByte(data []byte) ([]byte, error) { // 将指令字符串转换为字节数组
	// 初始化CRC值为0xFFFF
	crc := uint16(0xFFFF)
	// 遍历数据字节
	for _, b := range data {
		//将当前字节与CRC值进行异或操作
		crc ^= uint16(b)
		//对CRC值进行8次右移操作
		for i := 0; i < 8; i++ {
			//如果最低位为1，则进行异或运算
			if crc&0x0001 != 0 {
				crc >>= 1
				crc ^= 0xA001
			} else {
				//如果最低位为0，则直接右移
				crc >>= 1
			}
		}
	}
	//将计算出的CRC16值转换为字节数组
	crcBytes := make([]byte, 2)
	crcBytes[0] = byte(crc & 0xFF)
	crcBytes[1] = byte((crc >> 8) & 0xFF)
	//将CRC16校验码转换为十六进制字符串
	return crcBytes, nil
}
