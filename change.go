package eninfo

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

// 16进制转byte
func _int2Bytes(i int) ([]byte, error) {
	hex, err := hex.DecodeString(fmt.Sprintf("%08X", i))
	return []byte(hex), err
}

// byte转10进制
func _bytes2Dec(b_hex []byte) (int, error) {
	str_hex := hex.EncodeToString(b_hex)
	base, err := strconv.ParseInt(str_hex, 16, 10)
	return int(base), err
}

func _getEnObj(enType int) (iobj ieninfo) {
	switch enType {
	case 0:
		iobj = NewAes()
	default:
		iobj = NewAes()
	}
	return
}
