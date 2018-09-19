package qniblib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBackEnd_SetGet(t *testing.T) {
	be := NewBackEnd("http://127.0.0.1:2379")
	be.SetDevice("127.0.0.1", "0", "healthy")
	devs, err := be.GetDevices()
	assert.NoError(t, err, "GetDevices should return no error")
	assert.Equal(t, 1, len(devs))
}

func TestNewBackEnd_SetGet(t *testing.T) {
	be := NewBackEnd("http://127.0.0.1:2379")
	be.SetDevice("127.0.0.1", "0", "healthy")
	devs, err := be.GetDevices()
	assert.NoError(t, err, "GetDevices should return no error")
	assert.Equal(t, 1, len(devs))
}
