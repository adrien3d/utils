package utils

import (
	"github.com/snwfdhmp/errlog"
)

func CheckErr(e error) {
	if e != nil {
		errlog.Debug(e)
		panic(e)
	}
}
