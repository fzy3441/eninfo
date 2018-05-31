package eninfo

import (
	"crypto/aes"
	"crypto/md5"

	"crypto/cipher"
	"github.com/btcsuite/btcutil/base58"
	"github.com/fzy3441/nlog"
	// "time"
)

type Aes struct {
}

func NewAes() *Aes {
	return &Aes{}
}

//  加密私钥
func (obj *Aes) enValue(passwd string, value []byte) (*EnData, error) {
	b_pass := _g32md5([]byte(passwd))
	iv := _g16md5(append(b_pass, []byte("default")...)) // 生成加密参数

	cipher_value, err := _en_cbc(b_pass, iv, value)
	if err != nil {
		return nil, err
	}
	value_base := base58.Encode(cipher_value)

	return &EnData{
		Passwd:   passwd,
		EnCipher: value_base,
	}, nil
}

// 解密信息
func (obj *Aes) deValue(data *EnData) ([]byte, error) {
	b_pass := _g32md5([]byte(data.Passwd))
	iv := _g16md5(append(b_pass, []byte("default")...)) // 生成加密参数

	cipher_detail := base58.Decode(data.EnCipher)
	return _de_cbc(b_pass, iv, cipher_detail)
}

// 得到16位md5加密信息
func _g16md5(value []byte) []byte {
	h := md5.New()
	h.Write(value)
	return h.Sum(nil)
}

// 得到32位md5加密信息
func _g32md5(value []byte) []byte {
	Rpass := _g16md5(value)
	Lpass := _g16md5(append(value, []byte("fzyun")...))

	return append(Lpass, Rpass...)
}

// 加密
func _en_cbc(key, iv, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		nlog.Error(err)
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

// 解密
func _de_cbc(key, iv, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		nlog.Error(err)
		return nil, err
	}
	ciphertext = ciphertext[aes.BlockSize:]
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	return ciphertext, nil
}
