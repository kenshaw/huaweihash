package huaweihash

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"errors"
	"fmt"
)

// encrypt calculates the version 1 encryption hash for Huawei devices.
func encrypt(imei, key []byte) (string, error) {
	if len(imei) != 15 {
		return "", errors.New("imei must be length 15")
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
	err := binary.Read(r, binary.LittleEndian, &res)
	if err != nil {
		return "", err
	}

	// last bit manipulation
	res |= 0x2000000
	res &= 0x3FFFFFF

	return fmt.Sprintf("%d", res), nil
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
