// Package schemas
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package schemas

import (
	migrationuitl "github.com/ikaiguang/go-srv-kit/data/migration"
	datatypes "gorm.io/datatypes"
	gorm "gorm.io/gorm"
	time "time"
)

var _ = time.Time{}

var _ = datatypes.JSON{}

// UserRegPhoneSchema UserRegPhone
var UserRegPhoneSchema UserRegPhone

// NewUserRegPhone new schema
func NewUserRegPhone() *UserRegPhone {
	return &UserRegPhone{}
}

const (
	TableName = "as_user_reg_phone"

	FieldId          = "id"
	FieldCreatedTime = "created_time"
	FieldUpdatedTime = "updated_time"
	FieldDeletedTime = "deleted_time"
	FieldUserId      = "user_id"
	FieldUserPhone   = "user_phone"
)

// UserRegPhone ENGINE InnoDB CHARSET utf8mb4 COMMENT '用户_注册_手机'
type UserRegPhone struct {
	Id          uint64    `gorm:"column:id;primaryKey;type:uint;autoIncrement;comment:ID" json:"id"`
	CreatedTime time.Time `gorm:"column:created_time;type:time;not null;comment:创建时间" json:"created_time"`
	UpdatedTime time.Time `gorm:"column:updated_time;type:time;not null;comment:最后修改时间" json:"updated_time"`
	DeletedTime uint64    `gorm:"column:deleted_time;type:uint;not null;default:0;comment:删除时间" json:"deleted_time"`
	UserId      uint64    `gorm:"column:user_id;index;type:uint;not null;default:0;comment:uid" json:"user_id"`
	UserPhone   string    `gorm:"column:user_phone;unique;type:string;size:255;not null;default:'';comment:手机" json:"user_phone"`
}

// TableName table name
func (s *UserRegPhone) TableName() string {
	return TableName
}

// CreateTableMigrator create table migrator
func (s *UserRegPhone) CreateTableMigrator(migrator gorm.Migrator) migrationuitl.MigrationInterface {
	return migrationuitl.NewCreateTable(migrator, migrationuitl.Version, s)
}

// DropTableMigrator create table migrator
func (s *UserRegPhone) DropTableMigrator(migrator gorm.Migrator) migrationuitl.MigrationInterface {
	return migrationuitl.NewDropTable(migrator, migrationuitl.Version, s)
}

// TableSQL table SQL
func (s *UserRegPhone) TableSQL() string {
	return `
CREATE TABLE as_user_reg_phone (
	id BIGINT unsigned auto_increment comment 'ID',
	created_time DATETIME(3) NOT NULL comment '创建时间',
	updated_time DATETIME(3) NOT NULL comment '最后修改时间',
	deleted_time BIGINT NULL DEFAULT 0 comment '删除时间',
	user_id BIGINT NOT NULL DEFAULT 0 comment 'uid',
	user_phone varchar(255) NOT NULL DEFAULT '' comment '手机',
	PRIMARY KEY (id),
	UNIQUE KEY (user_phone),
	key (user_id)
) ENGINE InnoDB,
  CHARSET utf8mb4,
  COMMENT '用户_注册_手机'
`
}
