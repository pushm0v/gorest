package repository

import (
	"github.com/pushm0v/gorest/model"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(cust *model.Customer) (*model.Customer, error)
	Update(cust *model.Customer, updateValue interface{}) (*model.Customer, error)
	Delete(cust *model.Customer) error
	FindOne(id int) (*model.Customer, error)
}

type customerRepository struct {
	dbConnection *gorm.DB
}

func NewCustomerRepository(dbConnection *gorm.DB) CustomerRepository {
	return &customerRepository{dbConnection: dbConnection}
}

func (c *customerRepository) Create(cust *model.Customer) (*model.Customer, error) {
	var err = c.dbConnection.Create(cust).Error
	return cust, err
}

func (c *customerRepository) FindOne(id int) (cust *model.Customer, err error) {
	cust = &model.Customer{}
	err = c.dbConnection.First(cust, id).Error

	return
}

func (c *customerRepository) Update(cust *model.Customer, updateValue interface{}) (*model.Customer, error) {
	return cust, c.dbConnection.Model(cust).Updates(updateValue).Error
}

func (c *customerRepository) Delete(cust *model.Customer) error {
	return c.dbConnection.Delete(cust).Error
}
