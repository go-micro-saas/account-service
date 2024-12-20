// Package po
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package po

import (
	enumv1 "github.com/go-micro-saas/account-service/api/account-service/v1/enums"
	datatypes "gorm.io/datatypes"
	time "time"
)

var _ = time.Time{}

var _ = datatypes.JSON{}

// UserConfirmCode ENGINE InnoDB CHARSET utf8mb4 COMMENT '用户验证码确认表'
type UserConfirmCode struct {
	Id            uint64                                         `gorm:"column:id;primaryKey" json:"id"`              // ID
	CreatedTime   time.Time                                      `gorm:"column:created_time" json:"created_time"`     // 创建时间
	UpdatedTime   time.Time                                      `gorm:"column:updated_time" json:"updated_time"`     // 最后修改时间
	UserIdentify  string                                         `gorm:"column:user_identify" json:"user_identify"`   // 用户标识；手机、邮箱、。。。
	ConfirmType   enumv1.UserConfirmTypeEnum_UserConfirmType     `gorm:"column:confirm_type" json:"confirm_type"`     // 确认方式；1：邮箱，2：手机，3：密码，。。。
	ConfirmCode   string                                         `gorm:"column:confirm_code" json:"confirm_code"`     // 验证码
	ConfirmStatus enumv1.UserConfirmStatusEnum_UserConfirmStatus `gorm:"column:confirm_status" json:"confirm_status"` // 确认状态；0：未指定，1：确认中，2：已确认，3：已过期，2：已取消
	ConfirmTime   uint64                                         `gorm:"column:confirm_time" json:"confirm_time"`     // 确认时间
	CancelTime    uint64                                         `gorm:"column:cancel_time" json:"cancel_time"`       // 取消时间
}

// NewUserConfirmCode default UserConfirmCode
func NewUserConfirmCode(code string) *UserConfirmCode {
	var (
		now = time.Now()
	)
	dataModel := &UserConfirmCode{
		Id:            0,
		CreatedTime:   now,
		UpdatedTime:   now,
		UserIdentify:  "",
		ConfirmType:   enumv1.UserConfirmTypeEnum_UNSPECIFIED,
		ConfirmCode:   code,
		ConfirmStatus: enumv1.UserConfirmStatusEnum_UNSPECIFIED,
		ConfirmTime:   0,
		CancelTime:    0,
	}
	return dataModel
}
