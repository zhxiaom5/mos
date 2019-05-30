package apimanager

import (
	"github.com/jinzhu/gorm"
)

// APITab API管理
type APITab struct {
	APIURI  string `gorm:"type:varchar(256);not null;column:api_uri"`
	Content string `gorm:"type:varchar(256);not null;column:content"`
	APIKey  []APIKey
	gorm.Model
}

// TableName APITab 表名
func (APITab) TableName() string {
	return "APITab"
}

// APITabInfo API管理
type APITabInfo struct {
	ID         uint         `json:"id"`
	Content    string       `json:"content"`
	APIURI     string       `json:"api_uri"`
	APIKeyInfo []APIKeyInfo `json:"api_key_info"`
}

// APIKey API管理
type APIKey struct {
	AuthKey   string `gorm:"type:varchar(256);not null;column:auth_key"`
	IsDisable bool   `gorm:"not null;column:is_disable"`
	User      string `gorm:"type:varchar(256);not null;column:user"`
	APITabID  uint
	gorm.Model
}

// TableName APIKey 表名
func (APIKey) TableName() string {
	return "APIKey"
}

// APIKeyInfo API管理
type APIKeyInfo struct {
	ID        uint   `json:"id"`
	AuthKey   string `json:"auth_key"`
	IsDisable bool   `json:"is_disable"`
	User      string `json:"user"`
}
