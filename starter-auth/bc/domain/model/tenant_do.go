package model

import (
	"github.com/zhanjunjie2019/clover/global/defs"
	"github.com/zhanjunjie2019/clover/global/utils"
)

func NewTenant(id defs.ID, value TenantValue) Tenant {
	if len(value.TenantID) == 0 {
		value.TenantID = utils.UUID()
	}
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
	// SetPermissions 设置许可
	SetPermissions(permissions []string)
	// GetPermissions 获得许可
	GetPermissions() []string
}

type tenant struct {
	id          defs.ID
	value       TenantValue
	permissions []string
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

func (t *tenant) SetPermissions(permissions []string) {
	t.permissions = permissions
}

func (t *tenant) GetPermissions() []string {
	return t.permissions
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
