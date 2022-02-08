package huaweihash

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"unsafe"
)

/*
#include "encrypt.h"
*/
import "C"

// encrypt calculates the version 1 encryption hash for Huawei devices.
func encrypt(imei, key []byte) (uint32, error) {
	if len(imei) != 15 {
		return 0, errors.New("imei must be length 15")
	}
	// take 8:32 of hex(md5(key))
	hkey := fmt.Sprintf("%x", md5.Sum(key))[8:32]
	// copy imei and hkey into hbuf
	hbuf := make([]byte, 31)
	copy(hbuf, imei)
	copy(hbuf[len(imei):], hkey[:])
	// grab md5(hbuf)
	h := md5.Sum(hbuf)
	// compute intermediary result (checksum)
	ir := make([]byte, 4)
	for i := 0; i < 4; i++ {
		ir[3-i] = h[i] ^ h[i+4] ^ h[i+8] ^ h[i+12]
	}
	// convert ir to uint32
	r := bytes.NewReader(ir)
	var res uint32
	if err := binary.Read(r, binary.LittleEndian, &res); err != nil {
		return 0, err
	}
	// last bit manipulation
	return (res | 0x2000000) & 0x3ffffff, nil
}

func calc2(imei []byte) (uint32, error) {
	if len(imei) != 15 {
		return 0, errors.New("imei must be length 15")
	}
	buf := make([]byte, 40)
	C.calc2((*C.char)(unsafe.Pointer(&imei[0])), (*C.char)(unsafe.Pointer(&buf[0])))
	u, err := strconv.ParseUint(string(buf[:bytes.IndexByte(buf, 0x00)]), 10, 32)
	return uint32(u), err
}

func calc201(imei []byte) (uint32, error) {
	if len(imei) != 15 {
		return 0, errors.New("imei must be length 15")
	}
	buf := make([]byte, 40)
	C.calc201((*C.char)(unsafe.Pointer(&imei[0])), (*C.char)(unsafe.Pointer(&buf[0])))
	u, err := strconv.ParseUint(string(buf[:bytes.IndexByte(buf, 0x00)]), 10, 32)
	return uint32(u), err
}

// procIdx.
//func procIdx(imei []byte, ver int) int {
//	csum := 0 // checksum
//
//	//int c1, ch;
//	//long long cx;
//
//	for i := 0; i < len(imei); i++ {
//		ch = imei[i]
//		if ver == 201 {
//			csum += ((ch + i + 1) * ch) * (ch + 0x139)
//		} else {
//			csum += ((ch + i + 1) * (i + 1))
//		}
//	}
//
//	cx = (int64(csum) * -0x6db6db6d) >> 32
//	c1 = ((cx + csum) >> 2) - (csum >> 31)
//
//	return csum - ((c1 << 3) - c1)
//}
