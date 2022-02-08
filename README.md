# About huaweicalc

Package huaweicalc provides a simple Go package to calculate the unlock codes
for a Huawei device. These are used for flashing the firmware to different versions.

Originally, I was planning to rewrite all the hash calculations to Go, but gave
up on the v2 and v201 algorithms. As such, I merely imported the C code from
[forth32's GitHub repository](https://github.com/forth32/huaweicalc)

# Installation

Install in the usual way:

```sh
$ go get -u github.com/knq/huaweihash
```

# Command-Line Util
There is a command line tool that uses the package to easily generate the
flash, v1, v2, and v202 codes:
```sh
$ go get -u github.com/knq/huaweihash/cmd/huaweicalc
$ huaweicalc -imei 868757025499999
{
  "flash": "50702788",
  "v1": "48125080",
  "v2": "39842371",
  "v201": "46863554"
}
```
