package models

// Tenant type
type Tenant struct {
	BaseModel
	TenantHost         string
	TenantUser         string
	TenantUserPassword string  `json:"omitempty"`
	Admins             []*User `gorm:"many2many:tenant_user_rel;"`
	Solutions          []*Solution
}
