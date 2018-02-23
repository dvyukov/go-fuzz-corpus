// Copyright 2015 Dmitry Vyukov. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package flatbuffers

import (
	"github.com/dvyukov/go-fuzz-corpus/flatbuffers/mygame"
	flatbuffers "github.com/google/flatbuffers/go"
)

func Fuzz(data []byte) int {
	var m mygame.Monster
	m.Init(data, 0)
	//loop:
	pos := m.Pos(nil)
	pos.X()
	pos.Y()
	pos.Z()
	m.Mana()
	m.Hp()
	m.Name()
	for i := 0; i < m.InventoryLength(); i++ {
		m.Inventory(i)
	}
	m.Color()
	m.Name()
	var t flatbuffers.Table
	if m.Test(&t) {
		/*
			switch m.TestType() {
			case mygame.AnyMonster:
				m._tab = t
				goto loop
			case mygame.AnyWeapon:
				w := mygame.Weapon{t}
				w.Name()
				w.X()
			}
		*/
	}
	return 0
}
