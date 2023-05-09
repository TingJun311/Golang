package token

import (
	"crypto/rand"
	"encoding/base64"
	//"fmt"
	c "pkg/cache"
	"time"
)

type Token struct {
    Value      string
    Expiration time.Time
    Used       bool
}

func generateToken() string {
    b := make([]byte, 32)
    rand.Read(b)
    return base64.URLEncoding.EncodeToString(b)
}

func CreateToken(cache c.Cache) (Token) {
    oneTimeToken := generateToken()
    token := Token{
        Value:      oneTimeToken,
        Expiration: time.Now().Add(5 * time.Minute), // token will expire in 5 minutes
        Used:       false,
    }
    cache.Set(oneTimeToken, token, 5 * time.Minute)
    return token
}

func ValidateToken(tokenValue string, cache c.Cache) bool {
    // get token from database or cache
    
    if val, ok := cache.Get(tokenValue); !ok {
        var res Token
        res, ok = val.(Token)
        if !ok {
            return false
        }
        if time.Now().After(res.Expiration) || res.Used {
            return false
        }
        return false
    }
    return true
}

func UseToken(t *Token, cache c.Cache) (Token, *c.Cache, bool) {
    if _, ok := cache.Get(t.Value); !ok {
        return Token{}, nil, false
    }
    t.Used = true
    ok := cache.Update(t.Value, t, 1 * time.Minute)
    if !ok {
        return Token{}, nil, false
    }
    return *t, &cache, true 
}
