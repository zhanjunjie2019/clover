package uctx

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zhanjunjie2019/clover/global/confs"
	"github.com/zhanjunjie2019/clover/global/consts"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/errs"
	"github.com/zhanjunjie2019/clover/global/utils"
	"time"
)

func GetJwtClaims(c *gin.Context) (defs.JwtClaims, error) {
	jc, ok := c.Get(consts.CtxJwtVar)
	if ok {
		claims, ok := jc.(defs.JwtClaims)
		if ok {
			return claims, nil
		} else {
			return defs.JwtClaims{}, errs.TokenMalformedErr
		}
	} else {
		return GetJwtClaimsByBearerToken(c)
	}
}

func CreateJwtClaimsToken(claims defs.JwtClaims) (string, error) {
	jwtConfig := confs.GetGlobalConfig().JwtConfig
	expiresTime := jwtConfig.ExpiresTime
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Duration(expiresTime) * time.Second)},
		ID:        utils.UUID(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := jwtConfig.SigningKey + claims.TenantID
	return token.SignedString([]byte(key))
}

func GetJwtClaimsByBearerToken(c *gin.Context) (defs.JwtClaims, error) {
	token := c.Request.Header.Get(consts.TokenHeaderKey)
	if len(token) > 0 {
		jc, err := ParseYhUserTokenClaims(token, GetTenantID(c))
		if err == nil && jc != nil {
			c.Set(consts.CtxJwtVar, *jc)
			return *jc, nil
		} else {
			return defs.JwtClaims{}, err
		}
	}
	return defs.JwtClaims{}, errs.TokenNoExistErr
}

func ParseYhUserTokenClaims(tokenString, tenantID string) (*defs.JwtClaims, error) {
	key := confs.GetGlobalConfig().JwtConfig.SigningKey + tenantID
	token, err := jwt.ParseWithClaims(tokenString, &defs.JwtClaims{}, func(*jwt.Token) (any, error) {
		return []byte(key), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errs.TokenMalformedErr
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errs.TokenExpiredErr
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errs.TokenNotValidYetErr
			} else {
				return nil, errs.TokenInvalidErr
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*defs.JwtClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errs.TokenInvalidErr
	} else {
		return nil, errs.TokenInvalidErr
	}
}
