package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/knq/huaweihash"
)

var (
	flagImei = flag.String("imei", "", "imei")
)

func main() {
	flag.Parse()

	// calculate flash code
	flashCode, err := huaweihash.Flash([]byte(*flagImei))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// calculate v1 code
	v1Code, err := huaweihash.V1([]byte(*flagImei))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// calculate v2 code
	v2Code, err := huaweihash.V2([]byte(*flagImei))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// calculate v201 code
	v201Code, err := huaweihash.V201([]byte(*flagImei))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// format output
	m := map[string]string{
		"flash": string(flashCode),
		"v1":    string(v1Code),
		"v2":    string(v2Code),
		"v201":  string(v201Code),
	}

	// json encode
	buf, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%s\n", string(buf))
}
