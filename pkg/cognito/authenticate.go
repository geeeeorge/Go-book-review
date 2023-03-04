package cognito

import (
	"fmt"
	"sync"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"
)

var (
	ErrInvalidJWT = errors.New("invalid jwt")
)

type IDToken struct {
	Sub      string
	Iss      string
	Username string
	Iat      time.Time
}

var (
	jwks *keyfunc.JWKS
	once singleflight.Group
	mu   sync.Mutex
)

func cognitoJwkEndpoint(region, userPoolID string) string {
	return fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", region, userPoolID)
}

// NewValid middlewareで使用される
func NewValid(region, poolID, iss, raw string, now time.Time) (IDToken, error) {
	mu.Lock()
	if jwks == nil {
		_, err, _ := once.Do(iss, func() (v interface{}, err error) {
			jwks, err = keyfunc.Get(cognitoJwkEndpoint(region, poolID), keyfunc.Options{})
			return nil, err
		})
		if err != nil {
			once.Forget(iss)
			return IDToken{}, err
		}
	}
	mu.Unlock()

	tok, err := jwt.Parse(raw, jwks.Keyfunc, jwt.WithoutClaimsValidation())
	if err != nil {
		return IDToken{}, err
	}
	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return IDToken{}, ErrInvalidJWT
	}
	if !claims.VerifyExpiresAt(now.Unix(), true) {
		return IDToken{}, fmt.Errorf("%w: expired", ErrInvalidJWT)
	}
	if !claims.VerifyIssuer(iss, true) {
		return IDToken{}, fmt.Errorf("%w: unknown issuer", ErrInvalidJWT)
	}
	if claims["sub"] == nil || claims["iss"] == nil || claims["cognito:username"] == nil {
		return IDToken{}, fmt.Errorf("%w: invalid claim", ErrInvalidJWT)
	}

	return IDToken{
		Sub:      claims["sub"].(string),
		Iss:      claims["iss"].(string),
		Username: claims["cognito:username"].(string),
		Iat:      time.Unix(int64(claims["iat"].(float64)), 0),
	}, nil
}
