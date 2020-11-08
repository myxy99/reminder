/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/8 16:20
 */
package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const Secret = "kih**&hgyxshq##js"

type Info struct {
	Openid string `json:"open_id"`
}

func (user *Info) GenerateToken() (string, error) {
	claim := jwt.MapClaims{
		"id":  user.Openid,
		"nbf": time.Now().Unix(),
		"iat": time.Now().Unix(),
		"exp": time.Now().Unix() + 3*60*60,
		"iss": "myxy99.cn",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokens, err := token.SignedString([]byte(Secret))
	return tokens, err
}

func secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	}
}

func (user *Info) ParseToken(tokens string) (err error) {
	token, err := jwt.Parse(tokens, secret())
	if err != nil {
		return
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("cannot convert claim to mapclaim")
		return
	}
	//验证token，如果token被修改过则为false
	if !token.Valid {
		err = errors.New("token is invalid")
		return
	}
	user.Openid = claim["id"].(string)
	return err
}