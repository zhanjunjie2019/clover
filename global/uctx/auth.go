package uctx

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zhanjunjie2019/clover/global/confs"
	"github.com/zhanjunjie2019/clover/global/consts"
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/errs"
	"github.com/zhanjunjie2019/clover/global/utils"
	"go-micro.dev/v4/metadata"
	"time"
)

func GetJwtClaims(ctx context.Context) (defs.JwtClaims, error) {
	c, ok := ctx.(*gin.Context)
	if ok {
		jc, ok := c.Get(consts.CtxJwtVar)
		if ok {
			claims, ok := jc.(defs.JwtClaims)
			if ok {
				return claims, nil
			} else {
				return defs.JwtClaims{}, errs.TokenMalformedErr
			}
		}
		return GetJwtClaimsByBearerToken(c)
	} else {
		jc := ctx.Value(consts.CtxJwtVar)
		if jc != nil {
			claims, ok := jc.(defs.JwtClaims)
			if ok {
				return claims, nil
			}
		}
		return defs.JwtClaims{}, errs.TokenMalformedErr
	}
}

func NewJwtClaimsToken(
	tenantID string,
	userID uint64,
	username string,
	auths []string,
	accessTokenExpTime int64,
) (string, error) {
	// 构建accessToken
	accessExp := jwt.NumericDate{
		Time: time.Unix(accessTokenExpTime, 0),
	}
	accessJwtClaims := defs.JwtClaims{
		TenantID: tenantID,
		UserID:   userID,
		Username: username,
		Auths:    auths,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        utils.UUID(),
			ExpiresAt: &accessExp,
		},
	}
	return CreateJwtClaimsToken(accessJwtClaims)
}

func CreateJwtClaimsToken(claims defs.JwtClaims) (string, error) {
	jwtConfig := confs.GetGlobalConfig().JwtConfig
	expiresTime := jwtConfig.ExpiresTime
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Duration(expiresTime) * time.Second)},
		ID:        utils.UUID(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := jwtConfig.SigningKey
	return token.SignedString([]byte(key))
}

func GetJwtClaimsByBearerToken(c *gin.Context) (defs.JwtClaims, error) {
	token := c.Request.Header.Get(consts.TokenHeaderKey)
	if len(token) == 0 {
		cookie, err := c.Request.Cookie(consts.TokenHeaderKey)
		if err != nil || cookie == nil {
			return defs.JwtClaims{}, errs.TokenNoExistErr
		}
		token = cookie.Value
	}
	if len(token) > 0 {
		jc, err := ParseYhUserTokenClaims(token)
		if err == nil && jc != nil {
			c.Set(consts.CtxJwtVar, *jc)
			return *jc, nil
		} else {
			return defs.JwtClaims{}, err
		}
	}
	return defs.JwtClaims{}, errs.TokenNoExistErr
}

func GetJwtClaimsByGrpcCtx(ctx context.Context) (context.Context, defs.JwtClaims, error) {
	token, ok := metadata.Get(ctx, consts.TokenHeaderKey)
	if ok && len(token) > 0 {
		jc, err := ParseYhUserTokenClaims(token)
		if err == nil && jc != nil {
			ctx = context.WithValue(ctx, consts.CtxJwtVar, *jc)
			return ctx, *jc, nil
		} else {
			return ctx, defs.JwtClaims{}, err
		}
	}
	return ctx, defs.JwtClaims{}, errs.TokenNoExistErr
}

func ParseYhUserTokenClaims(tokenString string) (*defs.JwtClaims, error) {
	key := confs.GetGlobalConfig().JwtConfig.SigningKey
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
