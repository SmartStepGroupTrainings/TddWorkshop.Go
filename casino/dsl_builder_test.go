package casino

var create Builder

type Builder struct {
}

func (builder *Builder) Player() *PlayerBuilder {
	return &PlayerBuilder{}
}

func (builder *Builder) Game() *GameBuilder {
	return &GameBuilder{}
}
