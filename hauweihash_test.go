package huaweihash

import (
	"testing"
)

func TestCodes(t *testing.T) {
	for i, test := range []struct {
		imei  string
		flash uint32
		v1    uint32
		v2    uint32
		v201  uint32
	}{
		{"868757025499999", 50702788, 48125080, 39842371, 46863554},
	} {
		switch u, err := FlashString(test.imei); {
		case err != nil:
			t.Errorf("test %d expected no error, got: %v", i, err)
		default:
			if exp := test.flash; u != exp {
				t.Errorf("test %d expected Flash to be %d, got: %d", i, exp, u)
			}
		}
		switch u, err := V1String(test.imei); {
		case err != nil:
			t.Errorf("test %d expected no error, got: %v", i, err)
		default:
			if exp := test.v1; u != exp {
				t.Errorf("test %d expected V1 to be %d, got: %d", i, exp, u)
			}
		}
		switch u, err := V2String(test.imei); {
		case err != nil:
			t.Errorf("test %d expected no error, got: %v", i, err)
		default:
			if exp := test.v2; u != exp {
				t.Errorf("test %d expected V2 to be %d, got: %d", i, exp, u)
			}
		}
		switch u, err := V201String(test.imei); {
		case err != nil:
			t.Errorf("test %d expected no error, got: %v", i, err)
		default:
			if exp := test.v201; u != exp {
				t.Errorf("test %d expected V201 to be %d, got: %d", i, exp, u)
			}
		}
	}
}
