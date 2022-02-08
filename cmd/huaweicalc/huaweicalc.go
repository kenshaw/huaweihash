package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/kenshaw/huaweihash"
)

func main() {
	imei := flag.String("imei", "", "imei")
	flag.Parse()
	if err := run(*imei); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(imei string) error {
	// calculate flash, v1, v2, v201 codes
	flash, err := huaweihash.FlashString(imei)
	if err != nil {
		return err
	}
	v1, err := huaweihash.V1String(imei)
	if err != nil {
		return err
	}
	v2, err := huaweihash.V2String(imei)
	if err != nil {
		return err
	}
	v201, err := huaweihash.V201String(imei)
	if err != nil {
		return err
	}
	// format output
	m := map[string]string{
		"flash": strconv.FormatUint(uint64(flash), 10),
		"v1":    strconv.FormatUint(uint64(v1), 10),
		"v2":    strconv.FormatUint(uint64(v2), 10),
		"v201":  strconv.FormatUint(uint64(v201), 10),
	}
	// json encode
	buf, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(append(buf, '\n'))
	return err
}
