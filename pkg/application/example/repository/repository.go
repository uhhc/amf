package repository

import (
	"time"
)

// Example represents an example
type Example struct {
	// 主键
	ID int32 `json:"id,omitempty" example:"1"`
	// 样例ID
	ExampleID string `json:"example_id" example:"d3e62fac-27cd-45e6-83c3-826c519f7f7b"`
	// 样例名称
	ExampleName string `json:"example_name" example:"example_name"`
	// 样例状态
	Status string `json:"status" example:"Pending"`
	// 创建时间
	CreateTime time.Time `json:"create_time,omitempty" example:"2019-11-26 11:55:50"`
}

// TableName set the name of Example
func (Example) TableName() string {
	return "info_examples"
}

// ExampleList represents a list of examples
type ExampleList struct {
	// 数据列表
	Data []Example `json:"data"`
	// 总数
	Count int32 `json:"count" example:"100"`
}

// NewExampleList return a ExampleList instance
func NewExampleList(data []Example, count int32) *ExampleList {
	return &ExampleList{
		Data:  data,
		Count: count,
	}
}

// Repository describes the persistence on example model
type Repository interface {
	CreateExample(example Example) error
	GetExampleByExampleID(exampleID string) (Example, error)
	Update(exampleID string, attrs ...interface{}) error
	ChangeStatus(exampleID string, status string) error
	ListWhere(where Example, orderBy string, limit int32, offset int32) ([]Example, int32, error)
	ListWhereComplex(orderBy string, limit int32, offset int32, query interface{}, args ...interface{}) ([]Example, int32, error)
	DeleteIn(ids string) error
}
