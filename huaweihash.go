package huaweihash

// Huawei identifiers.
const (
	E630Upgrade    = "e630upgrade"
	Hwe620Datacard = "hwe620datacard"
	NotAvailable   = "[not available]"
)

// Flash calculates the a flash code for a specified imei.
func Flash(imei []byte) (uint32, error) {
	return encrypt(imei, []byte(E630Upgrade))
}

// FlashString calculates the a flash code for a specified imei.
func FlashString(imei string) (uint32, error) {
	return Flash([]byte(imei))
}

// V1 calculates the version 1 code for a specified imei.
func V1(imei []byte) (uint32, error) {
	return encrypt(imei, []byte(Hwe620Datacard))
}

// V1String calculates the version 1 code for a specified imei.
func V1String(imei string) (uint32, error) {
	return V1([]byte(imei))
}

// V2 calculates the version 2 code for a specified imei.
func V2(imei []byte) (uint32, error) {
	return calc2(imei)
}

// V2String calculates the version 2 code for a specified imei.
func V2String(imei string) (uint32, error) {
	return V2([]byte(imei))
}

// V201 calculates the version 201 code for a specified imei.
func V201(imei []byte) (uint32, error) {
	return calc201(imei)
}

// V201String calculates the version 201 code for a specified imei.
func V201String(imei string) (uint32, error) {
	return V201([]byte(imei))
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
