package model

import (
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/utils"
)

func NewTenant(id defs.ID, value TenantValue) Tenant {
	return &tenant{
		id:    id,
		value: value,
	}
}

type Tenant interface {
	ID() defs.ID
	// FullValue 核心数据
	FullValue() TenantValue
	// GenerateSecretKey 生成密钥
	GenerateSecretKey() string
	// VerifySecretKey 验证密钥
	VerifySecretKey(sk string) bool
}

type tenant struct {
	id    defs.ID
	value TenantValue
}

func (t tenant) ID() defs.ID {
	return t.id
}

func (t tenant) FullValue() TenantValue {
	return t.value
}

func (t *tenant) GenerateSecretKey() string {
	if len(t.value.SecretKey) == 0 {
		t.value.SecretKey = utils.UUID()
	}
	return t.value.SecretKey
}

func (t tenant) VerifySecretKey(sk string) bool {
	if len(t.value.SecretKey) == 0 {
		return false
	}
	return t.value.SecretKey == sk
}

type TenantValue struct {
	// TenantID 租户ID
	TenantID string
	// TenantName 租户名
	TenantName string
	// SecretKey 租户密钥
	SecretKey string
	// RedirectUrl 授权码重定向路径
	RedirectUrl string
	// AccessTokenTimeLimit 访问Token有效时限
	AccessTokenTimeLimit uint64
}
