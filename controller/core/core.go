package core

import "github.com/NilsKrause/blackwing/model"

var _internal = &core{Players: make(map[int]*model.Player, 0)}

func RegisterPlayer(id int) {
	if _, ok := _internal.Players[id]; !ok {
		_internal.Players[id] = model.NewPlayer(id)
	}
}
