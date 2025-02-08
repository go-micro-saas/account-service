// Package data
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package data

import (
	"bytes"
	context "context"
	"database/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-micro-saas/account-service/app/account-service/internal/data/po"
	datarepos "github.com/go-micro-saas/account-service/app/account-service/internal/data/repo"
	schemas "github.com/go-micro-saas/account-service/app/account-service/internal/data/schema/user_event_history"
	gormpkg "github.com/ikaiguang/go-srv-kit/data/gorm"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	gorm "gorm.io/gorm"
	"strings"
)

// userEventHistoryRepo repo
type userEventHistoryRepo struct {
	log                    *log.Helper
	dbConn                 *gorm.DB                 // *gorm.DB
	UserEventHistorySchema schemas.UserEventHistory // UserEventHistory
}

// NewUserEventHistoryRepo new data repo
func NewUserEventHistoryRepo(logger log.Logger, dbConn *gorm.DB) datarepos.UserEventHistoryRepo {
	logHelper := log.NewHelper(log.With(logger, "module", "account-service/data/user_event_history"))
	return &userEventHistoryRepo{
		log:    logHelper,
		dbConn: dbConn,
	}
}

func (s *userEventHistoryRepo) NewTransaction(ctx context.Context, opts ...*sql.TxOptions) gormpkg.TransactionInterface {
	return gormpkg.NewTransaction(ctx, s.dbConn, opts...)
}

// =============== 创建 ===============

// create insert one
func (s *userEventHistoryRepo) create(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserEventHistory) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.UserEventHistorySchema.TableName()).
		Create(dataModel).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		return errorpkg.Wrap(e, err)
	}
	return
}

// Create insert one
func (s *userEventHistoryRepo) Create(ctx context.Context, dataModel *po.UserEventHistory) error {
	return s.create(ctx, s.dbConn, dataModel)
}

// CreateWithDBConn create
func (s *userEventHistoryRepo) CreateWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserEventHistory) error {
	return s.create(ctx, dbConn, dataModel)
}

// existCreate exist create
func (s *userEventHistoryRepo) existCreate(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserEventHistory) (anotherModel *po.UserEventHistory, isNotFound bool, err error) {
	anotherModel = new(po.UserEventHistory)
	err = dbConn.WithContext(ctx).
		Table(s.UserEventHistorySchema.TableName()).
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
func (s *userEventHistoryRepo) ExistCreate(ctx context.Context, dataModel *po.UserEventHistory) (anotherModel *po.UserEventHistory, isNotFound bool, err error) {
	return s.existCreate(ctx, s.dbConn, dataModel)
}

// ExistCreateWithDBConn exist create
func (s *userEventHistoryRepo) ExistCreateWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserEventHistory) (anotherModel *po.UserEventHistory, isNotFound bool, err error) {
	return s.existCreate(ctx, dbConn, dataModel)
}

// createInBatches create many
func (s *userEventHistoryRepo) createInBatches(ctx context.Context, dbConn *gorm.DB, dataModels []*po.UserEventHistory, batchSize int) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.UserEventHistorySchema.TableName()).
		CreateInBatches(dataModels, batchSize).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		return errorpkg.Wrap(e, err)
	}
	return
}

// CreateInBatches create many
func (s *userEventHistoryRepo) CreateInBatches(ctx context.Context, dataModels []*po.UserEventHistory, batchSize int) error {
	return s.createInBatches(ctx, s.dbConn, dataModels, batchSize)
}

// CreateInBatchesWithDBConn create many
func (s *userEventHistoryRepo) CreateInBatchesWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModels []*po.UserEventHistory, batchSize int) error {
	return s.createInBatches(ctx, dbConn, dataModels, batchSize)
}

// =============== 更新 ===============

// update update
func (s *userEventHistoryRepo) update(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserEventHistory) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.UserEventHistorySchema.TableName()).
		// Where(schemas.FieldId+" = ?", dataModel.Id).
		Save(dataModel).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		return errorpkg.Wrap(e, err)
	}
	return
}

// Update update
func (s *userEventHistoryRepo) Update(ctx context.Context, dataModel *po.UserEventHistory) error {
	return s.update(ctx, s.dbConn, dataModel)
}

// UpdateWithDBConn update
func (s *userEventHistoryRepo) UpdateWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserEventHistory) error {
	return s.update(ctx, dbConn, dataModel)
}

// existUpdate exist update
func (s *userEventHistoryRepo) existUpdate(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserEventHistory) (anotherModel *po.UserEventHistory, isNotFound bool, err error) {
	anotherModel = new(po.UserEventHistory)
	err = dbConn.WithContext(ctx).
		Table(s.UserEventHistorySchema.TableName()).
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
func (s *userEventHistoryRepo) ExistUpdate(ctx context.Context, dataModel *po.UserEventHistory) (anotherModel *po.UserEventHistory, isNotFound bool, err error) {
	return s.existUpdate(ctx, s.dbConn, dataModel)
}

// ExistUpdateWithDBConn exist update
func (s *userEventHistoryRepo) ExistUpdateWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserEventHistory) (anotherModel *po.UserEventHistory, isNotFound bool, err error) {
	return s.existUpdate(ctx, dbConn, dataModel)
}

// =============== query one : 查一个 ===============

// queryOneById query one by id
func (s *userEventHistoryRepo) queryOneById(ctx context.Context, dbConn *gorm.DB, id interface{}) (dataModel *po.UserEventHistory, isNotFound bool, err error) {
	dataModel = new(po.UserEventHistory)
	err = dbConn.WithContext(ctx).
		Table(s.UserEventHistorySchema.TableName()).
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
func (s *userEventHistoryRepo) QueryOneById(ctx context.Context, id interface{}) (dataModel *po.UserEventHistory, isNotFound bool, err error) {
	return s.queryOneById(ctx, s.dbConn, id)
}

// QueryOneByIdWithDBConn query one by id
func (s *userEventHistoryRepo) QueryOneByIdWithDBConn(ctx context.Context, dbConn *gorm.DB, id interface{}) (dataModel *po.UserEventHistory, isNotFound bool, err error) {
	return s.queryOneById(ctx, dbConn, id)
}

// queryOneByConditions query one by conditions
func (s *userEventHistoryRepo) queryOneByConditions(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}) (dataModel *po.UserEventHistory, isNotFound bool, err error) {
	dataModel = new(po.UserEventHistory)
	dbConn = dbConn.WithContext(ctx).Table(s.UserEventHistorySchema.TableName())
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
func (s *userEventHistoryRepo) QueryOneByConditions(ctx context.Context, conditions map[string]interface{}) (dataModel *po.UserEventHistory, isNotFound bool, err error) {
	return s.queryOneByConditions(ctx, s.dbConn, conditions)
}

// QueryOneByConditionsWithDBConn query one by conditions
func (s *userEventHistoryRepo) QueryOneByConditionsWithDBConn(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}) (dataModel *po.UserEventHistory, isNotFound bool, err error) {
	return s.queryOneByConditions(ctx, dbConn, conditions)
}

// =============== query all : 查全部 ===============

// queryAllByConditions query all by conditions
func (s *userEventHistoryRepo) queryAllByConditions(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}) (dataModels []*po.UserEventHistory, err error) {
	dbConn = dbConn.WithContext(ctx).Table(s.UserEventHistorySchema.TableName())
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
func (s *userEventHistoryRepo) QueryAllByConditions(ctx context.Context, conditions map[string]interface{}) ([]*po.UserEventHistory, error) {
	return s.queryAllByConditions(ctx, s.dbConn, conditions)
}

// QueryAllByConditionsWithDBConn query all by conditions
func (s *userEventHistoryRepo) QueryAllByConditionsWithDBConn(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}) ([]*po.UserEventHistory, error) {
	return s.queryAllByConditions(ctx, dbConn, conditions)
}

// =============== list : 列表 ===============

// list 列表
func (s *userEventHistoryRepo) list(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}, paginatorArgs *gormpkg.PaginatorArgs) (dataModels []*po.UserEventHistory, recordCount int64, err error) {
	// query where
	dbConn = dbConn.WithContext(ctx).Table(s.UserEventHistorySchema.TableName())
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
func (s *userEventHistoryRepo) List(ctx context.Context, conditions map[string]interface{}, paginatorArgs *gormpkg.PaginatorArgs) ([]*po.UserEventHistory, int64, error) {
	return s.list(ctx, s.dbConn, conditions, paginatorArgs)
}

// ListWithDBConn 列表
func (s *userEventHistoryRepo) ListWithDBConn(ctx context.Context, dbConn *gorm.DB, conditions map[string]interface{}, paginatorArgs *gormpkg.PaginatorArgs) ([]*po.UserEventHistory, int64, error) {
	return s.list(ctx, dbConn, conditions, paginatorArgs)
}

// =============== delete : 删除 ===============

// delete delete one
func (s *userEventHistoryRepo) delete(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserEventHistory) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.UserEventHistorySchema.TableName()).
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
func (s *userEventHistoryRepo) Delete(ctx context.Context, dataModel *po.UserEventHistory) error {
	return s.delete(ctx, s.dbConn, dataModel)
}

// DeleteWithDBConn delete one
func (s *userEventHistoryRepo) DeleteWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModel *po.UserEventHistory) error {
	return s.delete(ctx, dbConn, dataModel)
}

// deleteByIds delete by ids
func (s *userEventHistoryRepo) deleteByIds(ctx context.Context, dbConn *gorm.DB, ids interface{}) (err error) {
	err = dbConn.WithContext(ctx).
		Table(s.UserEventHistorySchema.TableName()).
		Where(schemas.FieldId+" in (?)", ids).
		Delete(po.UserEventHistory{}).Error
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		err = errorpkg.Wrap(e, err)
		return err
	}
	return
}

// DeleteByIds delete by ids
func (s *userEventHistoryRepo) DeleteByIds(ctx context.Context, ids interface{}) error {
	return s.deleteByIds(ctx, s.dbConn, ids)
}

// DeleteByIdsWithDBConn delete by ids
func (s *userEventHistoryRepo) DeleteByIdsWithDBConn(ctx context.Context, dbConn *gorm.DB, ids interface{}) error {
	return s.deleteByIds(ctx, dbConn, ids)
}

// =============== insert : 批量入库 ===============

var _ gormpkg.BatchInsertRepo = new(UserEventHistorySlice)

// UserEventHistorySlice 表切片
type UserEventHistorySlice []*po.UserEventHistory

// TableName 表名
func (s *UserEventHistorySlice) TableName() string {
	return schemas.UserEventHistorySchema.TableName()
}

// Len 长度
func (s *UserEventHistorySlice) Len() int {
	return len(*s)
}

// InsertColumns 批量入库的列
func (s *UserEventHistorySlice) InsertColumns() (columnList []string, placeholder string) {
	// columns
	columnList = []string{
		schemas.FieldCreatedTime, schemas.FieldUpdatedTime,
		schemas.FieldEventName, schemas.FieldEventStatus,
		schemas.FieldEventContent, schemas.FieldEventError,
		schemas.FieldRetryEventTime, schemas.FieldRetryEventCounter,
		schemas.FieldRetryEventResult,
	}

	// placeholders
	insertPlaceholderBytes := bytes.Repeat([]byte("?, "), len(columnList))
	insertPlaceholderBytes = bytes.TrimSuffix(insertPlaceholderBytes, []byte(", "))

	return columnList, string(insertPlaceholderBytes)
}

// InsertValues 批量入库的值
func (s *UserEventHistorySlice) InsertValues(args *gormpkg.BatchInsertValueArgs) (prepareData []interface{}, placeholderSlice []string) {
	dataModels := (*s)[args.StepStart:args.StepEnd]
	for index := range dataModels {
		// placeholder
		placeholderSlice = append(placeholderSlice, "("+args.InsertPlaceholder+")")

		// prepare data
		prepareData = append(prepareData, dataModels[index].CreatedTime)
		prepareData = append(prepareData, dataModels[index].UpdatedTime)
		prepareData = append(prepareData, dataModels[index].EventName)
		prepareData = append(prepareData, dataModels[index].EventStatus)
		prepareData = append(prepareData, dataModels[index].EventContent)
		prepareData = append(prepareData, dataModels[index].EventError)
		prepareData = append(prepareData, dataModels[index].RetryEventTime)
		prepareData = append(prepareData, dataModels[index].RetryEventCounter)
		prepareData = append(prepareData, dataModels[index].RetryEventResult)
	}
	return prepareData, placeholderSlice
}

// UpdateColumns 批量入库的列
func (s *UserEventHistorySlice) UpdateColumns() (columnList []string) {
	// columns
	columnList = []string{
		schemas.FieldCreatedTime + "=excluded." + schemas.FieldCreatedTime,
		schemas.FieldUpdatedTime + "=excluded." + schemas.FieldUpdatedTime,
		schemas.FieldEventName + "=excluded." + schemas.FieldEventName,
		schemas.FieldEventStatus + "=excluded." + schemas.FieldEventStatus,
		schemas.FieldEventContent + "=excluded." + schemas.FieldEventContent,
		schemas.FieldEventError + "=excluded." + schemas.FieldEventError,
		schemas.FieldRetryEventTime + "=excluded." + schemas.FieldRetryEventTime,
		schemas.FieldRetryEventCounter + "=excluded." + schemas.FieldRetryEventCounter,
		schemas.FieldRetryEventResult + "=excluded." + schemas.FieldRetryEventResult,
	}
	return columnList
}

// ConflictActionForMySQL 更新冲突时的操作
func (s *UserEventHistorySlice) ConflictActionForMySQL() (req *gormpkg.BatchInsertConflictActionReq) {
	req = &gormpkg.BatchInsertConflictActionReq{
		OnConflictValueAlias:  "AS excluded",
		OnConflictTarget:      "ON DUPLICATE KEY",
		OnConflictAction:      "UPDATE " + strings.Join(s.UpdateColumns(), ", "),
		OnConflictPrepareData: nil,
	}
	return req
}

// ConflictActionForPostgres 更新冲突时的操作
func (s *UserEventHistorySlice) ConflictActionForPostgres() (req *gormpkg.BatchInsertConflictActionReq) {
	req = &gormpkg.BatchInsertConflictActionReq{
		OnConflictValueAlias:  "",
		OnConflictTarget:      "ON CONFLICT(id)",
		OnConflictAction:      "DO UPDATE SET " + strings.Join(s.UpdateColumns(), ", "),
		OnConflictPrepareData: nil,
	}
	return req
}

// insert 批量插入
func (s *userEventHistoryRepo) insert(ctx context.Context, dbConn *gorm.DB, dataModels UserEventHistorySlice) error {
	err := gormpkg.BatchInsertWithContext(ctx, dbConn, &dataModels)
	if err != nil {
		e := errorpkg.ErrorInternalServer("")
		err = errorpkg.Wrap(e, err)
		return err
	}
	return nil
}

// Insert 批量插入
func (s *userEventHistoryRepo) Insert(ctx context.Context, dataModels []*po.UserEventHistory) error {
	return s.insert(ctx, s.dbConn, dataModels)
}

// InsertWithDBConn 批量插入
func (s *userEventHistoryRepo) InsertWithDBConn(ctx context.Context, dbConn *gorm.DB, dataModels []*po.UserEventHistory) error {
	return s.insert(ctx, dbConn, dataModels)
}

// =============== conditions : 条件 ===============

// WhereConditions orm where
func (s *userEventHistoryRepo) WhereConditions(dbConn *gorm.DB, conditions map[string]interface{}) *gorm.DB {

	// table name
	// tableName := s.UserEventHistorySchema.TableName()

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

	// event_name
	// if data, ok := conditions[schemas.FieldEventName]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldEventName+" = ?", data)
	// }

	// event_status
	// if data, ok := conditions[schemas.FieldEventStatus]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldEventStatus+" = ?", data)
	// }

	// event_content
	// if data, ok := conditions[schemas.FieldEventContent]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldEventContent+" = ?", data)
	// }

	// event_error
	// if data, ok := conditions[schemas.FieldEventError]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldEventError+" = ?", data)
	// }

	// retry_event_time
	// if data, ok := conditions[schemas.FieldRetryEventTime]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldRetryEventTime+" = ?", data)
	// }

	// retry_event_counter
	// if data, ok := conditions[schemas.FieldRetryEventCounter]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldRetryEventCounter+" = ?", data)
	// }

	// retry_event_result
	// if data, ok := conditions[schemas.FieldRetryEventResult]; ok {
	// 	   dbConn = dbConn.Where(tableName+"."+schemas.FieldRetryEventResult+" = ?", data)
	// }

	return dbConn
}

// UpdateColumns update columns
func (s *userEventHistoryRepo) UpdateColumns(conditions map[string]interface{}) map[string]interface{} {

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

	// event_name
	//if data, ok := conditions[schemas.FieldEventName]; ok {
	//	updateColumns[schemas.FieldEventName] = data
	//}

	// event_status
	//if data, ok := conditions[schemas.FieldEventStatus]; ok {
	//	updateColumns[schemas.FieldEventStatus] = data
	//}

	// event_content
	//if data, ok := conditions[schemas.FieldEventContent]; ok {
	//	updateColumns[schemas.FieldEventContent] = data
	//}

	// event_error
	//if data, ok := conditions[schemas.FieldEventError]; ok {
	//	updateColumns[schemas.FieldEventError] = data
	//}

	// retry_event_time
	//if data, ok := conditions[schemas.FieldRetryEventTime]; ok {
	//	updateColumns[schemas.FieldRetryEventTime] = data
	//}

	// retry_event_counter
	//if data, ok := conditions[schemas.FieldRetryEventCounter]; ok {
	//	updateColumns[schemas.FieldRetryEventCounter] = data
	//}

	// retry_event_result
	//if data, ok := conditions[schemas.FieldRetryEventResult]; ok {
	//	updateColumns[schemas.FieldRetryEventResult] = data
	//}

	return updateColumns
}
