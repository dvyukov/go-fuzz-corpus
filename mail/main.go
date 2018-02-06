// Copyright 2015 Dmitry Vyukov. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package mail

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/mail"
	"reflect"
	"strings"
)

func Fuzz(data []byte) int {
	msg, err := mail.ReadMessage(bytes.NewReader(data))
	if err != nil {
		return 0
	}
	msg.Header.AddressList("to")
	msg.Header.Date()
	if addr, err := mail.ParseAddress(msg.Header.Get("from")); err == nil {
		addr1, err := mail.ParseAddress(addr.String())
		if err != nil {
			panic(err)
		}
		if !reflect.DeepEqual(addr, addr1) {
			panic("addr changed")
		}
	}
	io.Copy(ioutil.Discard, msg.Body)
	return 1
}

func FuzzParseAddressList(data []byte) int {
	list, err := mail.ParseAddressList(string(data))
	if err != nil {
		return 0
	}
	var addrs []string
	for _, addr := range list {
		addrs = append(addrs, addr.String())
	}
	list1, err := mail.ParseAddressList(strings.Join(addrs, ","))
	if err != nil {
		panic(err)
	}
	if !reflect.DeepEqual(list, list1) {
		panic("list changed")
	}
	return 1
}
