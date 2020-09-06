package repository

import (
	"strings"

	"github.com/uhhc/sdk-common-go/db"
	"github.com/uhhc/sdk-common-go/log"
)

type repository struct {
	db     *db.DbClient
	logger log.Logger
}

// NewRepository returns a concrete repository backed by database
func NewRepository(db *db.DbClient, logger log.Logger) Repository {
	logger.SugaredLogger = logger.With("repo", "gormdb")
	return &repository{
		db:     db,
		logger: logger,
	}
}

// CreateExample inserts a new example into db
func (repo *repository) CreateExample(example Example) error {
	ret := repo.db.Create(&example)
	return ret.Error
}

// Update to update an example
func (repo *repository) Update(exampleID string, attrs ...interface{}) error {
	model := Example{
		ExampleID: exampleID,
	}
	ret := repo.db.Model(&model).Where(model).Update(attrs...)
	return ret.Error
}

// ChangeStatus changes the example status
func (repo *repository) ChangeStatus(exampleID string, status string) error {
	return repo.Update(exampleID, "status", status)
}

// GetExampleByExampleID query the example by given id
func (repo *repository) GetExampleByExampleID(exampleID string) (Example, error) {
	ex := Example{
		ExampleID: exampleID,
	}
	ret := repo.db.Where(ex).First(&ex)
	return ex, ret.Error
}

// ListWhere to get companys in some conditions
// https://gorm.io/docs/query.html#Where
func (repo *repository) ListWhere(where Example, orderBy string, limit int32, offset int32) ([]Example, int32, error) {
	var (
		data  []Example
		count int32
		model = Example{}
	)
	ret := repo.db.Model(model).Where(where).Order(orderBy).Count(&count).Limit(limit).Offset(offset).Find(&data)
	return data, count, ret.Error
}

// ListWhereComplex to get examples in some complex conditions
func (repo *repository) ListWhereComplex(orderBy string, limit int32, offset int32, query interface{}, args ...interface{}) ([]Example, int32, error) {
	var (
		data  []Example
		count int32
		model = Example{}
	)
	ret := repo.db.Model(model).Where(query, args...).Order(orderBy).Count(&count).Limit(limit).Offset(offset).Find(&data)
	return data, count, ret.Error
}

// Delete to delete a specific example
func (repo *repository) Delete(example Example) error {
	ret := repo.db.Delete(&example)
	return ret.Error
}

// DeleteIn to delete several specific examples
func (repo *repository) DeleteIn(ids string) error {
	model := Example{}
	idList := strings.Split(ids, ",")
	ret := repo.db.Delete(&model, "example_id IN (?)", idList)
	return ret.Error
}

// Close implements DB.Close
func (repo *repository) Close() error {
	return repo.db.Close()
}
