package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)
	hashPassword, err := HashPassword(password)
	require.NoError(t, err)
	fmt.Println(password)
	require.NotEmpty(t, hashPassword)
	require.Equal(t, len(password), 6)
}
