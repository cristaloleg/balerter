package discord

import (
	"testing"

	"github.com/balerter/balerter/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	d, err := New(&config.ChannelDiscord{Name: "foo", Token: "123"}, nil)
	require.NoError(t, err)
	assert.IsType(t, &Discord{}, d)
	assert.Equal(t, "foo", d.name)
	assert.Equal(t, "123", d.conf.Token)
}

func TestName(t *testing.T) {
	d := &Discord{name: "foo"}
	assert.Equal(t, "foo", d.Name())
}
