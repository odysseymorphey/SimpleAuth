package services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"testing"
	"time"

	"github.com/odysseymorphey/SimpleAuth/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGeneratePairID(t *testing.T) {
	t.Run("valid format", func(t *testing.T) {
		uuid, err := GeneratePairID()
		require.NoError(t, err)

		parts := strings.Split(uuid, "-")
		require.Len(t, parts, 5)
		assert.Len(t, parts[0], 8)
		assert.Len(t, parts[1], 4)
		assert.Len(t, parts[2], 4)
		assert.Len(t, parts[3], 4)
		assert.Len(t, parts[4], 12)
	})

	t.Run("unique values", func(t *testing.T) {
		uuid1, err := GeneratePairID()
		require.NoError(t, err)

		uuid2, err := GeneratePairID()
		require.NoError(t, err)

		assert.NotEqual(t, uuid1, uuid2)
	})
}

func TestGenerateBCrypt(t *testing.T) {
	t.Run("valid token hash", func(t *testing.T) {
		token := "test_token"
		hash, err := generateBCrypt(token)
		require.NoError(t, err)

		err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(token))
		assert.NoError(t, err)
	})

	t.Run("empty token", func(t *testing.T) {
		_, err := generateBCrypt("")
		assert.Error(t, err)
	})
}

func TestGenerateAccessToken(t *testing.T) {
	uInfo := &models.UserInfo{
		GUID:   "test_guid",
		UserIP: "192.168.1.1",
	}
	pairID := "test_pair_id"

	tokenString, err := generateAccessToken(uInfo, pairID)
	require.NoError(t, err)

	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	require.NoError(t, err)

	claims, ok := token.Claims.(*models.CustomClaims)
	require.True(t, ok)

	assert.Equal(t, uInfo.GUID, claims.GUID)
	assert.Equal(t, uInfo.UserIP, claims.IP)
	assert.Equal(t, pairID, claims.PairID)
	assert.WithinDuration(t, time.Now().Add(5*time.Minute), claims.ExpiresAt.Time, time.Second)
}

func TestGenerateRefreshToken(t *testing.T) {
	t.Run("valid token", func(t *testing.T) {
		token, err := generateRefreshToken()
		require.NoError(t, err)

		assert.Len(t, token, 44) // 32 bytes => base64.URLEncoding is 44 chars
	})

	t.Run("unique tokens", func(t *testing.T) {
		token1, err := generateRefreshToken()
		require.NoError(t, err)

		token2, err := generateRefreshToken()
		require.NoError(t, err)

		assert.NotEqual(t, token1, token2)
	})
}

func TestParseToken(t *testing.T) {
	t.Run("valid token", func(t *testing.T) {
		expectedClaims := &models.CustomClaims{
			GUID:   "test_guid",
			IP:     "10.0.0.1",
			PairID: "test_pair",
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, expectedClaims)
		tokenString, err := token.SignedString([]byte("different_secret"))
		require.NoError(t, err)

		parsedClaims, err := parseToken(tokenString)
		require.NoError(t, err)

		assert.Equal(t, expectedClaims.GUID, parsedClaims.GUID)
		assert.Equal(t, expectedClaims.IP, parsedClaims.IP)
		assert.Equal(t, expectedClaims.PairID, parsedClaims.PairID)
	})

	t.Run("invalid token", func(t *testing.T) {
		_, err := parseToken("invalid.token.here")
		assert.Error(t, err)
	})
}
