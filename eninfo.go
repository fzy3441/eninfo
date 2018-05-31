package eninfo

import ()

func EnValue(passwd string, value []byte, enType int) (*EnData, error) {
	obj := _getEnObj(enType)
	l := len(value)
	head, err := _int2Bytes(l)
	if err != nil {
		return nil, err
	}

	value = append(head, value...)

	l_relenish := 64 - ((l + 4) % 64)

	if l_relenish > 0 {
		replenish := make([]byte, l_relenish, l_relenish)
		value = append(value, replenish...)

	}

	return obj.enValue(passwd, value)
}

func DeValue(passwd string, value string, enType int) ([]byte, error) {
	obj := _getEnObj(enType)

	b_value, err := obj.deValue(&EnData{
		Passwd:   passwd,
		EnCipher: value,
	})
	if err != nil {
		return nil, err
	}

	l, err := _bytes2Dec(b_value[0:4])
	if err != nil {
		return nil, err
	}
	return b_value[4:l], nil

}
