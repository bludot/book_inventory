package jwt

import (
	"github.com/bludot/tempmee/inventory/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJWTService_GenerateToken(t *testing.T) {
	t.Run("should generate token", func(t *testing.T) {
		a := assert.New(t)

		cfg := config.LoadConfigOrPanic()
		jwtService := NewJWTService(cfg.JWTConfig, map[string]interface{}{})

		token, err := jwtService.GenerateToken("test", map[string]interface{}{})
		a.NoError(err)
		a.NotNil(token)

		// verify token
		claims, err := jwtService.ValidateToken(token)
		a.NoError(err)
		a.NotNil(claims)

	})

	t.Run("should generate token with claims", func(t *testing.T) {
		a := assert.New(t)

		cfg := config.LoadConfigOrPanic()
		defaultClaims := map[string]interface{}{
			"test": "test",
		}
		jwtService := NewJWTService(cfg.JWTConfig, defaultClaims)

		token, err := jwtService.GenerateToken("test", map[string]interface{}{
			"test2": "test2",
		})
		a.NoError(err)
		a.NotNil(token)

		claims, err := jwtService.ValidateToken(token)
		claimsMap := *claims
		a.NoError(err)
		a.NotNil(claims)
		a.Equal("test", claimsMap["sub"])
		a.Equal("test", claimsMap["test"])
		a.Equal("test2", claimsMap["test2"])

	})

	t.Run("should generate token with default claims and override", func(t *testing.T) {
		a := assert.New(t)

		cfg := config.LoadConfigOrPanic()
		defaultClaims := map[string]interface{}{
			"test": "test",
		}
		jwtService := NewJWTService(cfg.JWTConfig, defaultClaims)

		token, err := jwtService.GenerateToken("test", map[string]interface{}{
			"test": "test2",
		})
		a.NoError(err)
		a.NotNil(token)

		claims, err := jwtService.ValidateToken(token)
		claimsMap := *claims
		a.NoError(err)
		a.NotNil(claims)
		a.Equal("test", claimsMap["sub"])
		a.Equal("test2", claimsMap["test"])

	})
}
