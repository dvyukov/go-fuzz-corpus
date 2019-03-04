// Copyright 2015 go-fuzz project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package gopacket

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func FuzzDNS(data []byte) int {
	gopacket.NewPacket(data, layers.LayerTypeDNS, gopacket.Default)
	return 0
}
