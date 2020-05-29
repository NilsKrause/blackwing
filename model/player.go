package model

type Player struct {
	ID         int                  // discord player id
	Characters map[string]Character // each character gets a uuid to identify them
}

func NewPlayer(id int) *Player {
	return &Player{ID: id, Characters: make(map[string]Character, 0)}
}

func (p *Player) HasChar(name string) bool {

}
