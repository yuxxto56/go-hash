package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"log"
)

type Hash struct {
	Str string
}

//MD5加密
func (h *Hash)HashMd5() string{
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(h.Str))
	Result := Md5Inst.Sum([]byte(""))
	return fmt.Sprintf("%x",Result)
}

//sha1加密
func (h *Hash)HashSha1() string{
	sha1Inst := sha1.New()
	_,err := sha1Inst.Write([]byte(h.Str))
	if err != nil{
		log.Fatal(err.Error())
	}
	Result := sha1Inst.Sum([]byte(""))
	return fmt.Sprintf("%x",Result)
}

func init() {
	//todo
}
