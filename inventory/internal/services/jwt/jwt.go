package jwt

import (
	"github.com/bludot/tempmee/inventory/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const minJWTTokenValidityMinutes = 15

type JWTServiceImpl interface {
	GenerateToken(subject string, claims map[string]interface{}) (string, error)
	ValidateToken(token string) (*map[string]interface{}, error)
}

type JWTService struct {
	Config config.JWTConfig
	TTL    time.Duration
	Claims jwt.MapClaims
}

func NewJWTService(cfg config.JWTConfig, claims jwt.MapClaims) JWTServiceImpl {
	return &JWTService{
		Config: cfg,
		TTL:    time.Minute * time.Duration(cfg.Expire),
		Claims: claims,
	}
}

func getDefault(duration *time.Duration, defaultDuration time.Duration) time.Duration {
	if duration == nil {
		return defaultDuration
	}

	return *duration
}

func (a *JWTService) applyClaims(claims map[string]interface{}) {
	for k, v := range claims {
		a.Claims[k] = v
	}
	a.Claims["iss"] = a.Config.Issuer
	a.Claims["aud"] = a.Config.Audience
	a.Claims["iat"] = time.Now().Unix()
	a.Claims["nbf"] = time.Now().Unix()
	a.Claims["exp"] = time.
		Now().
		Add(getDefault(&a.TTL, time.Minute*time.Duration(minJWTTokenValidityMinutes))).
		Unix()
}

func (a *JWTService) GenerateToken(subject string, claims map[string]interface{}) (string, error) {
	a.applyClaims(claims)
	a.Claims["sub"] = subject

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a.Claims)
	signedToken, err := token.SignedString([]byte(a.Config.PrivateKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (a *JWTService) ValidateToken(token string) (*map[string]interface{}, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.Config.PrivateKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !t.Valid {
		return nil, err
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	_, ok = claims["sub"].(string)
	if !ok {
		return nil, err
	}

	if claims["exp"].(float64) < float64(time.Now().Unix()) {
		return nil, err
	}

	// convert claims to map[string]interface{}
	claimsMap := make(map[string]interface{})
	for k, v := range claims {
		claimsMap[k] = v
	}

	return &claimsMap, nil
}
