package manager

import (
	"fmt"
	"github.com/balerter/balerter/internal/corestorage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGet_empty_name(t *testing.T) {
	m := &Manager{}

	_, err := m.Get(" ")

	require.Error(t, err)
	assert.Equal(t, "empty alert name", err.Error())
}

func TestGet_error(t *testing.T) {
	e1 := fmt.Errorf("e1")

	mockStorage := corestorage.NewMock("m")
	mockStorage.AlertMock().On("Get", mock.Anything).Return(nil, e1)

	m := &Manager{
		storage: mockStorage,
	}

	_, err := m.Get("foo")

	require.Error(t, err)
	assert.Equal(t, "error get alert, e1", err.Error())
}
