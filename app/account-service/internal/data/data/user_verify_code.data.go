// Package data
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package data

import (
	"bytes"
	context "context"
	"database/sql"
	"github.com/go-micro-saas/account-service/app/account-service/internal/data/po"
	datarepos "github.com/go-micro-saas/account-service/app/account-service/internal/data/repo"
	schemas "github.com/go-micro-saas/account-service/app/account-service/internal/data/schema/user_verify_code"
	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	gorm "gorm.io/gorm"
	"strings"
	"time"
)

// userVerifyCodeRepo repo
type userVerifyCodeRepo struct {
	dbConn               *gorm.DB               // *gorm.DB
	UserVerifyCodeSchema schemas.UserVerifyCode // UserVerifyCode
}

// NewUserVerifyCodeRepo new data repo
func NewUserVerifyCodeRepo(dbConn *gorm.DB) datarepos.UserVerifyCodeDataRepo {
	return &userVerifyCodeRepo{
		dbConn: dbConn,
	}
}

func (s *userVerifyCodeRepo) NewTransaction(ctx context.Context, opts ...*sql.TxOptions) gormpkg.TransactionInterface {
	return gormpkg.NewTransaction(ctx, s.dbConn, opts...)
}

// =============== 创建 ===============

// create insert one
func (s *userVerifyCodeRepo) create(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserVerifyCode) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.UserVerifyCodeSchema.TableName()).
		Create(dataModel).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		return errorpkg.Wrap(e, err)
	}
	return
}

// Create insert one
func (s *userVerifyCodeRepo) Create(ctx context.Context, dataModel *po.UserVerifyCode) error {
	return s.create(ctx, s.dbConn, dataModel)
}

// CreateWithDBConn create
func (s *userVerifyCodeRepo) CreateWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserVerifyCode) error {
	return s.create(ctx, dbConn, dataModel)
}

// existCreate exist create
func (s *userVerifyCodeRepo) existCreate(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserVerifyCode) (anotherModel *po.UserVerifyCode, isNotFound bool, err error) {
	anotherModel = new(po.UserVerifyCode)
	err = dbConn.WithContext(ctx).
		Table(s.UserVerifyCodeSchema.TableName()).
		Where(schemas.FieldId+" = ?", dataModel.Id).
		First(anotherModel).Error
	if err != nil {
		if gormpkg.IsErrRecordNotFound(err) {
			isNotFound = true
			err = nil
		} else {
			e := errorpkg.ErrorInternalServer("")
			err = errorpkg.Wrap(e, err)
		}
		return
	}
	return
}

// ExistCreate exist create
func (s *userVerifyCodeRepo) ExistCreate(ctx context.Context, dataModel *po.UserVerifyCode) (anotherModel *po.UserVerifyCode, isNotFound bool, err error) {
	return s.existCreate(ctx, s.dbConn, dataModel)
}

// ExistCreateWithDBConn exist create
func (s *userVerifyCodeRepo) ExistCreateWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserVerifyCode) (anotherModel *po.UserVerifyCode, isNotFound bool, err error) {
	return s.existCreate(ctx, dbConn, dataModel)
}

// createInBatches create many
func (s *userVerifyCodeRepo) createInBatches(ctx context.Context, dbConn *gorm.DB, dataModels []*po.UserVerifyCode, batchSize int) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.UserVerifyCodeSchema.TableName()).
		CreateInBatches(dataModels, batchSize).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		return errorpkg.Wrap(e, err)
	}
	return
}

// CreateInBatches create many
func (s *userVerifyCodeRepo) CreateInBatches(ctx context.Context, dataModels []*po.UserVerifyCode, batchSize int) error {
	return s.createInBatches(ctx, s.dbConn, dataModels, batchSize)
}

// CreateInBatchesWithDBConn create many
func (s *userVerifyCodeRepo) CreateInBatchesWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModels []*po.UserVerifyCode, batchSize int) error {
	return s.createInBatches(ctx, dbConn, dataModels, batchSize)
}

// =============== 更新 ===============

// update update
func (s *userVerifyCodeRepo) update(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserVerifyCode) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.UserVerifyCodeSchema.TableName()).
		// Where(schemas.FieldId+" = ?", dataModel.Id).
		Save(dataModel).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		return errorpkg.Wrap(e, err)
	}
	return
}

// Update update
func (s *userVerifyCodeRepo) Update(ctx context.Context, dataModel *po.UserVerifyCode) error {
	return s.update(ctx, s.dbConn, dataModel)
}

// UpdateWithDBConn update
func (s *userVerifyCodeRepo) UpdateWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserVerifyCode) error {
	return s.update(ctx, dbConn, dataModel)
}

func (s *userVerifyCodeRepo) UpdateVerifyStatus(ctx context.Context, dataModel *po.UserVerifyCode) (err error) {
	updates := map[string]interface{}{
		schemas.FieldUpdatedTime:  time.Now(),
		schemas.FieldVerifyStatus: dataModel.VerifyStatus,
		schemas.FieldConfirmTime:  dataModel.ConfirmTime,
	}
	err = s.dbConn.WithContext(ctx).
		Table(s.UserVerifyCodeSchema.TableName()).
		Where(schemas.FieldId+" = ?", dataModel.Id).
		UpdateColumns(updates).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		return errorpkg.Wrap(e, err)
	}
	return
}

// existUpdate exist update
func (s *userVerifyCodeRepo) existUpdate(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserVerifyCode) (anotherModel *po.UserVerifyCode, isNotFound bool, err error) {
	anotherModel = new(po.UserVerifyCode)
	err = dbConn.WithContext(ctx).
		Table(s.UserVerifyCodeSchema.TableName()).
		Where(schemas.FieldId+" = ?", dataModel.Id).
		Where(schemas.FieldId+" != ?", dataModel.Id).
		First(anotherModel).Error
	if err != nil {
		if gormpkg.IsErrRecordNotFound(err) {
			isNotFound = true
			err = nil
		} else {
			e := errorpkg.ErrorInternalServer("")
			err = errorpkg.Wrap(e, err)
		}
		return
	}
	return
}

// ExistUpdate exist update
func (s *userVerifyCodeRepo) ExistUpdate(ctx context.Context, dataModel *po.UserVerifyCode) (anotherModel *po.UserVerifyCode, isNotFound bool, err error) {
	return s.existUpdate(ctx, s.dbConn, dataModel)
}

// ExistUpdateWithDBConn exist update
func (s *userVerifyCodeRepo) ExistUpdateWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserVerifyCode) (anotherModel *po.UserVerifyCode, isNotFound bool, err error) {
	return s.existUpdate(ctx, dbConn, dataModel)
}

// =============== query one : 查一个 ===============

// queryOneById query one by id
func (s *userVerifyCodeRepo) queryOneById(ctx context.Context, dbConn *gorm.DB, id interface{}) (dataModel *po.UserVerifyCode, isNotFound bool, err error) {
	dataModel = new(po.UserVerifyCode)
	err = dbConn.WithContext(ctx).
		Table(s.UserVerifyCodeSchema.TableName()).
		Where(schemas.FieldId+" = ?", id).
		First(dataModel).Error
	if err != nil {
		if gormpkg.IsErrRecordNotFound(err) {
			err = nil
			isNotFound = true
		} else {
			e := errorpkg.ErrorInternalServer("")
			err = errorpkg.Wrap(e, err)
		}
		return
	}
	return
}

// QueryOneById query one by id
func (s *userVerifyCodeRepo) QueryOneById(ctx context.Context, id interface{}) (dataModel *po.UserVerifyCode, isNotFound bool, err error) {
	return s.queryOneById(ctx, s.dbConn, id)
}

// QueryOneByIdWithDBConn query one by id
func (s *userVerifyCodeRepo) QueryOneByIdWithDBConn(ctx context.Context, dbConn *gorm.DB, id interface{}) (dataModel *po.UserVerifyCode, isNotFound bool, err error) {
	return s.queryOneById(ctx, dbConn, id)
}

// queryOneByConditions query one by conditions
func (s *userVerifyCodeRepo) queryOneByConditions(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}) (dataModel *po.UserVerifyCode, isNotFound bool, err error) {
	dataModel = new(po.UserVerifyCode)
	dbConn = dbConn.WithContext(ctx).Table(s.UserVerifyCodeSchema.TableName())
	err = s.WhereConditions(dbConn, conditions).
		First(dataModel).Error
	if err != nil {
		if gormpkg.IsErrRecordNotFound(err) {
			err = nil
			isNotFound = true
		} else {
			e := errorpkg.ErrorInternalServer("")
			err = errorpkg.Wrap(e, err)
		}
		return
	}
	return
}

// QueryOneByConditions query one by conditions
func (s *userVerifyCodeRepo) QueryOneByConditions(ctx context.Context, conditions map[string]interface{}) (dataModel *po.UserVerifyCode, isNotFound bool, err error) {
	return s.queryOneByConditions(ctx, s.dbConn, conditions)
}

// QueryOneByConditionsWithDBConn query one by conditions
func (s *userVerifyCodeRepo) QueryOneByConditionsWithDBConn(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}) (dataModel *po.UserVerifyCode, isNotFound bool, err error) {
	return s.queryOneByConditions(ctx, dbConn, conditions)
}

// QueryOneVerifyCode query one by conditions
func (s *userVerifyCodeRepo) QueryOneVerifyCode(ctx context.Context, param *po.GetVerifyCodeParam) (dataModel *po.UserVerifyCode, isNotFound bool, err error) {
	dataModel = new(po.UserVerifyCode)
	dbConn := s.dbConn.WithContext(ctx).Table(s.UserVerifyCodeSchema.TableName())
	err = param.WhereConditions(dbConn).First(dataModel).Error
	if err != nil {
		if gormpkg.IsErrRecordNotFound(err) {
			err = nil
			isNotFound = true
		} else {
			e := errorpkg.ErrorInternalServer("")
			err = errorpkg.Wrap(e, err)
		}
		return
	}
	return
}

// =============== query all : 查全部 ===============

// queryAllByConditions query all by conditions
func (s *userVerifyCodeRepo) queryAllByConditions(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}) (dataModels []*po.UserVerifyCode, err error) {
	dbConn = dbConn.WithContext(ctx).Table(s.UserVerifyCodeSchema.TableName())
	err = s.WhereConditions(dbConn, conditions).
		Find(&dataModels).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		err = errorpkg.Wrap(e, err)
		return dataModels, err
	}
	return
}

// QueryAllByConditions query all by conditions
func (s *userVerifyCodeRepo) QueryAllByConditions(ctx context.Context, conditions map[string]interface{}) ([]*po.UserVerifyCode, error) {
	return s.queryAllByConditions(ctx, s.dbConn, conditions)
}

// QueryAllByConditionsWithDBConn query all by conditions
func (s *userVerifyCodeRepo) QueryAllByConditionsWithDBConn(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}) ([]*po.UserVerifyCode, error) {
	return s.queryAllByConditions(ctx, dbConn, conditions)
}

// =============== list : 列表 ===============

// list 列表
func (s *userVerifyCodeRepo) list(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}, paginatorArgs *gormpkg.PaginatorArgs) (dataModels []*po.UserVerifyCode, recordCount int64, err error) {
	// query where
	dbConn = dbConn.WithContext(ctx).Table(s.UserVerifyCodeSchema.TableName())
	dbConn = s.WhereConditions(dbConn, conditions)
	dbConn = gormpkg.AssembleWheres(dbConn, paginatorArgs.PageWheres)

	err = dbConn.Count(&recordCount).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		err = errorpkg.Wrap(e, err)
		return
	} else if recordCount == 0 {
		return // empty
	}

	// pagination
	dbConn = gormpkg.AssembleOrders(dbConn, paginatorArgs.PageOrders)
	err = gormpkg.Paginator(dbConn, paginatorArgs.PageOption).
		Find(&dataModels).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		err = errorpkg.Wrap(e, err)
		return
	}
	return
}

// List 列表
func (s *userVerifyCodeRepo) List(ctx context.Context, conditions map[string]interface{}, paginatorArgs *gormpkg.PaginatorArgs) ([]*po.UserVerifyCode, int64, error) {
	return s.list(ctx, s.dbConn, conditions, paginatorArgs)
}

// ListWithDBConn 列表
func (s *userVerifyCodeRepo) ListWithDBConn(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}, paginatorArgs *gormpkg.PaginatorArgs) ([]*po.UserVerifyCode, int64, error) {
	return s.list(ctx, dbConn, conditions, paginatorArgs)
}

// =============== delete : 删除 ===============

// delete delete one
func (s *userVerifyCodeRepo) delete(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserVerifyCode) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.UserVerifyCodeSchema.TableName()).
		Where(schemas.FieldId+" = ?", dataModel.Id).
		Delete(dataModel).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		err = errorpkg.Wrap(e, err)
		return err
	}
	return
}

// Delete delete one
func (s *userVerifyCodeRepo) Delete(ctx context.Context, dataModel *po.UserVerifyCode) error {
	return s.delete(ctx, s.dbConn, dataModel)
}

// DeleteWithDBConn delete one
func (s *userVerifyCodeRepo) DeleteWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserVerifyCode) error {
	return s.delete(ctx, dbConn, dataModel)
}

// deleteByIds delete by ids
func (s *userVerifyCodeRepo) deleteByIds(ctx context.Context, dbConn *gorm.DB, ids interface{}) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.UserVerifyCodeSchema.TableName()).
		Where(schemas.FieldId+" in (?)", ids).
		Delete(po.UserVerifyCode{}).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		err = errorpkg.Wrap(e, err)
		return err
	}
	return
}

// DeleteByIds delete by ids
func (s *userVerifyCodeRepo) DeleteByIds(ctx context.Context, ids interface{}) error {
	return s.deleteByIds(ctx, s.dbConn, ids)
}

// DeleteByIdsWithDBConn delete by ids
func (s *userVerifyCodeRepo) DeleteByIdsWithDBConn(ctx context.Context, dbConn *gorm.DB, ids interface{}) error {
	return s.deleteByIds(ctx, dbConn, ids)
}

// =============== insert : 批量入库 ===============

var _ gormpkg.BatchInsertRepo = new(UserVerifyCodeSlice)

// UserVerifyCodeSlice 表切片
type UserVerifyCodeSlice []*po.UserVerifyCode

// TableName 表名
func (s *UserVerifyCodeSlice) TableName() string {
	return schemas.UserVerifyCodeSchema.TableName()
}

// Len 长度
func (s *UserVerifyCodeSlice) Len() int {
	return len(*s)
}

// InsertColumns 批量入库的列
func (s *UserVerifyCodeSlice) InsertColumns() (columnList []string, placeholder string) {
	// columns
	columnList = []string{
		schemas.FieldCreatedTime, schemas.FieldUpdatedTime,
		schemas.FieldVerifyAccount, schemas.FieldVerifyType,
		schemas.FieldVerifyCode, schemas.FieldVerifyStatus,
		schemas.FieldConfirmTime, schemas.FieldCancelTime,
	}

	// placeholders
	insertPlaceholderBytes := bytes.Repeat([]byte("?, "), len(columnList))
	insertPlaceholderBytes = bytes.TrimSuffix(insertPlaceholderBytes, []byte(", "))

	return columnList, string(insertPlaceholderBytes)
}

// InsertValues 批量入库的值
func (s *UserVerifyCodeSlice) InsertValues(args *gormpkg.BatchInsertValueArgs) (prepareData []interface{}, placeholderSlice []string) {
	dataModels := (*s)[args.StepStart:args.StepEnd]
	for index := range dataModels {
		// placeholder
		placeholderSlice = append(placeholderSlice, "("+args.InsertPlaceholder+")")

		// prepare data
		prepareData = append(prepareData, dataModels[index].CreatedTime)
		prepareData = append(prepareData, dataModels[index].UpdatedTime)
		prepareData = append(prepareData, dataModels[index].VerifyAccount)
		prepareData = append(prepareData, dataModels[index].VerifyType)
		prepareData = append(prepareData, dataModels[index].VerifyCode)
		prepareData = append(prepareData, dataModels[index].VerifyStatus)
		prepareData = append(prepareData, dataModels[index].ConfirmTime)
		prepareData = append(prepareData, dataModels[index].CancelTime)
	}
	return prepareData, placeholderSlice
}

// UpdateColumns 批量入库的列
func (s *UserVerifyCodeSlice) UpdateColumns() (columnList []string) {
	// columns
	columnList = []string{
		schemas.FieldCreatedTime + "=excluded." + schemas.FieldCreatedTime,
		schemas.FieldUpdatedTime + "=excluded." + schemas.FieldUpdatedTime,
		schemas.FieldVerifyAccount + "=excluded." + schemas.FieldVerifyAccount,
		schemas.FieldVerifyType + "=excluded." + schemas.FieldVerifyType,
		schemas.FieldVerifyCode + "=excluded." + schemas.FieldVerifyCode,
		schemas.FieldVerifyStatus + "=excluded." + schemas.FieldVerifyStatus,
		schemas.FieldConfirmTime + "=excluded." + schemas.FieldConfirmTime,
		schemas.FieldCancelTime + "=excluded." + schemas.FieldCancelTime,
	}
	return columnList
}

// ConflictActionForMySQL 更新冲突时的操作
func (s *UserVerifyCodeSlice) ConflictActionForMySQL() (req *gormpkg.BatchInsertConflictActionReq) {
	req = &gormpkg.BatchInsertConflictActionReq{
		OnConflictValueAlias:  "AS excluded",
		OnConflictTarget:      "ON DUPLICATE KEY",
		OnConflictAction:      "UPDATE " + strings.Join(s.UpdateColumns(), ", "),
		OnConflictPrepareData: nil,
	}
	return req
}

// ConflictActionForPostgres 更新冲突时的操作
func (s *UserVerifyCodeSlice) ConflictActionForPostgres() (req *gormpkg.BatchInsertConflictActionReq) {
	req = &gormpkg.BatchInsertConflictActionReq{
		OnConflictValueAlias:  "",
		OnConflictTarget:      "ON CONFLICT(id)",
		OnConflictAction:      "DO UPDATE SET " + strings.Join(s.UpdateColumns(), ", "),
		OnConflictPrepareData: nil,
	}
	return req
}

// insert 批量插入
func (s *userVerifyCodeRepo) insert(ctx context.Context, dbConn *gorm.DB, dataModels UserVerifyCodeSlice) error {
	err := gormpkg.BatchInsertWithContext(ctx, dbConn, &dataModels)
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		err = errorpkg.Wrap(e, err)
		return err
	}
	return nil
}

// Insert 批量插入
func (s *userVerifyCodeRepo) Insert(ctx context.Context, dataModels []*po.UserVerifyCode) error {
	return s.insert(ctx, s.dbConn, dataModels)
}

// InsertWithDBConn 批量插入
func (s *userVerifyCodeRepo) InsertWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModels []*po.UserVerifyCode) error {
	return s.insert(ctx, dbConn, dataModels)
}

// =============== conditions : 条件 ===============

// WhereConditions orm where
func (s *userVerifyCodeRepo) WhereConditions(dbConn *gorm.DB, conditions map[string]interface{}) *gorm.DB {

	// table name
	// tableName := s.UserVerifyCodeSchema.TableName()

	// On-demand loading

	// id
	// if data, ok := conditions[schemas.FieldId]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldId+" = ?", data)
	// }

	// created_time
	// if data, ok := conditions[schemas.FieldCreatedTime]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldCreatedTime+" = ?", data)
	// }

	// updated_time
	// if data, ok := conditions[schemas.FieldUpdatedTime]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldUpdatedTime+" = ?", data)
	// }

	// verify_account
	// if data, ok := conditions[schemas.FieldVerifyAccount]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldVerifyAccount+" = ?", data)
	// }

	// verify_type
	// if data, ok := conditions[schemas.FieldVerifyType]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldVerifyType+" = ?", data)
	// }

	// verify_code
	// if data, ok := conditions[schemas.FieldVerifyCode]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldVerifyCode+" = ?", data)
	// }

	// verify_status
	// if data, ok := conditions[schemas.FieldVerifyStatus]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldVerifyStatus+" = ?", data)
	// }

	// confirm_time
	// if data, ok := conditions[schemas.FieldConfirmTime]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldConfirmTime+" = ?", data)
	// }

	// cancel_time
	// if data, ok := conditions[schemas.FieldCancelTime]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldCancelTime+" = ?", data)
	// }

	return dbConn
}

// UpdateColumns update columns
func (s *userVerifyCodeRepo) UpdateColumns(conditions map[string]interface{}) map[string]interface{} {

	// update columns
	updateColumns := make(map[string]interface{})

	// On-demand loading

	// id
	//if data, ok := conditions[schemas.FieldId]; ok {
	//	updateColumns[schemas.FieldId] = data
	//}

	// created_time
	//if data, ok := conditions[schemas.FieldCreatedTime]; ok {
	//	updateColumns[schemas.FieldCreatedTime] = data
	//}

	// updated_time
	//if data, ok := conditions[schemas.FieldUpdatedTime]; ok {
	//	updateColumns[schemas.FieldUpdatedTime] = data
	//}

	// verify_account
	//if data, ok := conditions[schemas.FieldVerifyAccount]; ok {
	//	updateColumns[schemas.FieldVerifyAccount] = data
	//}

	// verify_type
	//if data, ok := conditions[schemas.FieldVerifyType]; ok {
	//	updateColumns[schemas.FieldVerifyType] = data
	//}

	// verify_code
	//if data, ok := conditions[schemas.FieldVerifyCode]; ok {
	//	updateColumns[schemas.FieldVerifyCode] = data
	//}

	// verify_status
	//if data, ok := conditions[schemas.FieldVerifyStatus]; ok {
	//	updateColumns[schemas.FieldVerifyStatus] = data
	//}

	// confirm_time
	//if data, ok := conditions[schemas.FieldConfirmTime]; ok {
	//	updateColumns[schemas.FieldConfirmTime] = data
	//}

	// cancel_time
	//if data, ok := conditions[schemas.FieldCancelTime]; ok {
	//	updateColumns[schemas.FieldCancelTime] = data
	//}

	return updateColumns
}
