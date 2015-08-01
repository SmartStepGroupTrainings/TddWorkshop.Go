package casino

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestPlayer_ByDefault_IsNotInGame(t *testing.T) {
    player := Player{}

    assert.False(t, player.IsInGame())
}
