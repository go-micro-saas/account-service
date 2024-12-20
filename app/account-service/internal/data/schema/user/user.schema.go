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

// UserSchema User
var UserSchema User

// NewUser new schema
func NewUser() *User {
	return &User{}
}

const (
	TableName = "as_user"

	FieldId            = "id"
	FieldCreatedTime   = "created_time"
	FieldUpdatedTime   = "updated_time"
	FieldDeletedTime   = "deleted_time"
	FieldUserId        = "user_id"
	FieldUserPhone     = "user_phone"
	FieldUserEmail     = "user_email"
	FieldUserNickname  = "user_nickname"
	FieldUserAvatar    = "user_avatar"
	FieldUserGender    = "user_gender"
	FieldRegisterType  = "register_type"
	FieldUserStatus    = "user_status"
	FieldDisableTime   = "disable_time"
	FieldBlacklistTime = "blacklist_time"
	FieldPasswordHash  = "password_hash"
)

// User ENGINE InnoDB CHARSET utf8mb4 COMMENT '用户表'
type User struct {
	Id            uint64    `gorm:"column:id;primaryKey;type:uint;autoIncrement;comment:ID" json:"id"`
	CreatedTime   time.Time `gorm:"column:created_time;type:time;not null;comment:创建时间" json:"created_time"`
	UpdatedTime   time.Time `gorm:"column:updated_time;type:time;not null;comment:最后修改时间" json:"updated_time"`
	DeletedTime   uint64    `gorm:"column:deleted_time;index;type:uint;not null;default:0;comment:删除时间" json:"deleted_time"`
	UserId        uint64    `gorm:"column:user_id;unique;type:uint;not null;default:0;comment:uid" json:"user_id"`
	UserPhone     string    `gorm:"column:user_phone;type:string;size:255;not null;default:'';comment:手机" json:"user_phone"`
	UserEmail     string    `gorm:"column:user_email;type:string;size:255;not null;default:'';comment:邮箱" json:"user_email"`
	UserNickname  string    `gorm:"column:user_nickname;type:string;size:255;not null;default:'';comment:昵称" json:"user_nickname"`
	UserAvatar    string    `gorm:"column:user_avatar;type:string;size:255;not null;default:'';comment:头像" json:"user_avatar"`
	UserGender    uint32    `gorm:"column:user_gender;type:uint;not null;default:0;comment:性别；0：INIT，1：MALE，2：FEMALE，3：SECRET" json:"user_gender"`
	RegisterType  uint32    `gorm:"column:register_type;type:uint;not null;default:0;comment:注册类型；0：INIT，1：EMAIL，2：MOBILE，3：。。。参考ENUM定义" json:"register_type"`
	UserStatus    uint32    `gorm:"column:user_status;type:uint;not null;default:0;comment:状态；0：INIT，1：ENABLE，2：DISABLE，3：WHITELIST，4：BLACKLIST，5：DELETED" json:"user_status"`
	DisableTime   uint64    `gorm:"column:disable_time;type:uint;not null;default:0;comment:禁用时间" json:"disable_time"`
	BlacklistTime uint64    `gorm:"column:blacklist_time;type:uint;not null;default:0;comment:黑名单时间" json:"blacklist_time"`
	PasswordHash  string    `gorm:"column:password_hash;type:string;size:255;not null;default:'';comment:密码HASH" json:"password_hash"`
}

// TableName table name
func (s *User) TableName() string {
	return TableName
}

// CreateTableMigrator create table migrator
func (s *User) CreateTableMigrator(migrator gorm.Migrator) migrationuitl.MigrationInterface {
	return migrationuitl.NewCreateTable(migrator, migrationuitl.Version, s)
}

// DropTableMigrator create table migrator
func (s *User) DropTableMigrator(migrator gorm.Migrator) migrationuitl.MigrationInterface {
	return migrationuitl.NewDropTable(migrator, migrationuitl.Version, s)
}

// TableSQL table SQL
func (s *User) TableSQL() string {
	return `
CREATE TABLE as_user (
	id bigint unsigned auto_increment comment 'ID',
	created_time datetime(3) NOT NULL comment '创建时间',
	updated_time datetime(3) NOT NULL comment '最后修改时间',
	deleted_time bigint unsigned NOT NULL DEFAULT 0 comment '删除时间',
	user_id bigint unsigned NOT NULL DEFAULT 0 comment 'uid',
	user_phone varchar(255) NOT NULL DEFAULT '' comment '手机',
	user_email varchar(255) NOT NULL DEFAULT '' comment '邮箱',
	user_nickname varchar(255) NOT NULL DEFAULT '' comment '昵称',
	user_avatar varchar(255) NOT NULL DEFAULT '' comment '头像',
	user_gender integer unsigned NOT NULL DEFAULT 0 comment '性别；0：INIT，1：MALE，2：FEMALE，3：SECRET',
	register_type integer unsigned NOT NULL DEFAULT 0 comment '注册类型；0：INIT，1：EMAIL，2：MOBILE，3：。。。参考ENUM定义',
	user_status integer unsigned NOT NULL DEFAULT 0 comment '状态；0：INIT，1：ENABLE，2：DISABLE，3：WHITELIST，4：BLACKLIST，5：DELETED',
	disable_time bigint unsigned NOT NULL DEFAULT 0 comment '禁用时间',
	blacklist_time bigint unsigned NOT NULL DEFAULT 0 comment '黑名单时间',
	password_hash varchar(255) NOT NULL DEFAULT '' comment '密码HASH',
	PRIMARY KEY (id),
	UNIQUE KEY (user_id),
	key (deleted_time)
) ENGINE InnoDB,
  CHARSET utf8mb4,
  COMMENT '用户表'
`
}
