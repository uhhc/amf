package service

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/uhhc/sdk-common-go/log"
	"github.com/uhhc/sdk-common-go/util"
	timeUtil "github.com/uhhc/sdk-common-go/util/time"

	"github.com/uhhc/amf/pkg/application/example/repository"
	"github.com/uhhc/amf/pkg/config/errorcode"
	"github.com/uhhc/amf/pkg/grpc/client"
	"github.com/uhhc/amf/pkg/response"
)

// Service describes the Example service.
type Service interface {
	Create(example repository.Example) *response.Error
	Update(example repository.Example) *response.Error
	List(where repository.Example, order string, limit int32, page int32) (*repository.ExampleList, *response.Error)
	Get(exampleID string) (*repository.Example, *response.Error)
	GetFromGRPC(exampleID string) (*repository.Example, *response.Error)
	ChangeStatus(exampleID string, status string) *response.Error
	Delete(exampleIDs string) *response.Error
}

// business implements the Example Service
type business struct {
	repository  repository.Repository
	logger      log.Logger
	loggerClone log.Logger
}

// NewBusiness creates and returns a new Example service instance
func NewBusiness(rep repository.Repository, logger log.Logger) Service {
	return &business{
		repository:  rep,
		logger:      logger,
		loggerClone: logger,
	}
}

// Create to create an example
func (s *business) Create(example repository.Example) *response.Error {
	s.logger = s.loggerClone
	s.logger.SugaredLogger = s.logger.With("method", "Create")

	example.ExampleID = util.GetUUID()
	example.Status = "Pending"
	example.CreateTime = time.Now()

	err := s.repository.CreateExample(example)
	if err != nil {
		s.logger.Errorw("create data error", "error", err)
		return response.NewErrorFromCode(errorcode.CreateDataError)
	}

	return nil
}

// Update to update an example
func (s *business) Update(example repository.Example) *response.Error {
	s.logger = s.loggerClone
	s.logger.SugaredLogger = s.logger.With("method", "Update")

	body := repository.Example{
		ExampleName: example.ExampleName,
		Status:      example.Status,
	}
	err := s.repository.Update(example.ExampleID, body)
	if err != nil {
		s.logger.Errorw("update data error", "error", err)
		return response.NewErrorFromCode(errorcode.UpdateDataError)
	}

	return nil
}

// List to list examples
func (s *business) List(where repository.Example, order string, limit int32, page int32) (*repository.ExampleList, *response.Error) {
	s.logger = s.loggerClone
	s.logger.SugaredLogger = s.logger.With("method", "List")

	// Pagination
	if limit == 0 {
		limit = viper.GetInt32("PAGE_LIMIT")
	}
	if page <= 0 {
		page = 1
	}
	offset := limit * (page - 1)

	// Get data list
	examples, count, err := s.repository.ListWhere(where, order, limit, offset)
	if err != nil {
		s.logger.Errorw("list data error", "error", err)
		return nil, response.NewErrorFromCode(errorcode.GetDataError)
	}
	exampleList := repository.NewExampleList(examples, count)
	return exampleList, nil
}

// Get to get an example
func (s *business) Get(id string) (*repository.Example, *response.Error) {
	s.logger = s.loggerClone
	s.logger.SugaredLogger = s.logger.With("method", "Get")

	example, err := s.repository.GetExampleByExampleID(id)
	if err != nil {
		s.logger.Errorw("get data by id error", "error", err)

		var code int32
		if gorm.IsRecordNotFoundError(err) {
			code = errorcode.ResourceNotFound
		} else {
			code = errorcode.GetDataError
		}
		return nil, response.NewErrorFromCode(code)
	}
	return &example, nil
}

// GetFromGRPC is the same with Get method which data from grpc
func (s *business) GetFromGRPC(id string) (*repository.Example, *response.Error) {
	s.logger = s.loggerClone
	s.logger.SugaredLogger = s.logger.With("method", "GetFromGRPC")

	grpcClient := client.NewExampleGRPCClient(
		viper.GetString("GRPC_HOST_EXAMPLE"),
		viper.GetString("GRPC_PORT_EXAMPLE"),
		s.loggerClone,
	)
	res, err := grpcClient.Get(id)
	if err != nil {
		s.logger.Errorw("get data from grpc error", "error", err)
		return nil, response.NewErrorFromCode(errorcode.GetDataError)
	}

	t, _ := time.Parse(timeUtil.FullTimeFormat, res.CreateTime)
	example := &repository.Example{
		ID:          res.Id,
		ExampleID:   res.ExampleId,
		ExampleName: res.ExampleName,
		Status:      res.Status,
		CreateTime:  t,
	}

	return example, nil
}

// ChangeStatus to change the status of an example
func (s *business) ChangeStatus(id string, status string) *response.Error {
	s.logger = s.loggerClone
	s.logger.SugaredLogger = s.logger.With("method", "ChangeStatus")

	if id == "" {
		return response.NewErrorFromCode(errorcode.RequestParamError)
	}
	err := s.repository.ChangeStatus(id, status)
	if err != nil {
		s.logger.Errorw("change status error", "error", err)
		return response.NewErrorFromCode(errorcode.UpdateDataError)
	}
	return nil
}

// Delete to delete an example
func (s *business) Delete(exampleIDs string) *response.Error {
	s.logger = s.loggerClone
	s.logger.SugaredLogger = s.logger.With("method", "Delete")

	err := s.repository.DeleteIn(exampleIDs)
	if err != nil {
		s.logger.Errorw("delete data error", "error", err)
		return response.NewErrorFromCode(errorcode.DeleteDataError)
	}
	return nil
}
