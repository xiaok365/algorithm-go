package collections

import (
	"crypto/md5"
)

type KetamaHash struct {
}

func NewKetamaHash() *KetamaHash {
	return &KetamaHash{}
}

func (t *KetamaHash) GetHash(str string) uint {
	digest := md5.Sum([]byte(str))
	h := 0
	ret := uint(digest[3+h*4])<<24 | uint(digest[2+h*4])<<16 | uint(digest[1+h*4])<<8 | uint(digest[h*4])
	return ret
}
