package Utils

import (
	"github.com/xlzd/gotp"
)

var totp = gotp.NewDefaultTOTP(gotp.RandomSecret(16))

func GetOTP(timeStamp int64) string {
	return totp.At(int(timeStamp))
}

func ValidateOTP(code string, timeStamp int64) bool {
	return totp.Verify(code, int(timeStamp))
}
