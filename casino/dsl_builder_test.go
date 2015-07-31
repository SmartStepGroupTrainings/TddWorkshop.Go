package casino_new

type Builder struct {
}

var create *Builder

func (builder *Builder) Player() *PlayerBuilder {
	return &PlayerBuilder{}
}

func (builder *Builder) Game() *GameBuilder {
	return &GameBuilder{}
}
