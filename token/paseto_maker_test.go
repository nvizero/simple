package token

import (
	"simple/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPaseto(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	token, _, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload2, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload2.Username)

	require.NotZero(t, payload2.ID)
	require.Equal(t, username, payload2.Username)
}

func TestExpirePasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	token, _, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
