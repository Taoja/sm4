package main

import (
	"fmt"
	"sm4/sm4"
)

func main(){
	data := []byte("H5前端子应用HMBA，wasm项目")
	key := []byte("1234567890abcdef")
	ctx := sm4.NewSM4(key)
	enc := ctx.DoEncrypt(data)
	fmt.Printf("密文 = %x\n", enc)
	dec := ctx.DoDecrypt(enc)
	fmt.Printf("明文 = %s\n", dec)
}

