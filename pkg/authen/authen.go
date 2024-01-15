package authen

import (
	"errors"
	"math"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthFactory interface {
	SignToken() string
}

type Claims struct {
	PlayerId string `json:"player_id"`
}

type AuthMapClaims struct {
	*Claims
	jwt.RegisteredClaims
}

type authConcrete struct {
	Secret []byte
	Claims *AuthMapClaims `json:"claims"`
}

type accessToken struct {
	*authConcrete
}

type refreshToken struct {
	*authConcrete
}

// type refreshToken struct{ *authConcrete }

// type apikey struct{ *authConcrete }

func now() time.Time {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return time.Now().In(loc)
}

func (a *authConcrete) SignToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a.Claims)
	// only byte
	ss, _ := token.SignedString(a.Secret)

	return ss
}

func jwtTimeDurationCal(t int64) *jwt.NumericDate {
	return jwt.NewNumericDate(now().Add(time.Duration(t * int64(math.Pow10(9)))))
}

// func jwtTimeRepeatAdapter(t int64) *jwt.NumericDate {
// 	return jwt.NewNumericDate(time.Unix(t, 0))
// }

func NewAccessToken(secret string, expiredAt int64, claims *Claims) AuthFactory {

	return &accessToken{
		authConcrete: &authConcrete{
			Secret: []byte(secret),
			Claims: &AuthMapClaims{
				Claims: claims,
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwtTimeDurationCal(expiredAt),
					NotBefore: jwt.NewNumericDate(now()),
					IssuedAt:  jwt.NewNumericDate(now()),
				},
			},
		},
	}
}

func NewRefreshToken(secret string, expiredAt int64, claims *Claims) AuthFactory {
	return &refreshToken{
		authConcrete: &authConcrete{
			Secret: []byte(secret),
			Claims: &AuthMapClaims{
				Claims: claims,
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwtTimeDurationCal(expiredAt),
					NotBefore: jwt.NewNumericDate(now()),
					IssuedAt:  jwt.NewNumericDate(now()),
				},
			},
		},
	}
}

func ParseToken(secret string, tokenstring string) (*AuthMapClaims, error) {

	token, err := jwt.ParseWithClaims(tokenstring, &AuthMapClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error: unexpected signing method")
		}

		return []byte(secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, errors.New("error: token format is invalid")
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("error: token is expired")
		} else {
			return nil, errors.New("error: token is invalid")
		}
	}

	if claims, ok := token.Claims.(*AuthMapClaims); ok {
		return claims, nil
	} else {
		return nil, errors.New("error: claims type is invalid")

	}
}
