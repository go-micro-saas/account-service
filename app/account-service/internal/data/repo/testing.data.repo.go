package datarepos

import (
	"context"
	"github.com/go-micro-saas/account-service/app/account-service/internal/data/po"
)

type TestingDataRepo interface {
	HelloWorld(ctx context.Context, dataModel *po.HelloWorld) error
	QueryByID(ctx context.Context, id uint64) (dataModel *po.HelloWorld, isNotFound bool, err error)
	Create(ctx context.Context, dataModel *po.HelloWorld) error
}
