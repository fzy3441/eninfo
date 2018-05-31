package eninfo

const (
	Type_Aes = 0
)

type ieninfo interface {
	enValue(passwd string, value []byte) (*EnData, error)
	deValue(data *EnData) ([]byte, error)
}

type EnData struct {
	Passwd   string
	EnCipher string // 加密密文
}
