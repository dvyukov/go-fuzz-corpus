// Copyright 2020 go-fuzz project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package math

import (
	"math/big"
)

func Fuzz(data []byte) int {
	if len(data) < 3 {
		return 0
	}
	prec := int(data[0])
	roundingMode := big.RoundingMode(data[1])
	base := int(data[2])
	str := string(data[3:])
	_, _, err := big.ParseFloat(str, base, uint(prec), roundingMode)
	if err != nil {
		return 0
	}
	return 1
}
