package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	db "github.com/xdot2012/simple-bank/db/sqlc"
	"github.com/xdot2012/simple-bank/util"
)

func randomUser() (db.User, string) {
	password := util.RandomString(8)
	hashedPassword, _ := util.HashPassowrd(password)

	return db.User{
		ID:             util.RandomInt(1, 1000),
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}, password
}

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	user, _ := randomUser()
	token, payload, err := maker.CreateToken(user.ID, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, user.ID, payload.UserID)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	user, _ := randomUser()

	require.NoError(t, err)

	token, payload, err := maker.CreateToken(user.ID, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
