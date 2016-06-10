package huaweihash

import (
	"bytes"
	"errors"
	"unsafe"
)

/*

#include "encrypt.h"

*/
import "C"

const (
	E630Upgrade    = "e630upgrade"
	Hwe620Datacard = "hwe620datacard"
	NotAvailable   = "[not available]"
)

// Flash calculates the a flash code for a specified imei.
func Flash(imei []byte) (string, error) {
	return encrypt(imei, []byte(E630Upgrade))
}

// V1 calculates the version 1 code for a specified imei.
func V1(imei []byte) (string, error) {
	return encrypt(imei, []byte(Hwe620Datacard))
}

// V2 calculates the version 2 code for a specified imei.
func V2(imei []byte) (string, error) {
	if len(imei) != 15 {
		return "", errors.New("imei must be length 15")
	}

	buf := make([]byte, 40)
	C.calc2((*C.char)(unsafe.Pointer(&imei[0])), (*C.char)(unsafe.Pointer(&buf[0])))

	return string(buf[:bytes.IndexByte(buf, 0x00)]), nil
}

// V201 calculates the version 201 code for a specified imei.
func V201(imei []byte) (string, error) {
	if len(imei) != 15 {
		return "", errors.New("imei must be length 15")
	}

	buf := make([]byte, 40)
	C.calc201((*C.char)(unsafe.Pointer(&imei[0])), (*C.char)(unsafe.Pointer(&buf[0])))

	return string(buf[:bytes.IndexByte(buf, 0x00)]), nil
}

/*func dump(buf []byte) string {
	var b bytes.Buffer

	for i := 0; i < len(buf); i++ {
		b.WriteString(fmt.Sprintf("%x ", buf[i]))
	}

	return string(b.Bytes())
}

func dumpUint(buf []byte) string {
	var b bytes.Buffer

	for i := 0; i < len(buf); i++ {
		b.WriteString(fmt.Sprintf("%d ", buf[i]))
	}

	return string(b.Bytes())
}*/
