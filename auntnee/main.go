package main

import "github.com/grijul/otpgen"

func main() {
	otp := otpgen.TOTP{
		Secret:    "asdasd",
		Digits:    6,
		Algorithm: "SHA1",
		Period:    30,
	}

	str, err := otp.Generate()
	if err != nil {
		panic(err)
	}
	println(str)
}
