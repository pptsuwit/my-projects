package service

import (
	"go-fiber-crud/app/config/logs"
	"go-fiber-crud/app/model"
	"go-fiber-crud/app/repository"
	"go-fiber-crud/app/utils"
)

type customerService struct {
	repository repository.CustomerRepository
}
type CustomerService interface {
	GetCustomers(model.Pagination) (model.CustomerResponseWithPagination, error)
	GetCustomer(id int) (*model.CustomerResponse, error)
	CreateCustomer(customer *model.CustomerRequest) (*model.CustomerResponse, error)
	UpdateCustomer(id int, customer *model.CustomerRequest) (*model.CustomerResponse, error)
	DeleteCustomer(id int) error
}

func NewCustomerService(repository repository.CustomerRepository) customerService {
	return customerService{repository: repository}
}

func (s customerService) GetCustomers(page model.Pagination) (model.CustomerResponseWithPagination, error) {
	entities, err, count := s.repository.GetAll(page)
	if err != nil {
		logs.Error(err)
		return model.CustomerResponseWithPagination{}, err
	}

	responseEntity := []model.CustomerResponse{}
	for _, item := range entities {
		responseEntity = append(responseEntity, model.CustomerResponse{
			ID:        item.ID,
			Name:      item.Name,
			Email:     item.Email,
			Phone:     item.Phone,
			Address:   item.Address,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
			Gender:    item.Gender,
			GenderId:  item.GenderId,
		})
	}
	response := model.CustomerResponseWithPagination{
		Customer: responseEntity,
		Pagination: model.PaginationResponse{
			RecordPerPage: page.PageSize,
			CurrentPage:   page.Page + 1,
			TotalPage:     utils.GetTotalPage(int(count), page.PageSize),
			TotalRecord:   int(count),
		},
	}
	return response, nil
}
func (s customerService) GetCustomer(id int) (*model.CustomerResponse, error) {
	customer, err := s.repository.GetById(id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	response := &model.CustomerResponse{
		ID:        customer.ID,
		Name:      customer.Name,
		Email:     customer.Email,
		Phone:     customer.Phone,
		Address:   customer.Address,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
		Gender:    customer.Gender,
		GenderId:  customer.GenderId,
	}
	return response, nil
}
func (s customerService) CreateCustomer(customer *model.CustomerRequest) (*model.CustomerResponse, error) {

	entity, err := s.repository.Create(customer)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	newCustomer := &model.CustomerResponse{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		Phone:     entity.Phone,
		Address:   entity.Address,
		GenderId:  entity.GenderId,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
	return newCustomer, nil
}
func (s customerService) UpdateCustomer(id int, customer *model.CustomerRequest) (*model.CustomerResponse, error) {

	entity, err := s.repository.Update(id, customer)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	newCustomer := &model.CustomerResponse{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		Phone:     entity.Phone,
		Address:   entity.Address,
		GenderId:  entity.GenderId,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
	return newCustomer, nil
}

func (s customerService) DeleteCustomer(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		logs.Error(err)
		return err
	}
	return nil
}
